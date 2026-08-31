package httpapi_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"learn/go/day-96/internal/auth"
	"learn/go/day-96/internal/httpapi"
	"learn/go/day-96/internal/repository/memory"
	"learn/go/day-96/internal/service"
)

func testAuth(t *testing.T) *memory.Auth {
	t.Helper()
	svc, err := auth.NewTokenService("test-secret-key-at-least-32-bytes-long", time.Hour, nil)
	if err != nil {
		t.Fatal(err)
	}
	return memory.NewAuth(svc)
}

func TestRegister_ShortPassword(t *testing.T) {
	h := httpapi.NewAuthHandler(service.NewAuthService(testAuth(t)))
	body := strings.NewReader(`{"email":"a@b.com","password":"short"}`)
	req := httptest.NewRequest(http.MethodPost, "/auth/register", body)
	rec := httptest.NewRecorder()

	h.Register(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestLogin_InvalidCredentials(t *testing.T) {
	authStore := testAuth(t)
	h := httpapi.NewAuthHandler(service.NewAuthService(authStore))

	regBody := strings.NewReader(`{"email":"a@b.com","password":"password123"}`)
	regReq := httptest.NewRequest(http.MethodPost, "/auth/register", regBody)
	regRec := httptest.NewRecorder()
	h.Register(regRec, regReq)

	loginBody := strings.NewReader(`{"email":"a@b.com","password":"wrong"}`)
	loginReq := httptest.NewRequest(http.MethodPost, "/auth/login", loginBody)
	loginRec := httptest.NewRecorder()
	h.Login(loginRec, loginReq)

	if loginRec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", loginRec.Code)
	}
}

func TestCreateBookmark_MissingTitle(t *testing.T) {
	h := httpapi.NewBookmarkHandler(service.NewBookmarkService(memory.New(nil), 100*time.Millisecond, nil, nil))
	body := strings.NewReader(`{"title":"","url":"https://go.dev"}`)
	req := httptest.NewRequest(http.MethodPost, "/bookmarks", body)
	rec := httptest.NewRecorder()

	h.CreateBookmark(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401 without auth", rec.Code)
	}
}
