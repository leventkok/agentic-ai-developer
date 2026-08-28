package sqlite

import (
	"encoding/json"
	"time"

	"learn/go/day-47/internal/model"
)

type rowScanner interface {
	Scan(dest ...any) error
}

func ScanBookmark(scanner rowScanner) (model.Bookmark, error) {
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

func MarshalTags(tags []string) (string, error) {
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
