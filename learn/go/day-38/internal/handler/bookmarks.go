package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"learn/go/day-38/internal/model"
	"learn/go/day-38/internal/store"
	"learn/go/day-38/internal/validation"
)

type BookmarkHandler struct {
	Store       store.BookmarkRepository
	ListTimeout time.Duration
}

// handleStoreError maps store-layer errors to HTTP responses.
// Returns true when an error was written (caller should return).
func handleStoreError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		writeProblem(w, http.StatusRequestTimeout, err.Error())
		return true
	}
	if errors.Is(err, store.ErrNotFound) {
		writeProblem(w, http.StatusNotFound, "bookmark not found")
		return true
	}
	writeProblem(w, http.StatusInternalServerError, "internal error")
	return true
}

func (h *BookmarkHandler) ListBookmarks(w http.ResponseWriter, r *http.Request) {
	// Bound how long List may run — slow loop in store respects ctx deadline.
	timeout := h.ListTimeout
	if timeout <= 0 {
		timeout = 100 * time.Millisecond
	}
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	list, err := h.Store.List(ctx)
	if handleStoreError(w, err) {
		return
	}
	writeJSON(w, http.StatusOK, list)
}

func (h *BookmarkHandler) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	var req model.CreateBookmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if err := validation.ValidateCreate(req); err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}
	bookmark, err := h.Store.Create(r.Context(), req)
	if handleStoreError(w, err) {
		return
	}
	writeJSON(w, http.StatusCreated, bookmark)
}

func (h *BookmarkHandler) GetBookmark(w http.ResponseWriter, r *http.Request) {
	id, err := validation.ParseID(r.PathValue("id"))
	if err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}
	bookmark, err := h.Store.Get(r.Context(), id)
	if handleStoreError(w, err) {
		return
	}
	writeJSON(w, http.StatusOK, bookmark)
}

func (h *BookmarkHandler) UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	id, err := validation.ParseID(r.PathValue("id"))
	if err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}
	var req model.UpdateBookmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if err := validation.ValidateUpdate(req); err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}
	bookmark, err := h.Store.Update(r.Context(), id, req)
	if handleStoreError(w, err) {
		return
	}
	writeJSON(w, http.StatusOK, bookmark)
}

func (h *BookmarkHandler) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	id, err := validation.ParseID(r.PathValue("id"))
	if err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}
	if handleStoreError(w, h.Store.Delete(r.Context(), id)) {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
