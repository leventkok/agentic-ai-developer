package handler_test

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"learn/go/day-45/internal/db"
	"learn/go/day-45/internal/handler"
	"learn/go/day-45/internal/store"
)

func migrationsDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Join(filepath.Dir(file), "..", "..", "migrations")
}

func TestCRUDSmoke(t *testing.T) {
	path := filepath.Join(t.TempDir(), "smoke.db")
	database, err := db.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()

	if err := db.RunMigrations(database, migrationsDir(t)); err != nil {
		t.Fatal(err)
	}

	repo, err := store.NewDBStore(database)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Close()

	h := &handler.BookmarkHandler{Store: repo}

	// Create
	body := strings.NewReader(`{"title":"Smoke","url":"https://smoke.dev","tags":["test"]}`)
	req := httptest.NewRequest(http.MethodPost, "/bookmarks", body)
	rec := httptest.NewRecorder()
	h.CreateBookmark(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("create status = %d, want 201", rec.Code)
	}

	// List
	req = httptest.NewRequest(http.MethodGet, "/bookmarks", nil)
	rec = httptest.NewRecorder()
	h.ListBookmarks(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("list status = %d, want 200", rec.Code)
	}

	// Get
	req = httptest.NewRequest(http.MethodGet, "/bookmarks/1", nil)
	req.SetPathValue("id", "1")
	rec = httptest.NewRecorder()
	h.GetBookmark(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("get status = %d, want 200", rec.Code)
	}

	// Update
	body = strings.NewReader(`{"title":"Smoke Updated"}`)
	req = httptest.NewRequest(http.MethodPatch, "/bookmarks/1", body)
	req.SetPathValue("id", "1")
	rec = httptest.NewRecorder()
	h.UpdateBookmark(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("update status = %d, want 200", rec.Code)
	}

	// Delete
	req = httptest.NewRequest(http.MethodDelete, "/bookmarks/1", nil)
	req.SetPathValue("id", "1")
	rec = httptest.NewRecorder()
	h.DeleteBookmark(rec, req)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("delete status = %d, want 204", rec.Code)
	}
}
