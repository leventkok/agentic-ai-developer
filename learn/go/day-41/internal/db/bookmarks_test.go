package db_test

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"

	"learn/go/day-41/internal/db"
)

func migrationFile(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Join(filepath.Dir(file), "..", "..", "migrations", "001_init.sql")
}

func TestOpenAndListBookmarks(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")

	database, err := db.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()

	if err := db.Migrate(database, migrationFile(t)); err != nil {
		t.Fatal(err)
	}

	list, err := db.ListBookmarks(context.Background(), database)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) < 2 {
		t.Fatalf("got %d bookmarks, want at least 2", len(list))
	}
}
