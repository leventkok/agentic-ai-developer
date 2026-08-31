package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"learn/go/day-89/internal/auth"
	"learn/go/day-89/internal/ctxkey"
	"learn/go/day-89/internal/domain"
	"learn/go/day-89/internal/middleware"
	"learn/go/day-89/internal/repository/memory"
)

func testTokens(t *testing.T) *auth.TokenService {
	t.Helper()
	svc, err := auth.NewTokenService("test-secret-key-at-least-32-bytes-long", time.Hour, nil)
	if err != nil {
		t.Fatal(err)
	}
	return svc
}

func TestRequireAuth_MissingHeader(t *testing.T) {
	authStore := memory.NewAuth(testTokens(t))
	protected := middleware.RequireAuth(authStore)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
	rec := httptest.NewRecorder()
	protected.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", rec.Code)
	}
}

func TestRequireAuth_ValidToken(t *testing.T) {
	authStore := memory.NewAuth(testTokens(t))
	ctx := httptest.NewRequest(http.MethodGet, "/", nil).Context()
	user, token, err := authStore.RegisterAndLogin(ctx, "bob@example.com", "password123")
	if err != nil {
		t.Fatal(err)
	}

	var gotEmail string
	protected := middleware.RequireAuth(authStore)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, ok := ctxkey.UserFromContext(r.Context())
		if !ok {
			t.Fatal("user missing from context")
		}
		gotEmail = u.Email
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	protected.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if gotEmail != user.Email {
		t.Fatalf("email = %q, want %q", gotEmail, user.Email)
	}
}

func TestRequireRole_Forbidden(t *testing.T) {
	authStore := memory.NewAuth(testTokens(t))
	ctx := httptest.NewRequest(http.MethodGet, "/", nil).Context()
	_, token, err := authStore.RegisterAndLogin(ctx, "member@example.com", "password123")
	if err != nil {
		t.Fatal(err)
	}

	chain := middleware.RequireAuth(authStore)(middleware.RequireRole(domain.RoleAdmin)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})))

	req := httptest.NewRequest(http.MethodPost, "/bookmarks/bulk", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	chain.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("status = %d, want 403", rec.Code)
	}
}

func TestRateLimiter_BlocksBurst(t *testing.T) {
	limiter := middleware.NewRateLimiter(2, time.Minute)
	h := limiter.Limit(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	for i := 0; i < 2; i++ {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/auth/login", nil))
		if rec.Code != http.StatusOK {
			t.Fatalf("request %d status = %d", i, rec.Code)
		}
	}

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/auth/login", nil))
	if rec.Code != http.StatusTooManyRequests {
		t.Fatalf("status = %d, want 429", rec.Code)
	}
}
