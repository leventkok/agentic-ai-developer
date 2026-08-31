package bookmarksv1_test

import (
	"testing"

	"google.golang.org/protobuf/proto"

	bookmarksv1 "learn/go/day-79/internal/gen/bookmarksv1"
)

func TestBookmark_RoundTrip(t *testing.T) {
	userID := int32(42)
	original := &bookmarksv1.Bookmark{
		Id:        1,
		UserId:    userID,
		Title:     "Go docs",
		Url:       "https://go.dev",
		Tags:      []string{"go"},
		CreatedAt: "2026-01-01T00:00:00Z",
		UpdatedAt: "2026-01-01T00:00:00Z",
	}

	data, err := proto.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}

	var decoded bookmarksv1.Bookmark
	if err := proto.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}

	if decoded.GetId() != original.GetId() {
		t.Fatalf("id = %d, want %d", decoded.GetId(), original.GetId())
	}
	if decoded.GetTitle() != original.GetTitle() {
		t.Fatalf("title = %q, want %q", decoded.GetTitle(), original.GetTitle())
	}
}

func TestCreateBookmarkRequest_RoundTrip(t *testing.T) {
	original := &bookmarksv1.CreateBookmarkRequest{
		Title: "Go Blog",
		Url:   "https://go.dev/blog",
		Tags:  []string{"go"},
	}

	data, err := proto.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}

	var decoded bookmarksv1.CreateBookmarkRequest
	if err := proto.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.GetTitle() != original.GetTitle() {
		t.Fatalf("title = %q, want %q", decoded.GetTitle(), original.GetTitle())
	}
}

func TestBookmarkList_Empty(t *testing.T) {
	original := &bookmarksv1.BookmarkList{}

	data, err := proto.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}

	var decoded bookmarksv1.BookmarkList
	if err := proto.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if len(decoded.GetItems()) != 0 {
		t.Fatalf("items len = %d, want 0", len(decoded.GetItems()))
	}
}
