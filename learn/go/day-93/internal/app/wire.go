package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"learn/go/day-93/internal/auth"
	"learn/go/day-93/internal/clock"
	"learn/go/day-93/internal/config"
	"learn/go/day-93/internal/db"
	bookmarksv1 "learn/go/day-93/internal/gen/bookmarksv1"
	"learn/go/day-93/internal/grpcapi"
	"learn/go/day-93/internal/httpapi"
	"learn/go/day-93/internal/messaging/outbox"
	"learn/go/day-93/internal/messaging/relay"
	applog "learn/go/day-93/internal/observability/log"
	"learn/go/day-93/internal/repository"
	"learn/go/day-93/internal/repository/cached"
	"learn/go/day-93/internal/repository/sqlite"
	"learn/go/day-93/internal/repository/resilient"
	"learn/go/day-93/internal/service"
)

// Runtime bundles HTTP and gRPC servers with shared dependencies.
type Runtime struct {
	HTTPHandler http.Handler
	GRPCServer  *grpc.Server
	Cleanup     func() error
}

// Wire builds HTTP and gRPC runtimes from configuration.
func Wire(cfg config.Config) (*Runtime, error) {
	database, err := db.Open(cfg.DBPath, cfg.PoolConfig())
	if err != nil {
		return nil, err
	}

	if err := db.RunMigrations(database, "migrations"); err != nil {
		database.Close()
		return nil, err
	}

	sqliteStore, err := sqlite.New(database)
	if err != nil {
		database.Close()
		return nil, err
	}
	var bookmarkRepo repository.Bookmarks = resilient.NewBookmarks(sqliteStore)

	cacheStore, err := openCache(cfg)
	if err != nil {
		sqliteStore.Close()
		database.Close()
		return nil, err
	}
	bookmarkRepo = cached.New(bookmarkRepo, cacheStore, cfg.CacheTTL())

	clk := clock.RealClock{}
	tokens, err := auth.NewTokenService(cfg.JWTSecret, cfg.JWTTTL(), clk)
	if err != nil {
		_ = closeCache(cacheStore)
		sqliteStore.Close()
		database.Close()
		return nil, err
	}
	authStore := sqlite.NewAuthStore(database, tokens)

	eventBus, err := OpenBus(cfg)
	if err != nil {
		_ = closeCache(cacheStore)
		sqliteStore.Close()
		database.Close()
		return nil, err
	}

	outboxStore := outbox.NewStore(database)

	bookmarkSvc := service.NewBookmarkService(bookmarkRepo, time.Duration(cfg.ListTimeoutMS)*time.Millisecond, outboxStore, eventBus)
	authSvc := service.NewAuthService(authStore)

	relayCtx, relayCancel := context.WithCancel(context.Background())
	go func() {
		r := relay.Relay{
			Outbox:  outboxStore,
			Bus:     eventBus,
			Logger:  applog.New(cfg.Env),
			Workers: cfg.OutboxWorkers,
		}
		r.Run(relayCtx)
	}()

	httpHandler := httpapi.NewRouter(httpapi.Deps{
		Cfg:         cfg,
		Bookmarks:   bookmarkRepo,
		BookmarkSvc: bookmarkSvc,
		Auth:        authStore,
	})

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcapi.UnaryServerInterceptor(authSvc, log.Default())))
	bookmarksv1.RegisterBookmarkServiceServer(grpcServer, grpcapi.NewServer(bookmarkSvc))

	cleanup := func() error {
		var errs []error
		relayCancel()
		if err := busCloseErr(CloseBus(eventBus)); err != nil {
			errs = append(errs, err)
		}
		if err := cacheCloseErr(closeCache(cacheStore)); err != nil {
			errs = append(errs, err)
		}
		if err := sqliteStore.Close(); err != nil {
			errs = append(errs, err)
		}
		if err := database.Close(); err != nil {
			errs = append(errs, err)
		}
		if len(errs) > 0 {
			return fmt.Errorf("shutdown: %v", errs)
		}
		return nil
	}

	return &Runtime{
		HTTPHandler: httpHandler,
		GRPCServer:  grpcServer,
		Cleanup:     cleanup,
	}, nil
}

// WireGRPC is kept for callers that only need the gRPC server wiring path.
func WireGRPC(cfg config.Config) (*grpc.Server, func() error, error) {
	runtime, err := Wire(cfg)
	if err != nil {
		return nil, nil, err
	}
	return runtime.GRPCServer, runtime.Cleanup, nil
}
