package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"learn/go/day-45/internal/db"
	"learn/go/day-45/internal/model"
)

type DBStore struct {
	db         *sql.DB
	listStmt   *sql.Stmt
	getStmt    *sql.Stmt
	insertStmt *sql.Stmt
	updateStmt *sql.Stmt
	deleteStmt *sql.Stmt
}

func NewDBStore(database *sql.DB) (*DBStore, error) {
	s := &DBStore{db: database}

	var err error
	if s.listStmt, err = prepare(database, "list bookmarks", `
		SELECT id, title, url, tags, created_at, updated_at
		FROM bookmarks ORDER BY id`); err != nil {
		s.Close()
		return nil, err
	}
	if s.getStmt, err = prepare(database, "get bookmark", `
		SELECT id, title, url, tags, created_at, updated_at
		FROM bookmarks WHERE id = ?`); err != nil {
		s.Close()
		return nil, err
	}
	if s.insertStmt, err = prepare(database, "insert bookmark", `
		INSERT INTO bookmarks (title, url, tags, created_at, updated_at)
		VALUES (?, ?, ?, datetime('now'), datetime('now'))
		RETURNING id, title, url, tags, created_at, updated_at`); err != nil {
		s.Close()
		return nil, err
	}
	if s.updateStmt, err = prepare(database, "update bookmark", `
		UPDATE bookmarks
		SET title = ?, url = ?, tags = ?, updated_at = datetime('now')
		WHERE id = ?`); err != nil {
		s.Close()
		return nil, err
	}
	if s.deleteStmt, err = prepare(database, "delete bookmark", `
		DELETE FROM bookmarks WHERE id = ?`); err != nil {
		s.Close()
		return nil, err
	}

	return s, nil
}

func prepare(database *sql.DB, name, query string) (*sql.Stmt, error) {
	stmt, err := database.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare %s: %w", name, err)
	}
	return stmt, nil
}

