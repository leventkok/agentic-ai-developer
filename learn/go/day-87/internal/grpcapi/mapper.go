package grpcapi

import (
	"time"

	bookmarksv1 "learn/go/day-87/internal/gen/bookmarksv1"
	"learn/go/day-87/internal/domain"
)

// BookmarkToProto converts a domain bookmark to its protobuf representation.
func BookmarkToProto(b domain.Bookmark) *bookmarksv1.Bookmark {
	pb := &bookmarksv1.Bookmark{
		Id:        int32(b.ID),
		Title:     b.Title,
		Url:       b.URL,
		Tags:      append([]string(nil), b.Tags...),
		CreatedAt: b.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: b.UpdatedAt.UTC().Format(time.RFC3339),
	}
	if b.UserID != nil {
		pb.UserId = int32(*b.UserID)
	}
	return pb
}

// BookmarksToProto converts a slice of domain bookmarks to protobuf messages.
func BookmarksToProto(list []domain.Bookmark) []*bookmarksv1.Bookmark {
	items := make([]*bookmarksv1.Bookmark, 0, len(list))
	for _, b := range list {
		items = append(items, BookmarkToProto(b))
	}
	return items
}

// CreateInputFromProto maps a create request message to domain input.
func CreateInputFromProto(req *bookmarksv1.CreateBookmarkRequest) domain.CreateBookmarkInput {
	return domain.CreateBookmarkInput{
		Title: req.GetTitle(),
		URL:   req.GetUrl(),
		Tags:  append([]string(nil), req.GetTags()...),
	}
}
