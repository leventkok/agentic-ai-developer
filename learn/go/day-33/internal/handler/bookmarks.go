package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"learn/go/day-33/internal/model"
	"learn/go/day-33/internal/store"
)

// TODO: change Store field type from *store.MemoryStore to store.BookmarkRepository

type BookmarkHandler struct {
	Store store.BookmarkRepository
}

func (h *BookmarkHandler) ListBookmarks(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.Store.List())
}

func (h *BookmarkHandler) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	var req model.CreateBookmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}
	if req.URL == "" {
		writeError(w, http.StatusBadRequest, "url is required")
		return
	}
	writeJSON(w, http.StatusCreated, h.Store.Create(req))
}

func (h *BookmarkHandler) GetBookmark(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	bookmark, ok := h.Store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "bookmark not found")
		return
	}
	writeJSON(w, http.StatusOK, bookmark)
}

func (h *BookmarkHandler) UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	var req model.UpdateBookmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	bookmark, ok := h.Store.Update(id, req)
	if !ok {
		writeError(w, http.StatusNotFound, "bookmark not found")
		return
	}
	writeJSON(w, http.StatusOK, bookmark)
}

func (h *BookmarkHandler) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	if !h.Store.Delete(id) {
		writeError(w, http.StatusNotFound, "bookmark not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
