package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"learn/go/day-51/internal/repository/memory"
)

func TestRegister_ShortPassword(t *testing.T) {
	h := &AuthHandler{Auth: memory.NewAuth()}
	body := strings.NewReader(`{"email":"a@b.com","password":"short"}`)
	req := httptest.NewRequest(http.MethodPost, "/auth/register", body)
	rec := httptest.NewRecorder()

	h.Register(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestLogin_InvalidCredentials(t *testing.T) {
	auth := memory.NewAuth()
	h := &AuthHandler{Auth: auth}

	regBody := strings.NewReader(`{"email":"a@b.com","password":"password123"}`)
	regReq := httptest.NewRequest(http.MethodPost, "/auth/register", regBody)
	regRec := httptest.NewRecorder()
	h.Register(regRec, regReq)

	loginBody := strings.NewReader(`{"email":"a@b.com","password":"wrong"}`)
	loginReq := httptest.NewRequest(http.MethodPost, "/auth/login", loginBody)
	loginRec := httptest.NewRecorder()
	h.Login(loginRec, loginReq)

	if loginRec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", loginRec.Code)
	}
}
