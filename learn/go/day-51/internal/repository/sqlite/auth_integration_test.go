package sqlite_test

import (
	"context"
	"testing"
	"time"

	"learn/go/day-51/internal/db/testutil"
	"learn/go/day-51/internal/repository"
	"learn/go/day-51/internal/repository/sqlite"
)

func TestAuthStore_RegisterLoginAndSession(t *testing.T) {
	tdb := testutil.OpenTestDB(t)
	testutil.ResetTables(t, tdb)

	store := sqlite.NewAuthStore(tdb.DB, 24*time.Hour)
	ctx := context.Background()

	user, err := store.Register(ctx, "alice@example.com", "password123")
	if err != nil {
		t.Fatalf("register: %v", err)
	}
	if user.Email != "alice@example.com" {
		t.Fatalf("email = %q", user.Email)
	}

	_, err = store.Register(ctx, "alice@example.com", "other-pass")
	if err != repository.ErrDuplicateEmail {
		t.Fatalf("duplicate register err = %v", err)
	}

	resp, err := store.Login(ctx, "alice@example.com", "password123")
	if err != nil {
		t.Fatalf("login: %v", err)
	}
	if resp.Token == "" {
		t.Fatal("expected session token")
	}

	fromToken, err := store.UserFromToken(ctx, resp.Token)
	if err != nil {
		t.Fatalf("user from token: %v", err)
	}
	if fromToken.ID != user.ID {
		t.Fatalf("user id = %d, want %d", fromToken.ID, user.ID)
	}

	_, err = store.Login(ctx, "alice@example.com", "wrong-password")
	if err != repository.ErrInvalidCredentials {
		t.Fatalf("bad password err = %v", err)
	}

	_, err = store.UserFromToken(ctx, "not-a-real-token")
	if err != repository.ErrUnauthorized {
		t.Fatalf("bad token err = %v", err)
	}
}
