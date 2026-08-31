package httpapi_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"learn/go/day-99/internal/httpapi"
)

func TestHealthLiveness(t *testing.T) {
	h := httpapi.NewHealthHandler(nil)
	rec := httptest.NewRecorder()
	h.Liveness(rec, httptest.NewRequest(http.MethodGet, "/health", nil))
	if rec.Code != http.StatusOK {
		t.Fatal(rec.Code)
	}
}

func TestHealthReadinessFailsWhenDBDown(t *testing.T) {
	h := httpapi.NewHealthHandler(func(context.Context) error { return errors.New("down") })
	rec := httptest.NewRecorder()
	h.Readiness(rec, httptest.NewRequest(http.MethodGet, "/ready", nil))
	if rec.Code != http.StatusServiceUnavailable {
		t.Fatal(rec.Code)
	}
}
