package testutil

import (
	"path/filepath"
	"runtime"
	"testing"

	"learn/go/day-80/internal/db"
)

func MigrationsDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Join(filepath.Dir(file), "..", "..", "..", "migrations")
}

// OpenTestDB returns a migrated SQLite database in a temp file.
func OpenTestDB(t *testing.T) *db.TestDB {
	t.Helper()
	path := filepath.Join(t.TempDir(), "test.db")
	database, err := db.Open(path, db.DefaultPoolConfig())
	if err != nil {
		t.Fatal(err)
	}
	if err := db.RunMigrations(database, MigrationsDir(t)); err != nil {
		database.Close()
		t.Fatal(err)
	}
	tdb := &db.TestDB{DB: database, Path: path}
	t.Cleanup(func() { tdb.Close() })
	return tdb
}

// ResetTables truncates user data but keeps schema and migrations applied.
func ResetTables(t *testing.T, database *db.TestDB) {
	t.Helper()
	_, err := database.Exec(`
		DELETE FROM bookmark_audit;
		DELETE FROM bookmarks;
		DELETE FROM sessions;
		DELETE FROM users;
		DELETE FROM sqlite_sequence WHERE name IN ('bookmarks', 'bookmark_audit', 'users');
	`)
	if err != nil {
		t.Fatalf("reset tables: %v", err)
	}
}
