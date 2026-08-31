package bookmarksv1_test

import (
	"testing"

	bookmarksv1 "learn/go/day-75/internal/gen/bookmarksv1"
)

func TestBookmarkServiceServerInterfaceExists(t *testing.T) {
	var _ bookmarksv1.BookmarkServiceServer = (*bookmarksv1.UnimplementedBookmarkServiceServer)(nil)
}

func TestBookmarkServiceClientInterfaceExists(t *testing.T) {
	// Client is an interface generated in bookmarks_grpc.pb.go
	var _ bookmarksv1.BookmarkServiceClient = (bookmarksv1.BookmarkServiceClient)(nil)
}
