package sqlite_test

import (
	"context"
	"testing"
	"time"

	"learn/go/day-69/internal/auth"
	"learn/go/day-69/internal/db/testutil"
	"learn/go/day-69/internal/domain"
	"learn/go/day-69/internal/repository/sqlite"
)

func TestAuthStore_RegisterLoginJWT(t *testing.T) {
	tdb := testutil.OpenTestDB(t)
	testutil.ResetTables(t, tdb)

	tokens, err := auth.NewTokenService("test-secret-key-at-least-32-bytes-long", 24*time.Hour, nil)
	if err != nil {
		t.Fatal(err)
	}
	store := sqlite.NewAuthStore(tdb.DB, tokens)
	ctx := context.Background()

	user, err := store.Register(ctx, "alice@example.com", "password123")
	if err != nil {
		t.Fatalf("register: %v", err)
	}
	if user.Role != domain.RoleMember {
		t.Fatalf("role = %q", user.Role)
	}

	resp, err := store.Login(ctx, "alice@example.com", "password123")
	if err != nil {
		t.Fatalf("login: %v", err)
	}
	if resp.Token == "" {
		t.Fatal("expected jwt token")
	}

	fromToken, err := store.UserFromToken(ctx, resp.Token)
	if err != nil {
		t.Fatalf("user from token: %v", err)
	}
	if fromToken.ID != user.ID {
		t.Fatalf("user id = %d, want %d", fromToken.ID, user.ID)
	}

	_, err = store.Login(ctx, "alice@example.com", "wrong-password")
	if err != domain.ErrInvalidCredentials {
		t.Fatalf("bad password err = %v", err)
	}

	_, err = store.UserFromToken(ctx, "not.a.jwt")
	if err != domain.ErrUnauthorized {
		t.Fatalf("bad token err = %v", err)
	}
}
