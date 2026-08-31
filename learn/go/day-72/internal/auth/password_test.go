package auth_test

import (
	"testing"

	"learn/go/day-72/internal/auth"
)

func TestHashAndCheckPassword(t *testing.T) {
	hash, err := auth.HashPassword("secret-password")
	if err != nil {
		t.Fatal(err)
	}
	if hash == "secret-password" {
		t.Fatal("password stored in plaintext")
	}
	if err := auth.CheckPassword(hash, "secret-password"); err != nil {
		t.Fatalf("check password: %v", err)
	}
	if err := auth.CheckPassword(hash, "wrong"); err == nil {
		t.Fatal("expected wrong password to fail")
	}
}
