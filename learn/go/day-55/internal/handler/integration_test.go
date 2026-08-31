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

func TestCRUDSmoke(t *testing.T) {
	tdb := testutil.OpenTestDB(t)
	repo, err := sqlite.New(tdb.DB)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Close()

	tokens, err := auth.NewTokenService("test-secret-key-at-least-32-bytes-long", time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	authStore := sqlite.NewAuthStore(tdb.DB, tokens)
	ctx := context.Background()
	user, err := authStore.Register(ctx, "smoke@example.com", "password123")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := authStore.Login(ctx, "smoke@example.com", "password123")
	if err != nil {
		t.Fatal(err)
	}
	_ = resp

	h := &handler.BookmarkHandler{Repo: repo}
	requireAuth := middleware.RequireAuth(authStore)

	body := strings.NewReader(`{"title":"Smoke","url":"https://smoke.dev","tags":["test"]}`)
	req := httptest.NewRequest(http.MethodPost, "/bookmarks", body)
	req.Header.Set("Authorization", "Bearer "+resp.Token)
	rec := httptest.NewRecorder()
	requireAuth(http.HandlerFunc(h.CreateBookmark)).ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("create status = %d, want 201", rec.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/bookmarks", nil)
	rec = httptest.NewRecorder()
	h.ListBookmarks(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("list status = %d, want 200", rec.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/bookmarks/1", nil)
	req.SetPathValue("id", "1")
	rec = httptest.NewRecorder()
	h.GetBookmark(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("get status = %d, want 200", rec.Code)
	}

	body = strings.NewReader(`{"title":"Smoke Updated"}`)
	req = httptest.NewRequest(http.MethodPatch, "/bookmarks/1", body)
	req.SetPathValue("id", "1")
	req.Header.Set("Authorization", "Bearer "+resp.Token)
	rec = httptest.NewRecorder()
	requireAuth(http.HandlerFunc(h.UpdateBookmark)).ServeHTTP(rec, req)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("update seeded bookmark status = %d, want 403 for member", rec.Code)
	}

	// create own bookmark and update it
	body = strings.NewReader(`{"title":"Owned","url":"https://owned.dev"}`)
	req = httptest.NewRequest(http.MethodPost, "/bookmarks", body)
	req.Header.Set("Authorization", "Bearer "+resp.Token)
	rec = httptest.NewRecorder()
	requireAuth(http.HandlerFunc(h.CreateBookmark)).ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("create owned status = %d", rec.Code)
	}

	list, _ := repo.List(ctx)
	var ownedBookmarkID int
	for _, b := range list {
		if b.UserID != nil && *b.UserID == user.ID {
			ownedBookmarkID = b.ID
			break
		}
	}
	if ownedBookmarkID == 0 {
		t.Fatal("owned bookmark not found")
	}

	body = strings.NewReader(`{"title":"Owned Updated"}`)
	req = httptest.NewRequest(http.MethodPatch, "/bookmarks/"+strconv.Itoa(ownedBookmarkID), body)
	req.SetPathValue("id", strconv.Itoa(ownedBookmarkID))
	req.Header.Set("Authorization", "Bearer "+resp.Token)
	rec = httptest.NewRecorder()
	requireAuth(http.HandlerFunc(h.UpdateBookmark)).ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("update owned status = %d", rec.Code)
	}

	req = httptest.NewRequest(http.MethodDelete, "/bookmarks/"+strconv.Itoa(ownedBookmarkID), nil)
	req.SetPathValue("id", strconv.Itoa(ownedBookmarkID))
	req.Header.Set("Authorization", "Bearer "+resp.Token)
	rec = httptest.NewRecorder()
	requireAuth(http.HandlerFunc(h.DeleteBookmark)).ServeHTTP(rec, req)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("delete status = %d, want 204", rec.Code)
	}
}
