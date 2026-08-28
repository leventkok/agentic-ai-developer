package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"learn/go/day-49/internal/db/testutil"
	"learn/go/day-49/internal/handler"
	"learn/go/day-49/internal/repository/sqlite"
)

func TestCRUDSmoke(t *testing.T) {
	tdb := testutil.OpenTestDB(t)
	repo, err := sqlite.New(tdb.DB)
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Close()

	h := &handler.BookmarkHandler{Repo: repo}

	body := strings.NewReader(`{"title":"Smoke","url":"https://smoke.dev","tags":["test"]}`)
	req := httptest.NewRequest(http.MethodPost, "/bookmarks", body)
	rec := httptest.NewRecorder()
	h.CreateBookmark(rec, req)
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
	rec = httptest.NewRecorder()
	h.UpdateBookmark(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("update status = %d, want 200", rec.Code)
	}

	req = httptest.NewRequest(http.MethodDelete, "/bookmarks/1", nil)
	req.SetPathValue("id", "1")
	rec = httptest.NewRecorder()
	h.DeleteBookmark(rec, req)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("delete status = %d, want 204", rec.Code)
	}
}
