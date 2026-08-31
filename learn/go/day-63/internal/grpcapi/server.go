package grpcapi

import (
	"context"

	bookmarksv1 "learn/go/day-63/internal/gen/bookmarksv1"
	"learn/go/day-63/internal/service"
)

// Server implements the generated BookmarkService gRPC interface.
type Server struct {
	bookmarksv1.UnimplementedBookmarkServiceServer
	Bookmarks *service.BookmarkService
	Auth      *service.AuthService
}

// NewServer wires bookmark and auth services into a gRPC handler.
func NewServer(bookmarks *service.BookmarkService, auth *service.AuthService) *Server {
	return &Server{Bookmarks: bookmarks, Auth: auth}
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
	token, err := BearerToken(ctx)
	if err != nil {
		return nil, err
	}
	user, err := s.Auth.UserFromToken(ctx, token)
	if err != nil {
		return nil, ToStatus(err)
	}
	bookmark, err := s.Bookmarks.Create(ctx, user, CreateInputFromProto(req))
	if err != nil {
		return nil, ToStatus(err)
	}
	return BookmarkToProto(bookmark), nil
}

func (s *Server) DeleteBookmark(ctx context.Context, req *bookmarksv1.DeleteBookmarkRequest) (*bookmarksv1.DeleteBookmarkResponse, error) {
	token, err := BearerToken(ctx)
	if err != nil {
		return nil, err
	}
	user, err := s.Auth.UserFromToken(ctx, token)
	if err != nil {
		return nil, ToStatus(err)
	}
	if err := s.Bookmarks.Delete(ctx, user, int(req.GetId())); err != nil {
		return nil, ToStatus(err)
	}
	return &bookmarksv1.DeleteBookmarkResponse{}, nil
}
