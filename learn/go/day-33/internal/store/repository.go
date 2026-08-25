package store

import "learn/go/day-33/internal/model"

// BookmarkRepository abstracts storage — swap in a DB later without changing handlers.
type BookmarkRepository interface {
	List() []model.Bookmark
	Get(id int) (model.Bookmark, bool)
	Create(req model.CreateBookmarkRequest) model.Bookmark
	Update(id int, req model.UpdateBookmarkRequest) (model.Bookmark, bool)
	Delete(id int) bool
}
