package resilient_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"learn/go/day-93/internal/domain"
	"learn/go/day-93/internal/repository/resilient"
	"learn/go/day-93/internal/service/testing/fake"
)

func TestResilientBookmarks_List_RetriesTransientErrors(t *testing.T) {
	repo := fake.NewBookmarks(domain.Bookmark{ID: 1, Title: "T", URL: "https://a.dev"})
	attempts := 0
	repo.SetListHook(func(context.Context) ([]domain.Bookmark, error) {
		attempts++
		if attempts < 2 {
			return nil, errors.New("transient")
		}
		return []domain.Bookmark{{ID: 1, Title: "T", URL: "https://a.dev"}}, nil
	})

	wrapped := resilient.NewBookmarks(repo)
	list, err := wrapped.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 1 {
		t.Fatalf("len = %d", len(list))
	}
	if attempts < 2 {
		t.Fatalf("attempts = %d", attempts)
	}
}

func TestResilientBookmarks_List_RespectsContextTimeout(t *testing.T) {
	repo := fake.NewBookmarks()
	repo.SetListDelay(200 * time.Millisecond)
	wrapped := resilient.NewBookmarks(repo)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	_, err := wrapped.List(ctx)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("got %v", err)
	}
}
