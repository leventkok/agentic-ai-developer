package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"learn/go/day-55/internal/auth"
	"learn/go/day-55/internal/db/testutil"
	"learn/go/day-55/internal/handler"
	"learn/go/day-55/internal/middleware"
	"learn/go/day-55/internal/repository/sqlite"
)

func TestAuthCapstoneFlows(t *testing.T) {
	tdb := testutil.OpenTestDB(t)
	testutil.ResetTables(t, tdb)

	tokens, err := auth.NewTokenService("test-secret-key-at-least-32-bytes-long", time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	authStore := sqlite.NewAuthStore(tdb.DB, tokens)
	repo, err := sqlite.New(tdb.DB)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Close()

	authHandler := &handler.AuthHandler{Auth: authStore}
	bookmarkHandler := &handler.BookmarkHandler{Repo: repo}
	requireAuth := middleware.RequireAuth(authStore)

	// wrong password
	loginReq := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(
		`{"email":"u@example.com","password":"wrong1234"}`,
	))
	loginRec := httptest.NewRecorder()
	authHandler.Login(loginRec, loginReq)
	if loginRec.Code != http.StatusUnauthorized {
		t.Fatalf("wrong password status = %d", loginRec.Code)
	}

	// register + login
	regReq := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(
		`{"email":"u@example.com","password":"password123"}`,
	))
	regRec := httptest.NewRecorder()
	authHandler.Register(regRec, regReq)
	if regRec.Code != http.StatusCreated {
		t.Fatalf("register status = %d", regRec.Code)
	}

	loginReq = httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(
		`{"email":"u@example.com","password":"password123"}`,
	))
	loginRec = httptest.NewRecorder()
	authHandler.Login(loginRec, loginReq)
	if loginRec.Code != http.StatusOK {
		t.Fatalf("login status = %d", loginRec.Code)
	}
	token := extractToken(t, loginRec.Body.String())

	// missing token on create
	createReq := httptest.NewRequest(http.MethodPost, "/bookmarks", strings.NewReader(
		`{"title":"Mine","url":"https://mine.dev"}`,
	))
	createRec := httptest.NewRecorder()
	bookmarkHandler.CreateBookmark(createRec, createReq)
	if createRec.Code != http.StatusUnauthorized {
		t.Fatalf("missing token status = %d", createRec.Code)
	}

	// create with token
	createHandler := requireAuth(http.HandlerFunc(bookmarkHandler.CreateBookmark))
	createReq = httptest.NewRequest(http.MethodPost, "/bookmarks", strings.NewReader(
		`{"title":"Mine","url":"https://mine.dev"}`,
	))
	createReq.Header.Set("Authorization", "Bearer "+token)
	createRec = httptest.NewRecorder()
	createHandler.ServeHTTP(createRec, createReq)
	if createRec.Code != http.StatusCreated {
		t.Fatalf("create status = %d body=%s", createRec.Code, createRec.Body.String())
	}

	// second user cannot update first user's bookmark
	regReq = httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(
		`{"email":"other@example.com","password":"password123"}`,
	))
	regRec = httptest.NewRecorder()
	authHandler.Register(regRec, regReq)

	loginReq = httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(
		`{"email":"other@example.com","password":"password123"}`,
	))
	loginRec = httptest.NewRecorder()
	authHandler.Login(loginRec, loginReq)
	otherToken := extractToken(t, loginRec.Body.String())

	list, err := repo.List(context.Background())
	if err != nil || len(list) == 0 {
		t.Fatal("expected bookmarks")
	}
	targetID := list[0].ID

	patchReq := httptest.NewRequest(http.MethodPatch, "/bookmarks/1", strings.NewReader(`{"title":"Hacked"}`))
	patchReq.SetPathValue("id", strconv.Itoa(targetID))
	patchReq.Header.Set("Authorization", "Bearer "+otherToken)
	patchRec := httptest.NewRecorder()
	requireAuth(http.HandlerFunc(bookmarkHandler.UpdateBookmark)).ServeHTTP(patchRec, patchReq)
	if patchRec.Code != http.StatusForbidden {
		t.Fatalf("cross-user update status = %d, want 403", patchRec.Code)
	}
}

func extractToken(t *testing.T, body string) string {
	t.Helper()
	const marker = `"token":"`
	i := strings.Index(body, marker)
	if i < 0 {
		t.Fatalf("token not found in %q", body)
	}
	rest := body[i+len(marker):]
	end := strings.Index(rest, `"`)
	if end < 0 {
		t.Fatal("malformed token json")
	}
	return rest[:end]
}
