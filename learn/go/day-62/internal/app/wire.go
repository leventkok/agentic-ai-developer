package app

import (
	"fmt"
	"net/http"

	"learn/go/day-62/internal/auth"
	"learn/go/day-62/internal/clock"
	"learn/go/day-62/internal/config"
	"learn/go/day-62/internal/db"
	"learn/go/day-62/internal/httpapi"
	"learn/go/day-62/internal/repository/sqlite"
)

// Wire builds the HTTP handler and a cleanup function from configuration.
func Wire(cfg config.Config) (http.Handler, func() error, error) {
	database, err := db.Open(cfg.DBPath, cfg.PoolConfig())
	if err != nil {
		return nil, nil, err
	}

	if err := db.RunMigrations(database, "migrations"); err != nil {
		database.Close()
		return nil, nil, err
	}

	bookmarks, err := sqlite.New(database)
	if err != nil {
		database.Close()
		return nil, nil, err
	}

	clk := clock.RealClock{}
	tokens, err := auth.NewTokenService(cfg.JWTSecret, cfg.JWTTTL(), clk)
	if err != nil {
		bookmarks.Close()
		database.Close()
		return nil, nil, err
	}
	authStore := sqlite.NewAuthStore(database, tokens)

	handler := httpapi.NewRouter(httpapi.Deps{
		Cfg:       cfg,
		Bookmarks: bookmarks,
		Auth:      authStore,
	})

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

	return handler, cleanup, nil
}
