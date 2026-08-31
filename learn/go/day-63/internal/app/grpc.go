package app

import (
	"fmt"
	"time"

	"learn/go/day-63/internal/auth"
	"learn/go/day-63/internal/clock"
	"learn/go/day-63/internal/config"
	"learn/go/day-63/internal/db"
	"learn/go/day-63/internal/repository/sqlite"
	"learn/go/day-63/internal/service"
)

// GRPCDeps holds shared services for the standalone gRPC server.
type GRPCDeps struct {
	Bookmarks *service.BookmarkService
	Auth      *service.AuthService
}

// WireGRPCDeps builds bookmark and auth services backed by SQLite.
func WireGRPCDeps(cfg config.Config) (GRPCDeps, func() error, error) {
	database, err := db.Open(cfg.DBPath, cfg.PoolConfig())
	if err != nil {
		return GRPCDeps{}, nil, err
	}

	if err := db.RunMigrations(database, "migrations"); err != nil {
		database.Close()
		return GRPCDeps{}, nil, err
	}

	bookmarks, err := sqlite.New(database)
	if err != nil {
		database.Close()
		return GRPCDeps{}, nil, err
	}

	clk := clock.RealClock{}
	tokens, err := auth.NewTokenService(cfg.JWTSecret, cfg.JWTTTL(), clk)
	if err != nil {
		bookmarks.Close()
		database.Close()
		return GRPCDeps{}, nil, err
	}
	authStore := sqlite.NewAuthStore(database, tokens)

	bookmarkSvc := service.NewBookmarkService(bookmarks, time.Duration(cfg.ListTimeoutMS)*time.Millisecond)
	authSvc := service.NewAuthService(authStore)

	cleanup := func() error {
		var errs []error
		if err := bookmarks.Close(); err != nil {
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

	return GRPCDeps{Bookmarks: bookmarkSvc, Auth: authSvc}, cleanup, nil
}
