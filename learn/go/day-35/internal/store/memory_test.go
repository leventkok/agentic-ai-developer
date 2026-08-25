package store

import (
	"testing"

	"learn/go/day-35/internal/model"
)

func TestMemoryStore_ReturnsCopies(t *testing.T) {
	// TODO: create bookmark, mutate returned slice from List(), verify store unchanged
	s := NewMemoryStore()
	s.Create(model.CreateBookmarkRequest{Title: "Go", URL: "https://go.dev", Tags: []string{"lang"}})

	list := s.List()
	list[0].Tags[0] = "mutated"

	got, _ := s.Get(1)
	if got.Tags[0] == "mutated" {
		t.Fatal("store leaked internal slice reference")
	}
}

func TestMemoryStore_ConcurrentCreate(t *testing.T) {
	// TODO: run several goroutines calling Create, expect unique IDs
	s := NewMemoryStore()
	done := make(chan struct{}, 10)
	for range 10 {
		go func() {
			s.Create(model.CreateBookmarkRequest{Title: "x", URL: "https://x.dev"})
			done <- struct{}{}
		}()
	}
	for range 10 {
		<-done
	}
	if len(s.List()) != 10 {
		t.Fatalf("got %d bookmarks, want 10", len(s.List()))
	}
}
