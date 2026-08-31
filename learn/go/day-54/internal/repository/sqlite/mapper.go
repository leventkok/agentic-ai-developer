package sqlite

import (
	"database/sql"
	"encoding/json"
	"time"

	"learn/go/day-54/internal/auth"
	"learn/go/day-54/internal/model"
)

type rowScanner interface {
	Scan(dest ...any) error
}

func ScanBookmark(scanner rowScanner) (model.Bookmark, error) {
	var b model.Bookmark
	var tagsJSON, createdAt, updatedAt string
	var userID sql.NullInt64
	if err := scanner.Scan(&b.ID, &userID, &b.Title, &b.URL, &tagsJSON, &createdAt, &updatedAt); err != nil {
		return model.Bookmark{}, err
	}
	if userID.Valid {
		id := int(userID.Int64)
		b.UserID = &id
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

func canModifyBookmark(actor model.User, b model.Bookmark) bool {
	if actor.Role == auth.RoleAdmin {
		return true
	}
	if b.UserID == nil {
		return false
	}
	return *b.UserID == actor.ID
}
