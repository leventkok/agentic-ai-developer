package httpapi_test

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"learn/go/day-81/internal/test/env"
	"learn/go/day-81/internal/test/fixtures"
)

func TestIntegration_AuthRegisterLogin(t *testing.T) {
	h := env.SetupHTTP(t)

	reg := `{"email":"int@example.com","password":"password123"}`
	resp, err := http.Post(h.Server.URL+"/auth/register", "application/json", strings.NewReader(reg))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("register status = %d, body = %s", resp.StatusCode, body)
	}

	login := `{"email":"int@example.com","password":"password123"}`
	resp, err = http.Post(h.Server.URL+"/auth/login", "application/json", strings.NewReader(login))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("login status = %d", resp.StatusCode)
	}

	_, token := fixtures.RegisterUser(t, h.Auth, "other@example.com", "password123")
	if token == "" {
		t.Fatal("expected token")
	}
}

func TestIntegration_BookmarkCRUDViaHTTP(t *testing.T) {
	h := env.SetupHTTP(t)

	_, token := fixtures.RegisterUser(t, h.Auth, "crud@example.com", "password123")

	createBody := `{"title":"Integration","url":"https://integration.dev","tags":["test"]}`
	req, err := http.NewRequest(http.MethodPost, h.Server.URL+"/bookmarks", strings.NewReader(createBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("create status = %d, body = %s", resp.StatusCode, body)
	}

	var created struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		t.Fatal(err)
	}

	patchBody := `{"title":"Updated"}`
	req, err = http.NewRequest(http.MethodPatch, h.Server.URL+"/bookmarks/"+strconv.Itoa(created.ID), strings.NewReader(patchBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("update status = %d", resp.StatusCode)
	}

	req, err = http.NewRequest(http.MethodDelete, h.Server.URL+"/bookmarks/"+strconv.Itoa(created.ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("delete status = %d", resp.StatusCode)
	}

	resp, err = http.Get(h.Server.URL + "/bookmarks")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("list status = %d", resp.StatusCode)
	}
}

func TestIntegration_ForbiddenWhenNotOwner(t *testing.T) {
	h := env.SetupHTTP(t)

	_, ownerToken := fixtures.RegisterUser(t, h.Auth, "owner@example.com", "password123")
	_, otherToken := fixtures.RegisterUser(t, h.Auth, "other@example.com", "password123")

	createBody := `{"title":"Mine","url":"https://mine.dev"}`
	req, _ := http.NewRequest(http.MethodPost, h.Server.URL+"/bookmarks", strings.NewReader(createBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ownerToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	var created struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	patchBody := `{"title":"Stolen"}`
	req, err = http.NewRequest(http.MethodPatch, h.Server.URL+"/bookmarks/"+strconv.Itoa(created.ID), strings.NewReader(patchBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+otherToken)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("forbidden status = %d, want 403", resp.StatusCode)
	}
}

func TestIntegration_ParallelHTTP(t *testing.T) {
	t.Parallel()
	h := env.SetupHTTP(t)

	_, token := fixtures.RegisterUser(t, h.Auth, "parallel@example.com", "password123")
	req, err := http.NewRequest(http.MethodGet, h.Server.URL+"/bookmarks", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("list status = %d", resp.StatusCode)
	}
}
