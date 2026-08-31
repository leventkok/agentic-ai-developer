package httpapi_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"learn/go/day-66/internal/auth"
	"learn/go/day-66/internal/config"
	"learn/go/day-66/internal/db/testutil"
	"learn/go/day-66/internal/httpapi"
	"learn/go/day-66/internal/repository"
	"learn/go/day-66/internal/repository/sqlite"
	"learn/go/day-66/internal/test/fixtures"
)

func setupHTTPServer(t *testing.T) (*httptest.Server, repository.Auth) {
	t.Helper()
	tdb := testutil.OpenTestDB(t)
	testutil.ResetTables(t, tdb)

	bookmarks, err := sqlite.New(tdb.DB)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { bookmarks.Close() })

	tokens, err := auth.NewTokenService("test-secret-key-at-least-32-bytes-long", time.Hour, nil)
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
		JWTSecret:              "test-secret-key-at-least-32-bytes-long",
		JWTTTLHours:            24,
		AuthRateLimitPerMinute: 100,
		GRPCPort:               "9090",
	}

	handler := httpapi.NewRouter(httpapi.Deps{
		Cfg:       cfg,
		Bookmarks: bookmarks,
		Auth:      authStore,
	})
	return httptest.NewServer(handler), authStore
}

func TestIntegration_AuthRegisterLogin(t *testing.T) {
	srv, authStore := setupHTTPServer(t)
	defer srv.Close()

	reg := `{"email":"int@example.com","password":"password123"}`
	resp, err := http.Post(srv.URL+"/auth/register", "application/json", strings.NewReader(reg))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("register status = %d, body = %s", resp.StatusCode, body)
	}

	login := `{"email":"int@example.com","password":"password123"}`
	resp, err = http.Post(srv.URL+"/auth/login", "application/json", strings.NewReader(login))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("login status = %d", resp.StatusCode)
	}

	_, token := fixtures.RegisterUser(t, authStore, "other@example.com", "password123")
	if token == "" {
		t.Fatal("expected token")
	}
}

func TestIntegration_BookmarkCRUDViaHTTP(t *testing.T) {
	srv, authStore := setupHTTPServer(t)
	defer srv.Close()

	_, token := fixtures.RegisterUser(t, authStore, "crud@example.com", "password123")

	createBody := `{"title":"Integration","url":"https://integration.dev","tags":["test"]}`
	req, err := http.NewRequest(http.MethodPost, srv.URL+"/bookmarks", strings.NewReader(createBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("create status = %d, body = %s", resp.StatusCode, body)
	}

	resp, err = http.Get(srv.URL + "/bookmarks")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("list status = %d", resp.StatusCode)
	}
}
