package memory

import (
	"context"
	"sync"
	"time"

	"learn/go/day-54/internal/auth"
	"learn/go/day-54/internal/model"
	"learn/go/day-54/internal/repository"
)

type Store struct {
	mu        sync.RWMutex
	bookmarks map[int]model.Bookmark
	nextID    int
}

func New() *Store {
	return &Store{
		bookmarks: make(map[int]model.Bookmark),
		nextID:    1,
	}
}

func copyTags(tags []string) []string {
	if tags == nil {
		return nil
	}
	return append([]string(nil), tags...)
}

func copyBookmark(b model.Bookmark) model.Bookmark {
	b.Tags = copyTags(b.Tags)
	if b.UserID != nil {
		id := *b.UserID
		b.UserID = &id
	}
	return b
}

func canModify(actor model.User, b model.Bookmark) bool {
	if actor.Role == auth.RoleAdmin {
		return true
	}
	if b.UserID == nil {
		return false
	}
	return *b.UserID == actor.ID
}

func (s *Store) List(ctx context.Context) ([]model.Bookmark, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]model.Bookmark, 0, len(s.bookmarks))
	for _, b := range s.bookmarks {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		time.Sleep(10 * time.Millisecond)
		list = append(list, copyBookmark(b))
	}
	return list, nil
}

func (s *Store) Get(ctx context.Context, id int) (model.Bookmark, error) {
	select {
	case <-ctx.Done():
		return model.Bookmark{}, ctx.Err()
	default:
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.bookmarks[id]
	if !ok {
		return model.Bookmark{}, repository.ErrNotFound
	}
	return copyBookmark(b), nil
}

func (s *Store) Create(ctx context.Context, actor model.User, req model.CreateBookmarkRequest) (model.Bookmark, error) {
	select {
	case <-ctx.Done():
		return model.Bookmark{}, ctx.Err()
	default:
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UTC()
	ownerID := actor.ID
	b := model.Bookmark{
		ID:        s.nextID,
		UserID:    &ownerID,
		Title:     req.Title,
		URL:       req.URL,
		Tags:      copyTags(req.Tags),
		CreatedAt: now,
		UpdatedAt: now,
	}
	s.bookmarks[s.nextID] = b
	s.nextID++
	return copyBookmark(b), nil
}

func (s *Store) BulkCreate(ctx context.Context, actor model.User, reqs []model.CreateBookmarkRequest) ([]model.Bookmark, error) {
	created := make([]model.Bookmark, 0, len(reqs))
	for _, req := range reqs {
		b, err := s.Create(ctx, actor, req)
		if err != nil {
			return nil, err
		}
		created = append(created, b)
	}
	return created, nil
}

func (s *Store) Update(ctx context.Context, actor model.User, id int, req model.UpdateBookmarkRequest) (model.Bookmark, error) {
	select {
	case <-ctx.Done():
		return model.Bookmark{}, ctx.Err()
	default:
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.bookmarks[id]
	if !ok {
		return model.Bookmark{}, repository.ErrNotFound
	}
	if !canModify(actor, b) {
		return model.Bookmark{}, repository.ErrForbidden
	}
	if req.Title != nil {
		b.Title = *req.Title
	}
	if req.URL != nil {
		b.URL = *req.URL
	}
	if req.Tags != nil {
		b.Tags = copyTags(req.Tags)
	}
	b.UpdatedAt = time.Now().UTC()
	s.bookmarks[id] = b
	return copyBookmark(b), nil
}

func (s *Store) Delete(ctx context.Context, actor model.User, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.bookmarks[id]
	if !ok {
		return repository.ErrNotFound
	}
	if !canModify(actor, b) {
		return repository.ErrForbidden
	}
	delete(s.bookmarks, id)
	return nil
}
