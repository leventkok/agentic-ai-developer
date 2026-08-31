package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"learn/go/day-55/internal/model"
	"learn/go/day-55/internal/repository"
	"learn/go/day-55/internal/validation"
)

type AuthHandler struct {
	Auth repository.Auth
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if err := validation.ValidateRegister(req.Email, req.Password); err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.Auth.Register(r.Context(), req.Email, req.Password)
	if errors.Is(err, repository.ErrDuplicateEmail) {
		writeProblem(w, http.StatusConflict, "email already registered")
		return
	}
	if err != nil {
		writeProblem(w, http.StatusInternalServerError, "internal error")
		return
	}
	writeJSON(w, http.StatusCreated, user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if err := validation.ValidateLogin(req.Email, req.Password); err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Auth.Login(r.Context(), req.Email, req.Password)
	if errors.Is(err, repository.ErrInvalidCredentials) {
		writeProblem(w, http.StatusUnauthorized, "invalid email or password")
		return
	}
	if err != nil {
		writeProblem(w, http.StatusInternalServerError, "internal error")
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	user, ok := UserFromContext(r.Context())
	if !ok {
		writeProblem(w, http.StatusUnauthorized, "authentication required")
		return
	}
	writeJSON(w, http.StatusOK, user)
}
