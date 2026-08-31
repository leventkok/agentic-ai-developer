package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"learn/go/day-52/internal/handler"
	"learn/go/day-52/internal/middleware"
	"learn/go/day-52/internal/repository/memory"
	"learn/go/day-52/internal/auth"
)

func testTokens(t *testing.T) *auth.TokenService {
	t.Helper()
	svc, err := auth.NewTokenService("test-secret-key-at-least-32-bytes-long", time.Hour)
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

func TestRequireAuth_ValidJWT(t *testing.T) {
	authStore := memory.NewAuth(testTokens(t))
	ctx := httptest.NewRequest(http.MethodGet, "/", nil).Context()
	user, token, err := authStore.RegisterAndLogin(ctx, "bob@example.com", "password123")
	if err != nil {
		t.Fatal(err)
	}

	var gotEmail string
	protected := middleware.RequireAuth(authStore)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, ok := handler.UserFromContext(r.Context())
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
