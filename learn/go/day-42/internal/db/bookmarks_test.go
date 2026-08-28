package db_test

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"

	"learn/go/day-42/internal/db"
)

func migrationsDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Join(filepath.Dir(file), "..", "..", "migrations")
}

func TestRunMigrations(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	database, err := db.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()

	if err := db.RunMigrations(database, migrationsDir(t)); err != nil {
		t.Fatal(err)
	}

	list, err := db.ListBookmarks(context.Background(), database)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) < 2 {
		t.Fatalf("got %d bookmarks, want at least 2", len(list))
	}

	// running again should be idempotent
	if err := db.RunMigrations(database, migrationsDir(t)); err != nil {
		t.Fatal(err)
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

	list, err := db.ListBookmarks(context.Background(), database)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 0 {
		t.Fatalf("after rollback seed, got %d bookmarks, want 0", len(list))
	}
}
