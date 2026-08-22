package cafe

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func doRequest(t *testing.T, handler http.HandlerFunc, method, path string, body *strings.Reader) *httptest.ResponseRecorder {
	t.Helper()
	var req *http.Request
	if body != nil {
		req = httptest.NewRequest(method, path, body)
		req.Header.Set("Content-Type", "application/json")
	} else {
		req = httptest.NewRequest(method, path, nil)
	}
	rec := httptest.NewRecorder()
	handler(rec, req)
	return rec
}

func TestOrderHandler(t *testing.T) {
	body := strings.NewReader(`{"name":"Latte","qty":2}`)
	rec := doRequest(t, OrderHandler, http.MethodPost, "/order", body)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "90") {
		t.Errorf("body = %s, want total 90", rec.Body.String())
	}
}

func TestMenuHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		wantStatus int
		wantBody   string
	}{
		{name: "GET ok", method: http.MethodGet, wantStatus: http.StatusOK, wantBody: "Latte"},
		{name: "POST not allowed", method: http.MethodPost, wantStatus: http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := doRequest(t, MenuHandler, tt.method, "/menu", nil)
			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if tt.wantBody != "" && !strings.Contains(rec.Body.String(), tt.wantBody) {
				t.Errorf("body = %q, want to contain %q", rec.Body.String(), tt.wantBody)
			}
		})
	}
}
