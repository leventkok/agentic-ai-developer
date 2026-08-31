package grpcapi_test

import (
	"testing"
	"time"

	"learn/go/day-76/internal/domain"
	"learn/go/day-76/internal/grpcapi"
)

func TestBookmarkToProto(t *testing.T) {
	userID := 7
	created := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	updated := time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC)

	pb := grpcapi.BookmarkToProto(domain.Bookmark{
		ID:        1,
		UserID:    &userID,
		Title:     "Go docs",
		URL:       "https://go.dev",
		Tags:      []string{"go"},
		CreatedAt: created,
		UpdatedAt: updated,
	})

	if pb.GetId() != 1 {
		t.Fatalf("id = %d, want 1", pb.GetId())
	}
	if pb.GetUserId() != 7 {
		t.Fatalf("user_id = %d, want 7", pb.GetUserId())
	}
}
