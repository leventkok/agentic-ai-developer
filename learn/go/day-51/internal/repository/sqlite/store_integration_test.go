package sqlite_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"learn/go/day-51/internal/db"
	"learn/go/day-51/internal/db/testutil"
	"learn/go/day-51/internal/model"
	"learn/go/day-51/internal/repository"
	"learn/go/day-51/internal/repository/sqlite"
)

func newRepo(t *testing.T) *sqlite.Store {
	t.Helper()
	tdb := testutil.OpenTestDB(t)
	testutil.ResetTables(t, tdb)
	repo, err := sqlite.New(tdb.DB)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { repo.Close() })
	return repo
}

func TestStore_CreateGetDelete(t *testing.T) {
	repo := newRepo(t)
	ctx := context.Background()

	created, err := repo.Create(ctx, model.CreateBookmarkRequest{
		Title: "Go Blog",
		URL:   "https://go.dev/blog",
		Tags:  []string{"go"},
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.Get(ctx, created.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Title != "Go Blog" {
		t.Fatalf("title = %q", got.Title)
	}

	if err := repo.Delete(ctx, created.ID); err != nil {
		t.Fatal(err)
	}
	_, err = repo.Get(ctx, created.ID)
	if !errors.Is(err, repository.ErrNotFound) {
		t.Fatalf("got err=%v, want ErrNotFound", err)
	}
}

func TestStore_UpdateWritesAudit(t *testing.T) {
	repo := newRepo(t)
	ctx := context.Background()

	created, err := repo.Create(ctx, model.CreateBookmarkRequest{
		Title: "Audit",
		URL:   "https://audit.dev",
	})
	if err != nil {
		t.Fatal(err)
	}

	title := "Updated"
	if _, err := repo.Update(ctx, created.ID, model.UpdateBookmarkRequest{Title: &title}); err != nil {
		t.Fatal(err)
	}

	count, err := repo.AuditCount(ctx, created.ID)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("audit count = %d, want 1", count)
	}
}

func TestStore_BulkCreate(t *testing.T) {
	repo := newRepo(t)
	ctx := context.Background()

	created, err := repo.BulkCreate(ctx, []model.CreateBookmarkRequest{
		{Title: "A", URL: "https://a.dev"},
		{Title: "B", URL: "https://b.dev"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(created) != 2 {
		t.Fatalf("got %d, want 2", len(created))
	}
}

func TestStore_TransactionRollback(t *testing.T) {
	tdb := testutil.OpenTestDB(t)
	testutil.ResetTables(t, tdb)
	ctx := context.Background()

	boom := errors.New("boom")
	err := db.RunInTx(ctx, tdb.DB, func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			INSERT INTO bookmarks (title, url, tags, created_at, updated_at)
			VALUES ('Partial', 'https://partial.dev', '[]', datetime('now'), datetime('now'))`)
		if err != nil {
			return err
		}
		return boom
	})
	if !errors.Is(err, boom) {
		t.Fatalf("got err=%v", err)
	}

	var count int
	if err := tdb.DB.QueryRow(`SELECT COUNT(*) FROM bookmarks`).Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatalf("count = %d, want 0 after rollback", count)
	}
}

func TestStore_ResetIsolation(t *testing.T) {
	repo1 := newRepo(t)
	ctx := context.Background()
	if _, err := repo1.Create(ctx, model.CreateBookmarkRequest{Title: "One", URL: "https://one.dev"}); err != nil {
		t.Fatal(err)
	}

	repo2 := newRepo(t)
	list, err := repo2.List(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 0 {
		t.Fatalf("expected empty after reset, got %d", len(list))
	}
}
