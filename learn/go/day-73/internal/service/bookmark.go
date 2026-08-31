package service

import (
	"context"
	"time"

	"learn/go/day-73/internal/domain"
	"learn/go/day-73/internal/repository"
)

type BookmarkService struct {
	Repo        repository.Bookmarks
	ListTimeout time.Duration
}

func NewBookmarkService(repo repository.Bookmarks, listTimeout time.Duration) *BookmarkService {
	return &BookmarkService{Repo: repo, ListTimeout: listTimeout}
}

func (s *BookmarkService) List(ctx context.Context) ([]domain.Bookmark, error) {
	timeout := s.ListTimeout
	if timeout <= 0 {
		timeout = 100 * time.Millisecond
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return s.Repo.List(ctx)
}

func (s *BookmarkService) Get(ctx context.Context, id int) (domain.Bookmark, error) {
	return s.Repo.Get(ctx, id)
}

func (s *BookmarkService) Create(ctx context.Context, actor domain.User, in domain.CreateBookmarkInput) (domain.Bookmark, error) {
	validated, err := domain.ValidateCreateInput(in)
	if err != nil {
		return domain.Bookmark{}, err
	}
	return s.Repo.Create(ctx, actor, validated)
}

func (s *BookmarkService) BulkCreate(ctx context.Context, actor domain.User, inputs []domain.CreateBookmarkInput) ([]domain.Bookmark, error) {
	if !domain.CanBulkCreate(actor) {
		return nil, domain.ErrForbidden
	}
	validated, err := domain.ValidateBulkCreateInputs(inputs)
	if err != nil {
		return nil, err
	}
	return s.Repo.BulkCreate(ctx, actor, validated)
}

func (s *BookmarkService) Update(ctx context.Context, actor domain.User, id int, in domain.UpdateBookmarkInput) (domain.Bookmark, error) {
	existing, err := s.Repo.Get(ctx, id)
	if err != nil {
		return domain.Bookmark{}, err
	}
	if !domain.CanModifyBookmark(actor, existing) {
		return domain.Bookmark{}, domain.ErrForbidden
	}
	validated, err := domain.ValidateUpdateInput(in)
	if err != nil {
		return domain.Bookmark{}, err
	}
	return s.Repo.Update(ctx, actor, id, validated)
}

func (s *BookmarkService) Delete(ctx context.Context, actor domain.User, id int) error {
	existing, err := s.Repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !domain.CanModifyBookmark(actor, existing) {
		return domain.ErrForbidden
	}
	return s.Repo.Delete(ctx, actor, id)
}
