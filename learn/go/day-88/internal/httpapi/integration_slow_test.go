//go:build integration

package httpapi_test

import (
	"net/http"
	"strings"
	"testing"

	"learn/go/day-88/internal/test/env"
	"learn/go/day-88/internal/test/fixtures"
)

// Slow integration tests — run with: go test -tags=integration ./...
func TestIntegration_Slow_ListAfterManyCreates(t *testing.T) {
	h := env.SetupHTTP(t)

	_, token := fixtures.RegisterUser(t, h.Auth, "slow@example.com", "password123")

	for i := 0; i < 5; i++ {
		body := strings.NewReader(`{"title":"B","url":"https://b.dev"}`)
		req, _ := http.NewRequest(http.MethodPost, h.Server.URL+"/bookmarks", body)
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

	resp, err := http.Get(h.Server.URL + "/bookmarks")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("list status = %d", resp.StatusCode)
	}
}
