package store

import (
	"context"
	"testing"

	"learn/go/day-44/internal/model"
)

func TestMemoryStore_ReturnsCopies(t *testing.T) {
	s := NewMemoryStore()
	_, err := s.Create(context.Background(), model.CreateBookmarkRequest{
		Title: "Go",
		URL:   "https://go.dev",
		Tags:  []string{"lang"},
	})
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	list, err := s.List(context.Background())
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	list[0].Tags[0] = "mutated"

	got, err := s.Get(context.Background(), 1)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if got.Tags[0] == "mutated" {
		t.Fatal("store leaked internal slice reference")
	}
}

func TestMemoryStore_ConcurrentCreate(t *testing.T) {
	s := NewMemoryStore()
	done := make(chan struct{}, 10)
	for range 10 {
		go func() {
			_, _ = s.Create(context.Background(), model.CreateBookmarkRequest{
				Title: "x",
				URL:   "https://x.dev",
			})
			done <- struct{}{}
		}()
	}
	for range 10 {
		<-done
	}

	list, err := s.List(context.Background())
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(list) != 10 {
		t.Fatalf("got %d bookmarks, want 10", len(list))
	}
}
