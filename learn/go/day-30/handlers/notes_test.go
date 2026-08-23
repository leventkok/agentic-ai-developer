package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)


func TestCreateNoteHandler(t *testing.T) {
	body := strings.NewReader(`{"title":"Learn Go","content":"Day 30"}`)
	req := httptest.NewRequest(http.MethodPost, "/notes", body)
	rec := httptest.NewRecorder()

	CreateNoteHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want 201", rec.Code)
	}
}


func TestDeleteNoteHandler_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/notes/999", nil)
	req.SetPathValue("id", "999")
	rec := httptest.NewRecorder()
	DeleteNoteHandler(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", rec.Code)
	}
}
