package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"learn/go/day-42/internal/model"
)

func ListBookmarks(ctx context.Context, db *sql.DB) ([]model.Bookmark, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT id, title, url, tags, created_at, updated_at
		FROM bookmarks
		ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Bookmark
	for rows.Next() {
		var b model.Bookmark
		var tagsJSON, createdAt, updatedAt string
		if err := rows.Scan(&b.ID, &b.Title, &b.URL, &tagsJSON, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(tagsJSON), &b.Tags); err != nil {
			b.Tags = nil
		}
		b.CreatedAt = parseSQLiteTime(createdAt)
		b.UpdatedAt = parseSQLiteTime(updatedAt)
		list = append(list, b)
	}
	return list, rows.Err()
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
