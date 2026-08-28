package db_test

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"

	"learn/go/day-45/internal/db"
	"learn/go/day-45/internal/store"
)

func migrationsDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Join(filepath.Dir(file), "..", "..", "migrations")
}

func openMigratedStore(t *testing.T) *store.DBStore {
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

func TestRunMigrations(t *testing.T) {
	s := openMigratedStore(t)

	list, err := s.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(list) < 2 {
		t.Fatalf("got %d bookmarks, want at least 2", len(list))
	}
}

func TestRollbackMigration(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	database, err := db.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()

	dir := migrationsDir(t)
	if err := db.RunMigrations(database, dir); err != nil {
		t.Fatal(err)
	}

	if err := db.RollbackMigration(database, dir); err != nil {
		t.Fatal(err)
	}

	var auditTableCount int
	if err := database.QueryRow(`
		SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='bookmark_audit'`).Scan(&auditTableCount); err != nil {
		t.Fatal(err)
	}
	if auditTableCount != 0 {
		t.Fatal("expected bookmark_audit table to be dropped after rollback")
	}

	if err := db.RollbackMigration(database, dir); err != nil {
		t.Fatal(err)
	}

	s, err := store.NewDBStore(database)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	list, err := s.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 0 {
		t.Fatalf("after rollback seed, got %d bookmarks, want 0", len(list))
	}
}
