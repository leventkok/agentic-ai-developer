package outbox

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"learn/go/day-92/internal/messaging"
)

type Record struct {
	ID        int64
	EventID   string
	EventType messaging.EventType
	Payload   []byte
	CreatedAt time.Time
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Enqueue(ctx context.Context, evt messaging.Event) error {
	payload, err := evt.Marshal()
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx, `
		INSERT INTO outbox (event_id, event_type, payload)
		VALUES (?, ?, ?)`,
		evt.DeduplicationKey(), string(evt.Type), payload,
	)
	if err != nil {
		return fmt.Errorf("outbox enqueue: %w", err)
	}
	return nil
}

func (s *Store) FetchUnpublished(ctx context.Context, limit int) ([]Record, error) {
	if limit <= 0 {
		limit = 10
	}
	rows, err := s.db.QueryContext(ctx, `
		SELECT id, event_id, event_type, payload, created_at
		FROM outbox
		WHERE published_at IS NULL
		ORDER BY id ASC
		LIMIT ?`, limit)
	if err != nil {
		return nil, fmt.Errorf("outbox fetch: %w", err)
	}
	defer rows.Close()

	var records []Record
	for rows.Next() {
		var rec Record
		var createdAt string
		var eventType string
		if err := rows.Scan(&rec.ID, &rec.EventID, &eventType, &rec.Payload, &createdAt); err != nil {
			return nil, fmt.Errorf("outbox scan: %w", err)
		}
		rec.EventType = messaging.EventType(eventType)
		rec.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		records = append(records, rec)
	}
	return records, rows.Err()
}

func (s *Store) MarkPublished(ctx context.Context, id int64) error {
	_, err := s.db.ExecContext(ctx, `
		UPDATE outbox SET published_at = datetime('now') WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("outbox mark published: %w", err)
	}
	return nil
}
