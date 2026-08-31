package sqlite_test

import (
	"context"
	"testing"
	"time"

	"learn/go/day-52/internal/auth"
	"learn/go/day-52/internal/db/testutil"
	"learn/go/day-52/internal/repository"
	"learn/go/day-52/internal/repository/sqlite"
)

func TestAuthStore_RegisterLoginJWT(t *testing.T) {
	tdb := testutil.OpenTestDB(t)
	testutil.ResetTables(t, tdb)

	tokens, err := auth.NewTokenService("test-secret-key-at-least-32-bytes-long", 24*time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	store := sqlite.NewAuthStore(tdb.DB, tokens)
	ctx := context.Background()

	user, err := store.Register(ctx, "alice@example.com", "password123")
	if err != nil {
		t.Fatalf("register: %v", err)
	}

	resp, err := store.Login(ctx, "alice@example.com", "password123")
	if err != nil {
		t.Fatalf("login: %v", err)
	}
	if resp.Token == "" || len(resp.Token) < 20 {
		t.Fatal("expected jwt token")
	}

	fromToken, err := store.UserFromToken(ctx, resp.Token)
	if err != nil {
		t.Fatalf("user from token: %v", err)
	}
	if fromToken.ID != user.ID {
		t.Fatalf("user id = %d, want %d", fromToken.ID, user.ID)
	}

	_, err = store.UserFromToken(ctx, "not.a.jwt")
	if err != repository.ErrUnauthorized {
		t.Fatalf("bad token err = %v", err)
	}
}
