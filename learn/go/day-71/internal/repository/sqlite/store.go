package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"learn/go/day-71/internal/db"
	"learn/go/day-71/internal/domain"
)

type Store struct {
	db         *sql.DB
	listStmt   *sql.Stmt
	getStmt    *sql.Stmt
	insertStmt *sql.Stmt
	updateStmt *sql.Stmt
	deleteStmt *sql.Stmt
}

func New(database *sql.DB) (*Store, error) {
	s := &Store{db: database}
	var err error
	if s.listStmt, err = prepare(database, "list bookmarks", db.SQLListBookmarks); err != nil {
		s.Close()
		return nil, err
	}
	if s.getStmt, err = prepare(database, "get bookmark", db.SQLGetBookmarkByID); err != nil {
		s.Close()
		return nil, err
	}
	if s.insertStmt, err = prepare(database, "insert bookmark", db.SQLInsertBookmark); err != nil {
		s.Close()
		return nil, err
	}
	if s.updateStmt, err = prepare(database, "update bookmark", db.SQLUpdateBookmark); err != nil {
		s.Close()
		return nil, err
	}
	if s.deleteStmt, err = prepare(database, "delete bookmark", db.SQLDeleteBookmark); err != nil {
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

func (s *Store) Close() error {
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

func (s *Store) List(ctx context.Context) ([]domain.Bookmark, error) {
	rows, err := s.listStmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("list bookmarks: %w", err)
	}
	defer rows.Close()

	var list []domain.Bookmark
	for rows.Next() {
		b, err := ScanBookmark(rows)
		if err != nil {
			return nil, fmt.Errorf("scan bookmark row: %w", err)
		}
		list = append(list, b)
	}
	return list, rows.Err()
}

func (s *Store) Get(ctx context.Context, id int) (domain.Bookmark, error) {
	row := s.getStmt.QueryRowContext(ctx, id)
	b, err := ScanBookmark(row)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Bookmark{}, domain.ErrNotFound
	}
	if err != nil {
		return domain.Bookmark{}, fmt.Errorf("get bookmark id=%d: %w", id, err)
	}
	return b, nil
}

func (s *Store) Create(ctx context.Context, actor domain.User, req domain.CreateBookmarkInput) (domain.Bookmark, error) {
	tagsJSON, err := MarshalTags(req.Tags)
	if err != nil {
		return domain.Bookmark{}, fmt.Errorf("marshal tags: %w", err)
	}
	row := s.insertStmt.QueryRowContext(ctx, actor.ID, req.Title, req.URL, tagsJSON)
	b, err := ScanBookmark(row)
	if err != nil {
		return domain.Bookmark{}, fmt.Errorf("insert bookmark: %w", err)
	}
	return b, nil
}

func (s *Store) BulkCreate(ctx context.Context, actor domain.User, reqs []domain.CreateBookmarkInput) ([]domain.Bookmark, error) {
	if len(reqs) == 0 {
		return nil, nil
	}
	var created []domain.Bookmark
	err := db.RunInTx(ctx, s.db, func(tx *sql.Tx) error {
		stmt, err := tx.PrepareContext(ctx, db.SQLInsertBookmark)
		if err != nil {
			return fmt.Errorf("prepare bulk insert: %w", err)
		}
		defer stmt.Close()

		created = make([]domain.Bookmark, 0, len(reqs))
		for i, req := range reqs {
			tagsJSON, err := MarshalTags(req.Tags)
			if err != nil {
				return fmt.Errorf("marshal tags item %d: %w", i, err)
			}
			row := stmt.QueryRowContext(ctx, actor.ID, req.Title, req.URL, tagsJSON)
			b, err := ScanBookmark(row)
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

func (s *Store) Update(ctx context.Context, actor domain.User, id int, req domain.UpdateBookmarkInput) (domain.Bookmark, error) {
	existing, err := s.Get(ctx, id)
	if err != nil {
		return domain.Bookmark{}, err
	}

	var updated domain.Bookmark
	err = db.RunInTx(ctx, s.db, func(tx *sql.Tx) error {
		if req.Title != nil {
			existing.Title = *req.Title
		}
		if req.URL != nil {
			existing.URL = *req.URL
		}
		if req.Tags != nil {
			existing.Tags = append([]string(nil), req.Tags...)
		}

		tagsJSON, err := MarshalTags(existing.Tags)
		if err != nil {
			return fmt.Errorf("marshal tags: %w", err)
		}

		result, err := tx.ExecContext(ctx, db.SQLUpdateBookmark, existing.Title, existing.URL, tagsJSON, id)
		if err != nil {
			return fmt.Errorf("update bookmark id=%d: %w", id, err)
		}
		affected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if affected == 0 {
			return domain.ErrNotFound
		}

		if _, err := tx.ExecContext(ctx, db.SQLInsertAudit, id, auditDetail(req)); err != nil {
			return fmt.Errorf("insert audit bookmark_id=%d: %w", id, err)
		}

		updated, err = getBookmarkTx(ctx, tx, id)
		return err
	})
	if err != nil {
		return domain.Bookmark{}, err
	}
	return updated, nil
}

func (s *Store) Delete(ctx context.Context, actor domain.User, id int) error {
	if _, err := s.Get(ctx, id); err != nil {
		return err
	}

	result, err := s.deleteStmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("delete bookmark id=%d: %w", id, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (s *Store) AuditCount(ctx context.Context, bookmarkID int) (int, error) {
	var count int
	err := s.db.QueryRowContext(ctx, db.SQLCountAuditByBookmark, bookmarkID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count audit bookmark_id=%d: %w", bookmarkID, err)
	}
	return count, nil
}

func getBookmarkTx(ctx context.Context, tx *sql.Tx, id int) (domain.Bookmark, error) {
	row := tx.QueryRowContext(ctx, db.SQLGetBookmarkByID, id)
	b, err := ScanBookmark(row)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Bookmark{}, domain.ErrNotFound
	}
	if err != nil {
		return domain.Bookmark{}, fmt.Errorf("get bookmark id=%d: %w", id, err)
	}
	return b, nil
}

func auditDetail(req domain.UpdateBookmarkInput) string {
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
