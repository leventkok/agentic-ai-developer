package sqlite_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"learn/go/day-90/internal/db"
	"learn/go/day-90/internal/db/testutil"
	"learn/go/day-90/internal/domain"
	"learn/go/day-90/internal/repository/sqlite"
)

func testMember(id int) domain.User {
	return domain.User{ID: id, Role: domain.RoleMember}
}

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
	actor := testMember(1)

	created, err := repo.Create(ctx, actor, domain.CreateBookmarkInput{
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

	if err := repo.Delete(ctx, actor, created.ID); err != nil {
		t.Fatal(err)
	}
	_, err = repo.Get(ctx, created.ID)
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("got err=%v, want ErrNotFound", err)
	}
}

func TestStore_UpdateWritesAudit(t *testing.T) {
	repo := newRepo(t)
	ctx := context.Background()
	actor := testMember(1)

	created, err := repo.Create(ctx, actor, domain.CreateBookmarkInput{
		Title: "Audit",
		URL:   "https://audit.dev",
	})
	if err != nil {
		t.Fatal(err)
	}

	title := "Updated"
	if _, err := repo.Update(ctx, actor, created.ID, domain.UpdateBookmarkInput{Title: &title}); err != nil {
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
	actor := testMember(1)

	created, err := repo.BulkCreate(ctx, actor, []domain.CreateBookmarkInput{
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
			INSERT INTO bookmarks (user_id, title, url, tags, created_at, updated_at)
			VALUES (1, 'Partial', 'https://partial.dev', '[]', datetime('now'), datetime('now'))`)
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
	if _, err := repo1.Create(ctx, testMember(1), domain.CreateBookmarkInput{Title: "One", URL: "https://one.dev"}); err != nil {
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

func TestStore_ForbiddenCrossUser(t *testing.T) {
	t.Skip("ownership enforced in service layer, not repository")
}
