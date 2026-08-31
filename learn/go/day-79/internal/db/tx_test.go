package db_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"learn/go/day-79/internal/db"
)

func TestRunInTx_CommitsOnSuccess(t *testing.T) {
	database, err := db.Open(":memory:", db.DefaultPoolConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()

	err = db.RunInTx(context.Background(), database, func(tx *sql.Tx) error {
		_, err := tx.Exec(`CREATE TABLE counters (value INTEGER NOT NULL)`)
		if err != nil {
			return err
		}
		_, err = tx.Exec(`INSERT INTO counters (value) VALUES (1)`)
		return err
	})
	if err != nil {
		t.Fatal(err)
	}

	var count int
	if err := database.QueryRow(`SELECT COUNT(*) FROM counters`).Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("count = %d, want 1", count)
	}
}

func TestRunInTx_RollbacksOnError(t *testing.T) {
	database, err := db.Open(":memory:", db.DefaultPoolConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer database.Close()

	boom := errors.New("boom")
	err = db.RunInTx(context.Background(), database, func(tx *sql.Tx) error {
		if _, err := tx.Exec(`CREATE TABLE counters (value INTEGER NOT NULL)`); err != nil {
			return err
		}
		return boom
	})
	if !errors.Is(err, boom) {
		t.Fatalf("got err=%v, want boom", err)
	}

	var count int
	err = database.QueryRow(`SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='counters'`).Scan(&count)
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatal("expected counters table to be rolled back")
	}
}
