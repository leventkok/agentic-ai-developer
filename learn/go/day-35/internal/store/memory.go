package store

import (
	"sync"

	"learn/go/day-35/internal/model"
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

func (s *MemoryStore) List() []model.Bookmark {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]model.Bookmark, 0, len(s.bookmarks))
	for _, b := range s.bookmarks {
		list = append(list, copyBookmark(b))
	}
	return list
}

func (s *MemoryStore) Get(id int) (model.Bookmark, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.bookmarks[id]
	if !ok {
		return model.Bookmark{}, false
	}
	return copyBookmark(b), true
}

func (s *MemoryStore) Create(req model.CreateBookmarkRequest) model.Bookmark {
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
	return copyBookmark(b)
}

func (s *MemoryStore) Update(id int, req model.UpdateBookmarkRequest) (model.Bookmark, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.bookmarks[id]
	if !ok {
		return model.Bookmark{}, false
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
	return copyBookmark(b), true
}

func (s *MemoryStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.bookmarks[id]; !ok {
		return false
	}
	delete(s.bookmarks, id)
	return true
}
