package httpapi

import (
	"encoding/json"
	"net/http"

	"learn/go/day-100/internal/ctxkey"
	"learn/go/day-100/internal/domain"
	"learn/go/day-100/internal/service"
	"learn/go/day-100/internal/validation"
)

type BookmarkHandler struct {
	Svc *service.BookmarkService
}

func NewBookmarkHandler(svc *service.BookmarkService) *BookmarkHandler {
	return &BookmarkHandler{Svc: svc}
}

func (h *BookmarkHandler) ListBookmarks(w http.ResponseWriter, r *http.Request) {
	list, err := h.Svc.List(r.Context())
	if writeDomainError(w, err) {
		return
	}
	writeJSON(w, http.StatusOK, list)
}

func (h *BookmarkHandler) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	user, ok := ctxkey.UserFromContext(r.Context())
	if !ok {
		writeProblem(w, http.StatusUnauthorized, "authentication required")
		return
	}

	var req CreateBookmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	bookmark, err := h.Svc.Create(r.Context(), user, toCreateInput(req))
	if writeDomainError(w, err) {
		return
	}
	writeJSON(w, http.StatusCreated, bookmark)
}

func (h *BookmarkHandler) BulkCreateBookmarks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Deprecation", "true")
	w.Header().Set("Sunset", "2027-01-01")

	user, ok := ctxkey.UserFromContext(r.Context())
	if !ok {
		writeProblem(w, http.StatusUnauthorized, "authentication required")
		return
	}

	var req BulkCreateBookmarksRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	inputs := make([]domain.CreateBookmarkInput, len(req.Bookmarks))
	for i, b := range req.Bookmarks {
		inputs[i] = toCreateInput(b)
	}

	bookmarks, err := h.Svc.BulkCreate(r.Context(), user, inputs)
	if writeDomainError(w, err) {
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
	bookmark, err := h.Svc.Get(r.Context(), id)
	if writeDomainError(w, err) {
		return
	}
	writeJSON(w, http.StatusOK, bookmark)
}

func (h *BookmarkHandler) UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	user, ok := ctxkey.UserFromContext(r.Context())
	if !ok {
		writeProblem(w, http.StatusUnauthorized, "authentication required")
		return
	}

	id, err := validation.ParseID(r.PathValue("id"))
	if err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}
	var req UpdateBookmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	bookmark, err := h.Svc.Update(r.Context(), user, id, toUpdateInput(req))
	if writeDomainError(w, err) {
		return
	}
	writeJSON(w, http.StatusOK, bookmark)
}

func (h *BookmarkHandler) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	user, ok := ctxkey.UserFromContext(r.Context())
	if !ok {
		writeProblem(w, http.StatusUnauthorized, "authentication required")
		return
	}

	id, err := validation.ParseID(r.PathValue("id"))
	if err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}
	if writeDomainError(w, h.Svc.Delete(r.Context(), user, id)) {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func toCreateInput(req CreateBookmarkRequest) domain.CreateBookmarkInput {
	return domain.CreateBookmarkInput{Title: req.Title, URL: req.URL, Tags: req.Tags}
}

func toUpdateInput(req UpdateBookmarkRequest) domain.UpdateBookmarkInput {
	return domain.UpdateBookmarkInput{Title: req.Title, URL: req.URL, Tags: req.Tags}
}
