package store

import (
	"context"
	"errors"
	"testing"
	"time"

	"learn/go/day-39/internal/model"
)

func seedStore(t *testing.T, n int) *MemoryStore {
	t.Helper()
	s := NewMemoryStore()
	for range n {
		_, err := s.Create(context.Background(), model.CreateBookmarkRequest{
			Title: "Bookmark",
			URL:   "https://example.com",
			Tags:  []string{"tag"},
		})
		if err != nil {
			t.Fatalf("seed create failed: %v", err)
		}
	}
	return s
}

func TestList_WithCancel(t *testing.T) {
	s := seedStore(t, 20)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan error, 1)
	go func() {
		_, err := s.List(ctx)
		done <- err
	}()

	// List sleeps 10ms per item — after ~15ms one item processed, then we cancel.
	time.Sleep(15 * time.Millisecond)
	cancel()

	err := <-done
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("got err=%v, want context.Canceled", err)
	}
}

func TestList_WithTimeout(t *testing.T) {
	s := seedStore(t, 20) // 20 × 10ms ≈ 200ms if uncanceled

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	_, err := s.List(ctx)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("got err=%v, want context.DeadlineExceeded", err)
	}
}
