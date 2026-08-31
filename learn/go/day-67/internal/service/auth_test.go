package service_test

import (
	"context"
	"testing"

	"learn/go/day-67/internal/service"
	"learn/go/day-67/internal/service/testing/fake"
)

func TestAuthService_Register_DelegatesToRepo(t *testing.T) {
	repo := fake.NewAuth()
	svc := service.NewAuthService(repo)

	_, err := svc.Register(context.Background(), "a@b.com", "password123")
	if err != nil {
		t.Fatal(err)
	}
	if repo.RegisterCalls() != 1 {
		t.Fatalf("register calls = %d, want 1", repo.RegisterCalls())
	}
}

func TestAuthService_Login_ReturnsTokenFromRepo(t *testing.T) {
	repo := fake.NewAuth()
	svc := service.NewAuthService(repo)

	resp, err := svc.Login(context.Background(), "a@b.com", "password123")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Token != "fake-token" {
		t.Fatalf("token = %q", resp.Token)
	}
}
