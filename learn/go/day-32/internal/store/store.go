package store

import (
	"sync"

	"learn/go/day-32/internal/model"
)


type BookmarkStore struct {

	mu sync.RWMutex
	bookmarks map[int]model.Bookmark
	nextID int

}


func NewBookmarkStore() *BookmarkStore {
	return &BookmarkStore{
		bookmarks: make(map[int]model.Bookmark),
		nextID: 1,
	}
}

func (s *BookmarkStore) List() []model.Bookmark {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]model.Bookmark, 0, len(s.bookmarks))
	for _, b := range s.bookmarks {
		list = append(list, b)
	}
	return list
}

func (s *BookmarkStore) Get(id int) (model.Bookmark, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.bookmarks[id]
	return b, ok
}


func (s *BookmarkStore) Create(req model.CreateBookmarkRequest) model.Bookmark {
	s.mu.Lock()
	defer s.mu.Unlock()
	b := model.Bookmark{
		ID: s.nextID,
		Title: req.Title,
		URL: req.URL,
		Tags: req.Tags,
	}
	s.bookmarks[s.nextID] = b
	s.nextID++
	return b
}


func (s *BookmarkStore) Update(id int, req model.UpdateBookmarkRequest) (model.Bookmark, bool) {
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
		b.Tags = req.Tags
	}
	s.bookmarks[id] = b
	return b, true
}


func (s *BookmarkStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.bookmarks[id]; !ok {
		return false
	}
	delete(s.bookmarks, id)
	return true
}
