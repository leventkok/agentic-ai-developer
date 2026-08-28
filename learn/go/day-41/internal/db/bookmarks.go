package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"learn/go/day-41/internal/model"
)

func ListBookmarks(ctx context.Context, db *sql.DB) ([]model.Bookmark, error) {
	rows, err := db.QueryContext(ctx, `SELECT id, title, url, tags FROM bookmarks ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close() 

	var list []model.Bookmark
	for rows.Next() {
		var b model.Bookmark
		var tagsJSON string
		if err := rows.Scan(&b.ID, &b.Title, &b.URL, &tagsJSON); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(tagsJSON), &b.Tags); err != nil {
			b.Tags = nil
		}
		list = append(list, b)
	}
	return list, rows.Err()
}