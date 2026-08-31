package domain_test

import (
	"errors"
	"strings"
	"testing"

	"learn/go/day-88/internal/domain"
)

func TestNewTitle_RejectsEmpty(t *testing.T) {
	_, err := domain.NewTitle("")
	if err == nil {
		t.Fatal("expected error")
	}
	var ve domain.ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("expected ValidationError, got %v", err)
	}
}

func TestNewTitle_RejectsTooLong(t *testing.T) {
	_, err := domain.NewTitle(strings.Repeat("a", domain.MaxTitleLen+1))
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestNewTitle_AcceptsValid(t *testing.T) {
	title, err := domain.NewTitle("Go docs")
	if err != nil {
		t.Fatal(err)
	}
	if title.String() != "Go docs" {
		t.Fatalf("got %q", title.String())
	}
}

func TestNewBookmarkURL_RejectsNonHTTP(t *testing.T) {
	_, err := domain.NewBookmarkURL("ftp://example.com")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestNewBookmarkURL_AcceptsHTTPS(t *testing.T) {
	u, err := domain.NewBookmarkURL("https://go.dev")
	if err != nil {
		t.Fatal(err)
	}
	if u.String() != "https://go.dev" {
		t.Fatalf("got %q", u.String())
	}
}

func TestValidateCreateInput_EnforcesInvariants(t *testing.T) {
	_, err := domain.ValidateCreateInput(domain.CreateBookmarkInput{
		Title: "",
		URL:   "https://go.dev",
	})
	if err == nil {
		t.Fatal("expected validation error")
	}
}
