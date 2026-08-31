package service

import (
	"context"
	"time"

	"learn/go/day-96/internal/domain"
	"learn/go/day-96/internal/messaging"
	"learn/go/day-96/internal/messaging/outbox"
	"learn/go/day-96/internal/repository"
)

type BookmarkService struct {
	Repo        repository.Bookmarks
	ListTimeout time.Duration
	Outbox      *outbox.Store
	Events      messaging.Bus
}

func NewBookmarkService(repo repository.Bookmarks, listTimeout time.Duration, outboxStore *outbox.Store, events messaging.Bus) *BookmarkService {
	return &BookmarkService{Repo: repo, ListTimeout: listTimeout, Outbox: outboxStore, Events: events}
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
	created, err := s.Repo.Create(ctx, actor, validated)
	if err != nil {
		return domain.Bookmark{}, err
	}
	s.publish(ctx, messaging.NewEvent(messaging.EventBookmarkCreated, created.ID, actor.ID))
	return created, nil
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
	updated, err := s.Repo.Update(ctx, actor, id, validated)
	if err != nil {
		return domain.Bookmark{}, err
	}
	s.publish(ctx, messaging.NewEvent(messaging.EventBookmarkUpdated, updated.ID, actor.ID))
	return updated, nil
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

func (s *BookmarkService) publish(ctx context.Context, evt messaging.Event) {
	if s.Outbox != nil {
		_ = s.Outbox.Enqueue(ctx, evt)
		return
	}
	if s.Events == nil {
		return
	}
	_ = s.Events.Publish(ctx, evt)
}
