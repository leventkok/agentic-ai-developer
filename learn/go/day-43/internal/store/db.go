package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"learn/go/day-43/internal/model"
)

type DBStore struct {
	db         *sql.DB
	listStmt   *sql.Stmt
	getStmt    *sql.Stmt
	insertStmt *sql.Stmt
	updateStmt *sql.Stmt
	deleteStmt *sql.Stmt
}

func NewDBStore(db *sql.DB) (*DBStore, error) {
	s := &DBStore{db: db}

	var err error
	if s.listStmt, err = prepare(db, "list bookmarks", `
		SELECT id, title, url, tags, created_at, updated_at
		FROM bookmarks ORDER BY id`); err != nil {
		s.Close()
		return nil, err
	}
	if s.getStmt, err = prepare(db, "get bookmark", `
		SELECT id, title, url, tags, created_at, updated_at
		FROM bookmarks WHERE id = ?`); err != nil {
		s.Close()
		return nil, err
	}
	if s.insertStmt, err = prepare(db, "insert bookmark", `
		INSERT INTO bookmarks (title, url, tags, created_at, updated_at)
		VALUES (?, ?, ?, datetime('now'), datetime('now'))
		RETURNING id, title, url, tags, created_at, updated_at`); err != nil {
		s.Close()
		return nil, err
	}
	if s.updateStmt, err = prepare(db, "update bookmark", `
		UPDATE bookmarks
		SET title = ?, url = ?, tags = ?, updated_at = datetime('now')
		WHERE id = ?`); err != nil {
		s.Close()
		return nil, err
	}
	if s.deleteStmt, err = prepare(db, "delete bookmark", `
		DELETE FROM bookmarks WHERE id = ?`); err != nil {
		s.Close()
		return nil, err
	}

	return s, nil
}

func prepare(db *sql.DB, name, query string) (*sql.Stmt, error) {
	stmt, err := db.Prepare(query)
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

func (s *DBStore) Update(ctx context.Context, id int, req model.UpdateBookmarkRequest) (model.Bookmark, error) {
	existing, err := s.Get(ctx, id)
	if err != nil {
		return model.Bookmark{}, err
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
		return model.Bookmark{}, fmt.Errorf("marshal tags: %w", err)
	}

	result, err := s.updateStmt.ExecContext(ctx, existing.Title, existing.URL, tagsJSON, id)
	if err != nil {
		return model.Bookmark{}, fmt.Errorf("update bookmark id=%d: %w", id, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return model.Bookmark{}, fmt.Errorf("rows affected update id=%d: %w", id, err)
	}
	if affected == 0 {
		return model.Bookmark{}, ErrNotFound
	}

	return s.Get(ctx, id)
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
