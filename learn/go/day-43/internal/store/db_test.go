package store_test

import (
	"context"
	"errors"
	"path/filepath"
	"runtime"
	"testing"

	"learn/go/day-43/internal/db"
	"learn/go/day-43/internal/model"
	"learn/go/day-43/internal/store"
)

func migrationsDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Join(filepath.Dir(file), "..", "..", "migrations")
}

func newTestDBStore(t *testing.T) *store.DBStore {
	t.Helper()
	path := filepath.Join(t.TempDir(), "test.db")
	database, err := db.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { database.Close() })

	if err := db.RunMigrations(database, migrationsDir(t)); err != nil {
		t.Fatal(err)
	}

	s, err := store.NewDBStore(database)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { s.Close() })
	return s
}

func TestDBStore_CreateAndGet(t *testing.T) {
	s := newTestDBStore(t)
	ctx := context.Background()

	created, err := s.Create(ctx, model.CreateBookmarkRequest{
		Title: "Go Blog",
		URL:   "https://go.dev/blog",
		Tags:  []string{"go"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if created.ID == 0 {
		t.Fatal("expected non-zero id")
	}

	got, err := s.Get(ctx, created.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Title != "Go Blog" {
		t.Fatalf("title = %q, want Go Blog", got.Title)
	}
}

func TestDBStore_GetNotFound(t *testing.T) {
	s := newTestDBStore(t)

	_, err := s.Get(context.Background(), 9999)
	if !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("got err=%v, want ErrNotFound", err)
	}
}

func TestDBStore_Update(t *testing.T) {
	s := newTestDBStore(t)
	ctx := context.Background()

	created, err := s.Create(ctx, model.CreateBookmarkRequest{
		Title: "Old",
		URL:   "https://old.dev",
	})
	if err != nil {
		t.Fatal(err)
	}

	title := "New"
	updated, err := s.Update(ctx, created.ID, model.UpdateBookmarkRequest{Title: &title})
	if err != nil {
		t.Fatal(err)
	}
	if updated.Title != "New" {
		t.Fatalf("title = %q, want New", updated.Title)
	}
}

func TestDBStore_Delete(t *testing.T) {
	s := newTestDBStore(t)
	ctx := context.Background()

	created, err := s.Create(ctx, model.CreateBookmarkRequest{
		Title: "Delete me",
		URL:   "https://delete.dev",
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := s.Delete(ctx, created.ID); err != nil {
		t.Fatal(err)
	}

	_, err = s.Get(ctx, created.ID)
	if !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("got err=%v, want ErrNotFound", err)
	}
}

func TestDBStore_DeleteNotFound(t *testing.T) {
	s := newTestDBStore(t)

	err := s.Delete(context.Background(), 9999)
	if !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("got err=%v, want ErrNotFound", err)
	}
}
