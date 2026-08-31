package httpapi_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"learn/go/day-98/internal/httpapi"
)

func TestHealthLiveness(t *testing.T) {
	h := httpapi.NewHealthHandler(nil)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	h.Liveness(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
}

func TestHealthReadinessFailsWhenDBDown(t *testing.T) {
	h := httpapi.NewHealthHandler(func(context.Context) error {
		return errors.New("db unreachable")
	})
	req := httptest.NewRequest(http.MethodGet, "/ready", nil)
	rec := httptest.NewRecorder()
	h.Readiness(rec, req)
	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d", rec.Code)
	}
}
