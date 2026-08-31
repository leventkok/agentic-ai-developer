package env

import (
	"net/http/httptest"
	"testing"
	"time"

	"learn/go/day-67/internal/auth"
	"learn/go/day-67/internal/config"
	"learn/go/day-67/internal/db/testutil"
	"learn/go/day-67/internal/httpapi"
	"learn/go/day-67/internal/repository"
	"learn/go/day-67/internal/repository/sqlite"
)

// HTTP bundles an httptest server wired to real SQLite repositories.
type HTTP struct {
	Server *httptest.Server
	Auth   repository.Auth
}

// SetupHTTP creates an isolated test database and HTTP server.
// Resources are torn down automatically via t.Cleanup, even when tests fail.
func SetupHTTP(t *testing.T) *HTTP {
	t.Helper()

	tdb := testutil.OpenTestDB(t)
	testutil.ResetTables(t, tdb)

	bookmarks, err := sqlite.New(tdb.DB)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { bookmarks.Close() })

	const jwtSecret = "test-secret-key-at-least-32-bytes-long"
	tokens, err := auth.NewTokenService(jwtSecret, time.Hour, nil)
	if err != nil {
		t.Fatal(err)
	}
	authStore := sqlite.NewAuthStore(tdb.DB, tokens)

	cfg := config.Config{
		Port:                   "8080",
		ListTimeoutMS:          500,
		Env:                    "development",
		ReadTimeoutSec:         5,
		WriteTimeoutSec:        10,
		ShutdownTimeoutSec:     15,
		DBPath:                 tdb.Path,
		DBMaxOpenConns:         1,
		DBMaxIdleConns:         1,
		JWTSecret:              jwtSecret,
		JWTTTLHours:            24,
		AuthRateLimitPerMinute: 100,
		GRPCPort:               "9090",
	}

	handler := httpapi.NewRouter(httpapi.Deps{
		Cfg:       cfg,
		Bookmarks: bookmarks,
		Auth:      authStore,
	})
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	return &HTTP{Server: srv, Auth: authStore}
}
