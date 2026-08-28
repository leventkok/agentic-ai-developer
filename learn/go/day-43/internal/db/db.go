package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	db.SetMaxOpenConns(1)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return db, nil
}

func ensureMigrationTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
		version INTEGER PRIMARY KEY,
		applied_at TEXT NOT NULL DEFAULT (datetime('now'))
	)`)
	return err
}

func appliedVersions(db *sql.DB) (map[int]bool, error) {
	rows, err := db.Query(`SELECT version FROM schema_migrations`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	versions := make(map[int]bool)
	for rows.Next() {
		var v int
		if err := rows.Scan(&v); err != nil {
			return nil, err
		}
		versions[v] = true
	}
	return versions, rows.Err()
}

func migrationFiles(dir string, suffix string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasSuffix(name, suffix) {
			files = append(files, filepath.Join(dir, name))
		}
	}
	sort.Strings(files)
	return files, nil
}

func parseVersion(filename string) (int, error) {
	base := filepath.Base(filename)
	part := strings.SplitN(base, "_", 2)[0]
	return strconv.Atoi(part)
}

// RunMigrations applies pending *.up.sql files in order.
func RunMigrations(db *sql.DB, dir string) error {
	if err := ensureMigrationTable(db); err != nil {
		return err
	}

	applied, err := appliedVersions(db)
	if err != nil {
		return err
	}

	files, err := migrationFiles(dir, ".up.sql")
	if err != nil {
		return err
	}

	for _, file := range files {
		version, err := parseVersion(file)
		if err != nil {
			return fmt.Errorf("parse version %s: %w", file, err)
		}
		if applied[version] {
			continue
		}

		raw, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		if _, err := db.Exec(string(raw)); err != nil {
			return fmt.Errorf("migration %d up failed: %w", version, err)
		}

		if _, err := db.Exec(`INSERT INTO schema_migrations (version) VALUES (?)`, version); err != nil {
			return fmt.Errorf("record migration %d: %w", version, err)
		}
	}

	return nil
}

// RollbackMigration reverts the latest applied migration using *.down.sql.
func RollbackMigration(db *sql.DB, dir string) error {
	if err := ensureMigrationTable(db); err != nil {
		return err
	}

	var version int
	err := db.QueryRow(`SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1`).Scan(&version)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return err
	}

	matches, err := filepath.Glob(filepath.Join(dir, fmt.Sprintf("%03d_*.down.sql", version)))
	if err != nil {
		return err
	}
	if len(matches) == 0 {
		return fmt.Errorf("no down migration for version %d", version)
	}

	raw, err := os.ReadFile(matches[0])
	if err != nil {
		return err
	}

	if _, err := db.Exec(string(raw)); err != nil {
		return fmt.Errorf("migration %d down failed: %w", version, err)
	}

	_, err = db.Exec(`DELETE FROM schema_migrations WHERE version = ?`, version)
	return err
}
