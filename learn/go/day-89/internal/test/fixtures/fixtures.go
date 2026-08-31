package fixtures

import (
	"context"
	"testing"

	"learn/go/day-89/internal/repository"
	"learn/go/day-89/internal/domain"
)

// RegisterUser registers a test user and returns the user plus JWT token.
func RegisterUser(t *testing.T, auth repository.Auth, email, password string) (domain.User, string) {
	t.Helper()
	ctx := context.Background()
	user, err := auth.Register(ctx, email, password)
	if err != nil {
		t.Fatalf("register user: %v", err)
	}
	resp, err := auth.Login(ctx, email, password)
	if err != nil {
		t.Fatalf("login user: %v", err)
	}
	return user, resp.Token
}
