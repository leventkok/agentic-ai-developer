package handler

import (
	"encoding/json"
	"net/http"

	"learn/go/day-35/internal/model"
	"learn/go/day-35/internal/store"
	"learn/go/day-35/internal/validation"
)

type BookmarkHandler struct {
	Store store.BookmarkRepository
}

func (h *BookmarkHandler) ListBookmarks(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.Store.List())
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
	writeJSON(w, http.StatusCreated, h.Store.Create(req))
}

func (h *BookmarkHandler) GetBookmark(w http.ResponseWriter, r *http.Request) {
	id, err := validation.ParseID(r.PathValue("id"))
	if err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}
	bookmark, ok := h.Store.Get(id)
	if !ok {
		writeProblem(w, http.StatusNotFound, "bookmark not found")
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
	bookmark, ok := h.Store.Update(id, req)
	if !ok {
		writeProblem(w, http.StatusNotFound, "bookmark not found")
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
	if !h.Store.Delete(id) {
		writeProblem(w, http.StatusNotFound, "bookmark not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
