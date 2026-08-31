package messaging

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type EventType string

const (
	EventBookmarkCreated EventType = "bookmark.created"
	EventBookmarkUpdated EventType = "bookmark.updated"
)

type Event struct {
	ID         string    `json:"id"`
	Type       EventType `json:"type"`
	BookmarkID int       `json:"bookmark_id"`
	UserID     int       `json:"user_id"`
	OccurredAt time.Time `json:"occurred_at"`
}

func NewEvent(eventType EventType, bookmarkID, userID int) Event {
	return Event{
		ID:         uuid.NewString(),
		Type:       eventType,
		BookmarkID: bookmarkID,
		UserID:     userID,
		OccurredAt: time.Now().UTC(),
	}
}

func (e Event) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func UnmarshalEvent(data []byte) (Event, error) {
	var e Event
	if err := json.Unmarshal(data, &e); err != nil {
		return Event{}, fmt.Errorf("messaging: decode event: %w", err)
	}
	return e, nil
}

func (e Event) DeduplicationKey() string {
	if e.ID != "" {
		return e.ID
	}
	return fmt.Sprintf("%s:%d:%d", e.Type, e.BookmarkID, e.UserID)
}
