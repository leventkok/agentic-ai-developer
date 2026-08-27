package store

import (
	"sync"
	"time"
	"context"
	"learn/go/day-39/internal/model"
)

type MemoryStore struct {
	mu        sync.RWMutex
	bookmarks map[int]model.Bookmark
	nextID    int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
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
	return b
}

func (s *MemoryStore) List(ctx context.Context) ([]model.Bookmark, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]model.Bookmark, 0, len(s.bookmarks))
	for _, b := range s.bookmarks {

		select{
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
		}

		time.Sleep(10 * time.Millisecond)


		list = append(list, copyBookmark(b))
	}
	return list, nil
}

func (s *MemoryStore) Get(ctx context.Context, id int) (model.Bookmark, error) {
	
	select{
		case <-ctx.Done():
			return model.Bookmark{}, ctx.Err()
		default:
	}
	
	
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.bookmarks[id]
	if !ok {
		return model.Bookmark{}, ErrNotFound
	}
	return copyBookmark(b), nil
}

func (s *MemoryStore) Create(ctx context.Context, req model.CreateBookmarkRequest) (model.Bookmark, error) {
	select{
		case <-ctx.Done():
			return model.Bookmark{}, ctx.Err()
		default:
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	b := model.Bookmark{
		ID:    s.nextID,
		Title: req.Title,
		URL:   req.URL,
		Tags:  copyTags(req.Tags),
	}
	s.bookmarks[s.nextID] = b
	s.nextID++
	return copyBookmark(b), nil
}

func (s *MemoryStore) Update(ctx context.Context, id int, req model.UpdateBookmarkRequest) (model.Bookmark, error) {


	select{
		case <-ctx.Done():
			return model.Bookmark{}, ctx.Err()
		default:
	}


	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.bookmarks[id]
	if !ok {
		return model.Bookmark{}, ErrNotFound
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
	s.bookmarks[id] = b
	return copyBookmark(b), nil
}

func (s *MemoryStore) Delete(ctx context.Context, id int) error {

	select{
		case <-ctx.Done():
			return ctx.Err()
		default:
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.bookmarks[id]; !ok {
		return ErrNotFound
	}
	delete(s.bookmarks, id)
	return nil
}
