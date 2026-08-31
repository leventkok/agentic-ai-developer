package cached_test

import (
	"context"
	"testing"
	"time"

	"learn/go/day-96/internal/cache/memory"
	"learn/go/day-96/internal/domain"
	"learn/go/day-96/internal/repository/cached"
	"learn/go/day-96/internal/service/testing/fake"
)

func TestCachedBookmarks_ListUsesCacheAfterFirstLoad(t *testing.T) {
	repo := fake.NewBookmarks(domain.Bookmark{ID: 1, Title: "T", URL: "https://a.dev"})
	calls := 0
	repo.SetListHook(func(context.Context) ([]domain.Bookmark, error) {
		calls++
		return []domain.Bookmark{{ID: 1, Title: "T", URL: "https://a.dev"}}, nil
	})

	store := memory.New()
	wrapped := cached.New(repo, store, time.Minute)

	if _, err := wrapped.List(context.Background()); err != nil {
		t.Fatal(err)
	}
	if _, err := wrapped.List(context.Background()); err != nil {
		t.Fatal(err)
	}
	if calls != 1 {
		t.Fatalf("inner list calls = %d, want 1", calls)
	}
}

func TestCachedBookmarks_CreateInvalidatesList(t *testing.T) {
	repo := fake.NewBookmarks()
	store := memory.New()
	wrapped := cached.New(repo, store, time.Minute)

	calls := 0
	repo.SetListHook(func(context.Context) ([]domain.Bookmark, error) {
		calls++
		return []domain.Bookmark{}, nil
	})

	if _, err := wrapped.List(context.Background()); err != nil {
		t.Fatal(err)
	}
	if _, err := wrapped.Create(context.Background(), domain.User{ID: 1}, domain.CreateBookmarkInput{
		Title: "Go", URL: "https://go.dev",
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := wrapped.List(context.Background()); err != nil {
		t.Fatal(err)
	}
	if calls != 2 {
		t.Fatalf("list calls = %d, want 2 after invalidation", calls)
	}
}
