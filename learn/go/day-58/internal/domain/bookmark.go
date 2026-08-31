package domain

import "time"

type Bookmark struct {
	ID        int       `json:"id"`
	UserID    *int      `json:"user_id,omitempty"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Tags      []string  `json:"tags,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
