package store_test

import (
	"context"
	"database/sql"
	"errors"
	"path/filepath"
	"runtime"
	"testing"

	"learn/go/day-44/internal/db"
	"learn/go/day-44/internal/model"
	"learn/go/day-44/internal/store"
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

func TestDBStore_BulkCreate(t *testing.T) {
	s := newTestDBStore(t)
	ctx := context.Background()

	before, err := s.List(ctx)
	if err != nil {
		t.Fatal(err)
	}

	created, err := s.BulkCreate(ctx, []model.CreateBookmarkRequest{
		{Title: "One", URL: "https://one.dev"},
		{Title: "Two", URL: "https://two.dev"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(created) != 2 {
		t.Fatalf("created %d bookmarks, want 2", len(created))
	}

	after, err := s.List(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(after) != len(before)+2 {
		t.Fatalf("count = %d, want %d", len(after), len(before)+2)
	}
}

func TestDBStore_BulkCreateRollsBackOnFailure(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	database, err := db.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()

	if err := db.RunMigrations(database, migrationsDir(t)); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	var before int
	if err := database.QueryRow(`SELECT COUNT(*) FROM bookmarks`).Scan(&before); err != nil {
		t.Fatal(err)
	}

	boom := errors.New("boom")
	err = db.RunInTx(ctx, database, func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			INSERT INTO bookmarks (title, url, tags, created_at, updated_at)
			VALUES (?, ?, '[]', datetime('now'), datetime('now'))`,
			"Partial", "https://partial.dev"); err != nil {
			return err
		}
		return boom
	})
	if !errors.Is(err, boom) {
		t.Fatalf("got err=%v, want boom", err)
	}

	var after int
	if err := database.QueryRow(`SELECT COUNT(*) FROM bookmarks`).Scan(&after); err != nil {
		t.Fatal(err)
	}
	if after != before {
		t.Fatalf("count changed from %d to %d after failed tx", before, after)
	}
}

func TestDBStore_UpdateWritesAuditRow(t *testing.T) {
	s := newTestDBStore(t)
	ctx := context.Background()

	created, err := s.Create(ctx, model.CreateBookmarkRequest{
		Title: "Audit me",
		URL:   "https://audit.dev",
	})
	if err != nil {
		t.Fatal(err)
	}

	title := "Audited"
	if _, err := s.Update(ctx, created.ID, model.UpdateBookmarkRequest{Title: &title}); err != nil {
		t.Fatal(err)
	}

	count, err := s.AuditCount(ctx, created.ID)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("audit count = %d, want 1", count)
	}
}
