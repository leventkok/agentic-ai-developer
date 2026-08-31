package service

import (
	"context"

	"learn/go/day-59/internal/domain"
)

// CreateValidated orchestrates domain validation before persistence.
// TODO (Day 59): Call domain.NewBookmark(in) then repo.Create; return domain validation errors.
func (s *BookmarkService) CreateValidated(ctx context.Context, actor domain.User, in domain.CreateBookmarkInput) (domain.Bookmark, error) {
	panic("TODO: implement — domain.NewBookmark → repo.Create")
}
