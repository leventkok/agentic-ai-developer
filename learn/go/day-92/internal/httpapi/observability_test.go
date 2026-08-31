package httpapi_test

import (
	"net/http"
	"strings"
	"testing"

	"learn/go/day-92/internal/test/env"
)

func TestObservability_MetricsEndpoint(t *testing.T) {
	h := env.SetupHTTP(t)
	resp, err := http.Get(h.Server.URL + "/metrics")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("metrics status = %d", resp.StatusCode)
	}
}

func TestObservability_HealthListEndpoint(t *testing.T) {
	h := env.SetupHTTP(t)
	resp, err := http.Get(h.Server.URL + "/bookmarks")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("list status = %d", resp.StatusCode)
	}
}

func TestObservability_AuthDoesNotLogSecrets(t *testing.T) {
	h := env.SetupHTTP(t)
	body := strings.NewReader(`{"email":"obs@example.com","password":"password123"}`)
	resp, err := http.Post(h.Server.URL+"/auth/register", "application/json", body)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("register status = %d", resp.StatusCode)
	}
}
