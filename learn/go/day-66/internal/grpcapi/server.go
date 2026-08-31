package grpcapi

import (
	"context"

	"learn/go/day-66/internal/ctxkey"
	"learn/go/day-66/internal/domain"
	bookmarksv1 "learn/go/day-66/internal/gen/bookmarksv1"
	"learn/go/day-66/internal/service"
)

// Server implements the generated BookmarkService gRPC interface.
type Server struct {
	bookmarksv1.UnimplementedBookmarkServiceServer
	Bookmarks *service.BookmarkService
}

// NewServer wires bookmark services into a gRPC handler.
func NewServer(bookmarks *service.BookmarkService) *Server {
	return &Server{Bookmarks: bookmarks}
}

func (s *Server) ListBookmarks(ctx context.Context, _ *bookmarksv1.ListBookmarksRequest) (*bookmarksv1.BookmarkList, error) {
	list, err := s.Bookmarks.List(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}
	return &bookmarksv1.BookmarkList{Items: BookmarksToProto(list)}, nil
}

func (s *Server) GetBookmark(ctx context.Context, req *bookmarksv1.GetBookmarkRequest) (*bookmarksv1.Bookmark, error) {
	bookmark, err := s.Bookmarks.Get(ctx, int(req.GetId()))
	if err != nil {
		return nil, ToStatus(err)
	}
	return BookmarkToProto(bookmark), nil
}

func (s *Server) CreateBookmark(ctx context.Context, req *bookmarksv1.CreateBookmarkRequest) (*bookmarksv1.Bookmark, error) {
	user, ok := ctxkey.UserFromContext(ctx)
	if !ok {
		return nil, ToStatus(domain.ErrUnauthorized)
	}
	bookmark, err := s.Bookmarks.Create(ctx, user, CreateInputFromProto(req))
	if err != nil {
		return nil, ToStatus(err)
	}
	return BookmarkToProto(bookmark), nil
}

func (s *Server) DeleteBookmark(ctx context.Context, req *bookmarksv1.DeleteBookmarkRequest) (*bookmarksv1.DeleteBookmarkResponse, error) {
	user, ok := ctxkey.UserFromContext(ctx)
	if !ok {
		return nil, ToStatus(domain.ErrUnauthorized)
	}
	if err := s.Bookmarks.Delete(ctx, user, int(req.GetId())); err != nil {
		return nil, ToStatus(err)
	}
	return &bookmarksv1.DeleteBookmarkResponse{}, nil
}
