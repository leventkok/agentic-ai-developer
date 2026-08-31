package service

import (
	"context"

	"learn/go/day-59/internal/domain"
)

// TODO (Day 57): Replace pass-through Update with a real use case:
//  1. existing, err := s.Repo.Get(ctx, id)
//  2. if !domain.CanModifyBookmark(actor, existing) { return domain.ErrForbidden }
//  3. return s.Repo.Update(ctx, actor, id, in)
//
// Same pattern for Delete. Remove ownership checks from sqlite/store.go once service owns the flow.

func (s *BookmarkService) UpdateUseCase(ctx context.Context, actor domain.User, id int, in domain.UpdateBookmarkInput) (domain.Bookmark, error) {
	_ = actor
	_ = id
	_ = in
	panic("TODO: orchestrate Get → domain rule → Update")
}
