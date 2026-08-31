package idempotency

import (
	"context"
	"database/sql"
	"fmt"

	"learn/go/day-90/internal/messaging"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) TryMarkProcessed(ctx context.Context, dedupKey string) (bool, error) {
	res, err := s.db.ExecContext(ctx, `
		INSERT OR IGNORE INTO processed_events (dedup_key) VALUES (?)`, dedupKey)
	if err != nil {
		return false, fmt.Errorf("idempotency mark: %w", err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return n == 1, nil
}

func Wrap(store *Store, handler messaging.Handler) messaging.Handler {
	return func(ctx context.Context, evt messaging.Event) error {
		ok, err := store.TryMarkProcessed(ctx, evt.DeduplicationKey())
		if err != nil {
			return err
		}
		if !ok {
			return nil
		}
		return handler(ctx, evt)
	}
}
