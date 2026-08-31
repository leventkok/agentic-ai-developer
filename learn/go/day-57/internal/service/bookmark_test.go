package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"learn/go/day-57/internal/domain"
	"learn/go/day-57/internal/repository/memory"
	"learn/go/day-57/internal/service"
)

func TestBookmarkService_Update_ForbiddenWhenNotOwner(t *testing.T) {
	store := memory.New()
	ctx := context.Background()
	owner := domain.User{ID: 1, Role: domain.RoleMember}
	created, err := store.Create(ctx, owner, domain.CreateBookmarkInput{Title: "T", URL: "https://a.dev"})
	if err != nil {
		t.Fatal(err)
	}

	svc := service.NewBookmarkService(store, time.Second)
	title := "Hacked"
	_, err = svc.Update(ctx, domain.User{ID: 2, Role: domain.RoleMember}, created.ID, domain.UpdateBookmarkInput{Title: &title})
	if !errors.Is(err, domain.ErrForbidden) {
		t.Fatalf("got %v, want ErrForbidden", err)
	}
}

func TestBookmarkService_Delete_AdminCanDeleteAny(t *testing.T) {
	store := memory.New()
	ctx := context.Background()
	owner := domain.User{ID: 2, Role: domain.RoleMember}
	created, err := store.Create(ctx, owner, domain.CreateBookmarkInput{Title: "T", URL: "https://a.dev"})
	if err != nil {
		t.Fatal(err)
	}

	svc := service.NewBookmarkService(store, time.Second)
	if err := svc.Delete(ctx, domain.User{ID: 99, Role: domain.RoleAdmin}, created.ID); err != nil {
		t.Fatal(err)
	}
}
