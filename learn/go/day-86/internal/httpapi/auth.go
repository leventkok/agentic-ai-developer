package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"learn/go/day-86/internal/ctxkey"
	"learn/go/day-86/internal/domain"
	"learn/go/day-86/internal/service"
	"learn/go/day-86/internal/validation"
)

type AuthHandler struct {
	Svc *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{Svc: svc}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if err := validation.ValidateRegister(req.Email, req.Password); err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.Svc.Register(r.Context(), req.Email, req.Password)
	if errors.Is(err, domain.ErrDuplicateEmail) {
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
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if err := validation.ValidateLogin(req.Email, req.Password); err != nil {
		writeProblem(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Svc.Login(r.Context(), req.Email, req.Password)
	if errors.Is(err, domain.ErrInvalidCredentials) {
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
	user, ok := ctxkey.UserFromContext(r.Context())
	if !ok {
		writeProblem(w, http.StatusUnauthorized, "authentication required")
		return
	}
	writeJSON(w, http.StatusOK, user)
}
