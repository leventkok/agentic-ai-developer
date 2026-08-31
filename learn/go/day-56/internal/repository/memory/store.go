package memory

import (
	"context"
	"sync"
	"time"

	"learn/go/day-56/internal/domain"
)

type Store struct {
	mu        sync.RWMutex
	bookmarks map[int]domain.Bookmark
	nextID    int
}

func New() *Store {
	return &Store{
		bookmarks: make(map[int]domain.Bookmark),
		nextID:    1,
	}
}

func copyTags(tags []string) []string {
	if tags == nil {
		return nil
	}
	return append([]string(nil), tags...)
}

func copyBookmark(b domain.Bookmark) domain.Bookmark {
	b.Tags = copyTags(b.Tags)
	if b.UserID != nil {
		id := *b.UserID
		b.UserID = &id
	}
	return b
}

func canModify(actor domain.User, b domain.Bookmark) bool {
	if actor.Role == domain.RoleAdmin {
		return true
	}
	if b.UserID == nil {
		return false
	}
	return *b.UserID == actor.ID
}

func (s *Store) List(ctx context.Context) ([]domain.Bookmark, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]domain.Bookmark, 0, len(s.bookmarks))
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

func (s *Store) Get(ctx context.Context, id int) (domain.Bookmark, error) {
	select {
	case <-ctx.Done():
		return domain.Bookmark{}, ctx.Err()
	default:
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.bookmarks[id]
	if !ok {
		return domain.Bookmark{}, domain.ErrNotFound
	}
	return copyBookmark(b), nil
}

func (s *Store) Create(ctx context.Context, actor domain.User, in domain.CreateBookmarkInput) (domain.Bookmark, error) {
	select {
	case <-ctx.Done():
		return domain.Bookmark{}, ctx.Err()
	default:
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UTC()
	ownerID := actor.ID
	b := domain.Bookmark{
		ID:        s.nextID,
		UserID:    &ownerID,
		Title:     in.Title,
		URL:       in.URL,
		Tags:      copyTags(in.Tags),
		CreatedAt: now,
		UpdatedAt: now,
	}
	s.bookmarks[s.nextID] = b
	s.nextID++
	return copyBookmark(b), nil
}

func (s *Store) BulkCreate(ctx context.Context, actor domain.User, inputs []domain.CreateBookmarkInput) ([]domain.Bookmark, error) {
	created := make([]domain.Bookmark, 0, len(inputs))
	for _, in := range inputs {
		b, err := s.Create(ctx, actor, in)
		if err != nil {
			return nil, err
		}
		created = append(created, b)
	}
	return created, nil
}

func (s *Store) Update(ctx context.Context, actor domain.User, id int, in domain.UpdateBookmarkInput) (domain.Bookmark, error) {
	select {
	case <-ctx.Done():
		return domain.Bookmark{}, ctx.Err()
	default:
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.bookmarks[id]
	if !ok {
		return domain.Bookmark{}, domain.ErrNotFound
	}
	if !canModify(actor, b) {
		return domain.Bookmark{}, domain.ErrForbidden
	}
	if in.Title != nil {
		b.Title = *in.Title
	}
	if in.URL != nil {
		b.URL = *in.URL
	}
	if in.Tags != nil {
		b.Tags = copyTags(in.Tags)
	}
	b.UpdatedAt = time.Now().UTC()
	s.bookmarks[id] = b
	return copyBookmark(b), nil
}

func (s *Store) Delete(ctx context.Context, actor domain.User, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.bookmarks[id]
	if !ok {
		return domain.ErrNotFound
	}
	if !canModify(actor, b) {
		return domain.ErrForbidden
	}
	delete(s.bookmarks, id)
	return nil
}
