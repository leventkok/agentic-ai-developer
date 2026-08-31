//go:build integration

package httpapi_test

import (
	"net/http"
	"strings"
	"testing"

	"learn/go/day-66/internal/test/fixtures"
)

// Slow integration tests — run with: go test -tags=integration ./internal/httpapi/...
func TestIntegration_Slow_ListAfterManyCreates(t *testing.T) {
	srv, authStore := setupHTTPServer(t)
	defer srv.Close()

	_, token := fixtures.RegisterUser(t, authStore, "slow@example.com", "password123")

	for i := 0; i < 5; i++ {
		body := strings.NewReader(`{"title":"B","url":"https://b.dev"}`)
		req, _ := http.NewRequest(http.MethodPost, srv.URL+"/bookmarks", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("create %d status = %d", i, resp.StatusCode)
		}
	}

	resp, err := http.Get(srv.URL + "/bookmarks")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("list status = %d", resp.StatusCode)
	}
}
