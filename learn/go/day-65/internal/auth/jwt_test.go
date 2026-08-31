package auth_test

import (
	"testing"
	"time"

	"learn/go/day-65/internal/auth"
)

func TestIssueAndParseJWT(t *testing.T) {
	svc, err := auth.NewTokenService("test-secret-key-at-least-32-bytes-long", time.Hour, nil)
	if err != nil {
		t.Fatal(err)
	}

	token, err := svc.Issue(1, "a@b.com", auth.RoleMember)
	if err != nil {
		t.Fatal(err)
	}

	claims, err := svc.Parse(token)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if claims.UserID != 1 || claims.Email != "a@b.com" || claims.Role != auth.RoleMember {
		t.Fatalf("claims = %+v", claims)
	}

	_, err = svc.Parse("not.a.jwt")
	if err != auth.ErrInvalidToken {
		t.Fatalf("bad token err = %v", err)
	}
}
