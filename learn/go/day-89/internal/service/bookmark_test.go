package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"learn/go/day-89/internal/domain"
	"learn/go/day-89/internal/service"
	"learn/go/day-89/internal/service/testing/fake"
)

func TestBookmarkService_Create_Succeeds(t *testing.T) {
	svc := service.NewBookmarkService(fake.NewBookmarks(), time.Second, nil, nil)
	got, err := svc.Create(context.Background(), domain.User{ID: 1, Role: domain.RoleMember}, domain.CreateBookmarkInput{
		Title: "Go", URL: "https://go.dev", Tags: []string{"lang"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got.Title != "Go" {
		t.Fatalf("title = %q", got.Title)
	}
}

func TestBookmarkService_List_ReturnsBookmarks(t *testing.T) {
	repo := fake.NewBookmarks(domain.Bookmark{ID: 1, Title: "T", URL: "https://a.dev"})
	svc := service.NewBookmarkService(repo, time.Second, nil, nil)
	list, err := svc.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 1 {
		t.Fatalf("len = %d", len(list))
	}
}

func TestBookmarkService_Update_ForbiddenWhenNotOwner(t *testing.T) {
	ownerID := 1
	repo := fake.NewBookmarks(domain.Bookmark{
		ID: 1, UserID: &ownerID, Title: "T", URL: "https://a.dev",
	})
	svc := service.NewBookmarkService(repo, time.Second, nil, nil)

	title := "Hacked"
	_, err := svc.Update(context.Background(), domain.User{ID: 2, Role: domain.RoleMember}, 1, domain.UpdateBookmarkInput{Title: &title})
	if !errors.Is(err, domain.ErrForbidden) {
		t.Fatalf("got %v, want ErrForbidden", err)
	}
}

func TestBookmarkService_Update_SucceedsForOwner(t *testing.T) {
	ownerID := 1
	repo := fake.NewBookmarks(domain.Bookmark{
		ID: 1, UserID: &ownerID, Title: "T", URL: "https://a.dev",
	})
	svc := service.NewBookmarkService(repo, time.Second, nil, nil)

	title := "Updated"
	got, err := svc.Update(context.Background(), domain.User{ID: 1, Role: domain.RoleMember}, 1, domain.UpdateBookmarkInput{Title: &title})
	if err != nil {
		t.Fatal(err)
	}
	if got.Title != "Updated" {
		t.Fatalf("title = %q", got.Title)
	}
}

func TestBookmarkService_Delete_AdminCanDeleteAny(t *testing.T) {
	ownerID := 2
	repo := fake.NewBookmarks(domain.Bookmark{
		ID: 1, UserID: &ownerID, Title: "T", URL: "https://a.dev",
	})
	svc := service.NewBookmarkService(repo, time.Second, nil, nil)

	if err := svc.Delete(context.Background(), domain.User{ID: 99, Role: domain.RoleAdmin}, 1); err != nil {
		t.Fatal(err)
	}
}

func TestBookmarkService_List_RespectsTimeout(t *testing.T) {
	repo := fake.NewBookmarks(domain.Bookmark{ID: 1, Title: "T", URL: "https://a.dev"})
	repo.SetListDelay(200 * time.Millisecond)
	svc := service.NewBookmarkService(repo, 50*time.Millisecond, nil, nil)

	_, err := svc.List(context.Background())
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("got %v, want deadline exceeded", err)
	}
}

func TestBookmarkService_Create_ValidationError(t *testing.T) {
	svc := service.NewBookmarkService(fake.NewBookmarks(), time.Second, nil, nil)
	_, err := svc.Create(context.Background(), domain.User{ID: 1}, domain.CreateBookmarkInput{
		Title: "", URL: "https://go.dev",
	})
	if !domain.IsValidation(err) {
		t.Fatalf("got %v, want validation error", err)
	}
}
