package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestID_SetsHeaderAndContext(t *testing.T) {
	var gotID string
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, ok := RequestIDFromContext(r.Context())
		if !ok || id == "" {
			t.Fatal("request ID missing from context")
		}
		gotID = id
	})

	h := RequestID(next)
	req := httptest.NewRequest(http.MethodGet, "/bookmarks", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	headerID := rec.Header().Get("X-Request-ID")
	if headerID == "" {
		t.Fatal("X-Request-ID header not set")
	}
	if headerID != gotID {
		t.Fatalf("header %q != context %q", headerID, gotID)
	}
}