package db

import "database/sql"

// TestDB wraps *sql.DB with helpers for integration tests.
type TestDB struct {
	DB   *sql.DB
	Path string
}

func (t *TestDB) Exec(query string, args ...any) (sql.Result, error) {
	return t.DB.Exec(query, args...)
}

func (t *TestDB) Close() error {
	return t.DB.Close()
}
