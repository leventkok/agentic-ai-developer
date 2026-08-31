package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"learn/go/day-81/internal/auth"
	"learn/go/day-81/internal/clock"
	"learn/go/day-81/internal/config"
	"learn/go/day-81/internal/db"
	bookmarksv1 "learn/go/day-81/internal/gen/bookmarksv1"
	"learn/go/day-81/internal/grpcapi"
	"learn/go/day-81/internal/httpapi"
	"learn/go/day-81/internal/repository"
	"learn/go/day-81/internal/repository/cached"
	"learn/go/day-81/internal/repository/sqlite"
	"learn/go/day-81/internal/repository/resilient"
	"learn/go/day-81/internal/service"
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

	bookmarkSvc := service.NewBookmarkService(bookmarkRepo, time.Duration(cfg.ListTimeoutMS)*time.Millisecond)
	authSvc := service.NewAuthService(authStore)

	httpHandler := httpapi.NewRouter(httpapi.Deps{
		Cfg:       cfg,
		Bookmarks: bookmarkRepo,
		Auth:      authStore,
	})

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcapi.UnaryServerInterceptor(authSvc, log.Default())))
	bookmarksv1.RegisterBookmarkServiceServer(grpcServer, grpcapi.NewServer(bookmarkSvc))

	cleanup := func() error {
		var errs []error
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
