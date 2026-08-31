package grpcapi_test

import (
	"testing"
	"time"

	"learn/go/day-63/internal/domain"
	"learn/go/day-63/internal/grpcapi"
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
	if pb.GetTitle() != "Go docs" {
		t.Fatalf("title = %q, want Go docs", pb.GetTitle())
	}
	if pb.GetUrl() != "https://go.dev" {
		t.Fatalf("url = %q, want https://go.dev", pb.GetUrl())
	}
	if pb.GetCreatedAt() != "2026-01-01T00:00:00Z" {
		t.Fatalf("created_at = %q", pb.GetCreatedAt())
	}
}

func TestBookmarksToProto_Empty(t *testing.T) {
	items := grpcapi.BookmarksToProto(nil)
	if len(items) != 0 {
		t.Fatalf("len = %d, want 0", len(items))
	}
}
