package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	notes   = map[int]Note{}
	nextID  = 1
	notesMu sync.RWMutex
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]any{"code": status, "message": message})
}

// ListNotesHandler handles GET /notes
func ListNotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	notesMu.RLock()
	defer notesMu.RUnlock()

	list := make([]Note, 0, len(notes))
	for _, n := range notes {
		list = append(list, n)
	}
	writeJSON(w, http.StatusOK, list)
}

// CreateNoteHandler handles POST /notes
func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	notesMu.Lock()
	note := Note{ID: nextID, Title: req.Title, Content: req.Content}
	notes[nextID] = note
	nextID++
	notesMu.Unlock()

	writeJSON(w, http.StatusCreated, note)
}

// DeleteNoteHandler handles DELETE /notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	notesMu.Lock()
	defer notesMu.Unlock()
	if _, ok := notes[id]; !ok {
		writeError(w, http.StatusNotFound, "note not found")
		return
	}
	delete(notes, id)
	w.WriteHeader(http.StatusNoContent)
}