func (s *DBStore) Close() error {
	var errs []error
	for _, stmt := range []*sql.Stmt{s.listStmt, s.getStmt, s.insertStmt, s.updateStmt, s.deleteStmt} {
		if stmt != nil {
			if err := stmt.Close(); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errors.Join(errs...)
}

func (s *DBStore) List(ctx context.Context) ([]model.Bookmark, error) {
	rows, err := s.listStmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("list bookmarks: %w", err)
	}
	defer rows.Close()

	var list []model.Bookmark
	for rows.Next() {
		b, err := scanBookmark(rows)
		if err != nil {
			return nil, fmt.Errorf("scan bookmark row: %w", err)
		}
		list = append(list, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate bookmarks: %w", err)
	}
	return list, nil
}

func (s *DBStore) Get(ctx context.Context, id int) (model.Bookmark, error) {
	row := s.getStmt.QueryRowContext(ctx, id)
	b, err := scanBookmark(row)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Bookmark{}, ErrNotFound
	}
	if err != nil {
		return model.Bookmark{}, fmt.Errorf("get bookmark id=%d: %w", id, err)
	}
	return b, nil
}

func (s *DBStore) Create(ctx context.Context, req model.CreateBookmarkRequest) (model.Bookmark, error) {
	tagsJSON, err := marshalTags(req.Tags)
	if err != nil {
		return model.Bookmark{}, fmt.Errorf("marshal tags: %w", err)
	}

	row := s.insertStmt.QueryRowContext(ctx, req.Title, req.URL, tagsJSON)
	b, err := scanBookmark(row)
	if err != nil {
		return model.Bookmark{}, fmt.Errorf("insert bookmark: %w", err)
	}
	return b, nil
}

// BulkCreate inserts every bookmark in one transaction — all succeed or none do.
func (s *DBStore) BulkCreate(ctx context.Context, reqs []model.CreateBookmarkRequest) ([]model.Bookmark, error) {
	if len(reqs) == 0 {
		return nil, nil
	}

	var created []model.Bookmark
	err := db.RunInTx(ctx, s.db, func(tx *sql.Tx) error {
		stmt, err := tx.PrepareContext(ctx, `
			INSERT INTO bookmarks (title, url, tags, created_at, updated_at)
			VALUES (?, ?, ?, datetime('now'), datetime('now'))
			RETURNING id, title, url, tags, created_at, updated_at`)
		if err != nil {
			return fmt.Errorf("prepare bulk insert: %w", err)
		}
		defer stmt.Close()

		created = make([]model.Bookmark, 0, len(reqs))
		for i, req := range reqs {
			tagsJSON, err := marshalTags(req.Tags)
			if err != nil {
				return fmt.Errorf("marshal tags item %d: %w", i, err)
			}

			row := stmt.QueryRowContext(ctx, req.Title, req.URL, tagsJSON)
			b, err := scanBookmark(row)
			if err != nil {
				return fmt.Errorf("insert bookmark item %d: %w", i, err)
			}
			created = append(created, b)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return created, nil
}

// Update changes a bookmark and writes an audit row in the same transaction.
func (s *DBStore) Update(ctx context.Context, id int, req model.UpdateBookmarkRequest) (model.Bookmark, error) {
	var updated model.Bookmark

	err := db.RunInTx(ctx, s.db, func(tx *sql.Tx) error {
		existing, err := getBookmarkTx(ctx, tx, id)
		if err != nil {
			return err
		}

		if req.Title != nil {
			existing.Title = *req.Title
		}
		if req.URL != nil {
			existing.URL = *req.URL
		}
		if req.Tags != nil {
			existing.Tags = append([]string(nil), req.Tags...)
		}

		tagsJSON, err := marshalTags(existing.Tags)
		if err != nil {
			return fmt.Errorf("marshal tags: %w", err)
		}

		result, err := tx.ExecContext(ctx, `
			UPDATE bookmarks
			SET title = ?, url = ?, tags = ?, updated_at = datetime('now')
			WHERE id = ?`, existing.Title, existing.URL, tagsJSON, id)
		if err != nil {
			return fmt.Errorf("update bookmark id=%d: %w", id, err)
		}
		affected, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("rows affected update id=%d: %w", id, err)
		}
		if affected == 0 {
			return ErrNotFound
		}

		detail := auditDetail(req)
		if _, err := tx.ExecContext(ctx, `
			INSERT INTO bookmark_audit (bookmark_id, action, detail)
			VALUES (?, 'updated', ?)`, id, detail); err != nil {
			return fmt.Errorf("insert audit bookmark_id=%d: %w", id, err)
		}

		updated, err = getBookmarkTx(ctx, tx, id)
		return err
	})
	if err != nil {
		return model.Bookmark{}, err
	}
	return updated, nil
}

func (s *DBStore) Delete(ctx context.Context, id int) error {
	result, err := s.deleteStmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("delete bookmark id=%d: %w", id, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected delete id=%d: %w", id, err)
	}
	if affected == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *DBStore) AuditCount(ctx context.Context, bookmarkID int) (int, error) {
	var count int
	err := s.db.QueryRowContext(ctx, `
		SELECT COUNT(*) FROM bookmark_audit WHERE bookmark_id = ?`, bookmarkID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count audit bookmark_id=%d: %w", bookmarkID, err)
	}
	return count, nil
}

func getBookmarkTx(ctx context.Context, tx *sql.Tx, id int) (model.Bookmark, error) {
	row := tx.QueryRowContext(ctx, `
		SELECT id, title, url, tags, created_at, updated_at
		FROM bookmarks WHERE id = ?`, id)
	b, err := scanBookmark(row)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Bookmark{}, ErrNotFound
	}
	if err != nil {
		return model.Bookmark{}, fmt.Errorf("get bookmark id=%d: %w", id, err)
	}
	return b, nil
}

func auditDetail(req model.UpdateBookmarkRequest) string {
	var parts []string
	if req.Title != nil {
		parts = append(parts, "title")
	}
	if req.URL != nil {
		parts = append(parts, "url")
	}
	if req.Tags != nil {
		parts = append(parts, "tags")
	}
	if len(parts) == 0 {
		return "no fields changed"
	}
	return "changed: " + strings.Join(parts, ", ")
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanBookmark(scanner rowScanner) (model.Bookmark, error) {
	var b model.Bookmark
	var tagsJSON, createdAt, updatedAt string
	if err := scanner.Scan(&b.ID, &b.Title, &b.URL, &tagsJSON, &createdAt, &updatedAt); err != nil {
		return model.Bookmark{}, err
	}
	if err := json.Unmarshal([]byte(tagsJSON), &b.Tags); err != nil {
		b.Tags = nil
	}
	b.CreatedAt = parseSQLiteTime(createdAt)
	b.UpdatedAt = parseSQLiteTime(updatedAt)
	return b, nil
}

func marshalTags(tags []string) (string, error) {
	if tags == nil {
		tags = []string{}
	}
	raw, err := json.Marshal(tags)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

func parseSQLiteTime(raw string) time.Time {
	if t, err := time.Parse(time.RFC3339, raw); err == nil {
		return t
	}
	if t, err := time.Parse("2006-01-02 15:04:05", raw); err == nil {
		return t
	}
	return time.Time{}
}
