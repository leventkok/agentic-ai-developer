package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"learn/go/day-51/internal/model"
	"learn/go/day-51/internal/repository"
	"learn/go/day-51/internal/validation"
)

type BookmarkHandler struct {
	Repo        repository.Bookmarks
	ListTimeout time.Duration
}

func handleRepoError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		writeProblem(w, http.StatusRequestTimeout, err.Error())
		return true
	}
	if errors.Is(err, repository.ErrNotFound) {
		writeProblem(w, http.StatusNotFound, "bookmark not found")
		return true
	}
	writeProblem(w, http.StatusInternalServerError, "internal error")
	return true
}

func (h *BookmarkHandler) ListBookmarks(w http.ResponseWriter, r *http.Request) {
	timeout := h.ListTimeout
	if timeout <= 0 {
		timeout = 100 * time.Millisecond
	}
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	list, err := h.Repo.List(ctx)
	if handleRepoError(w, err) {
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
	bookmark, err := h.Repo.Create(r.Context(), req)
	if handleRepoError(w, err) {
		return
	}
	writeJSON(w, http.StatusCreated, bookmark)
}

func (h *BookmarkHandler) BulkCreateBookmarks(w http.ResponseWriter, r *http.Request) {
	var req model.BulkCreateBookmarksRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if err := validation.ValidateBulkCreate(req); err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}
	bookmarks, err := h.Repo.BulkCreate(r.Context(), req.Bookmarks)
	if handleRepoError(w, err) {
		return
	}
	writeJSON(w, http.StatusCreated, bookmarks)
}

func (h *BookmarkHandler) GetBookmark(w http.ResponseWriter, r *http.Request) {
	id, err := validation.ParseID(r.PathValue("id"))
	if err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}
	bookmark, err := h.Repo.Get(r.Context(), id)
	if handleRepoError(w, err) {
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
	bookmark, err := h.Repo.Update(r.Context(), id, req)
	if handleRepoError(w, err) {
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
	if handleRepoError(w, h.Repo.Delete(r.Context(), id)) {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
