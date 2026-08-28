package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"learn/go/day-50/internal/repository/memory"
)

func newTestHandler() *BookmarkHandler {
	return &BookmarkHandler{Repo: memory.New()}
}

func TestCreateBookmark_MissingTitle(t *testing.T) {
	h := newTestHandler()
	body := strings.NewReader(`{"title":"","url":"https://go.dev"}`)
	req := httptest.NewRequest(http.MethodPost, "/bookmarks", body)
	rec := httptest.NewRecorder()

	h.CreateBookmark(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestCreateBookmark_InvalidURL(t *testing.T) {
	h := newTestHandler()
	body := strings.NewReader(`{"title":"Go","url":"not-a-url"}`)
	req := httptest.NewRequest(http.MethodPost, "/bookmarks", body)
	rec := httptest.NewRecorder()

	h.CreateBookmark(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestGetBookmark_NotFound(t *testing.T) {
	h := newTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/bookmarks/999", nil)
	req.SetPathValue("id", "999")
	rec := httptest.NewRecorder()

	h.GetBookmark(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
}

func TestGetBookmark_InvalidID(t *testing.T) {
	h := newTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/bookmarks/abc", nil)
	req.SetPathValue("id", "abc")
	rec := httptest.NewRecorder()

	h.GetBookmark(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}
