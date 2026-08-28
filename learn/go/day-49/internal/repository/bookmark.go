package repository

import (
	"context"
	"errors"

	"learn/go/day-49/internal/model"
)

var ErrNotFound = errors.New("bookmark not found")

// Bookmarks exposes domain operations — no SQL in handlers.
type Bookmarks interface {
	List(ctx context.Context) ([]model.Bookmark, error)
	Get(ctx context.Context, id int) (model.Bookmark, error)
	Create(ctx context.Context, req model.CreateBookmarkRequest) (model.Bookmark, error)
	BulkCreate(ctx context.Context, reqs []model.CreateBookmarkRequest) ([]model.Bookmark, error)
	Update(ctx context.Context, id int, req model.UpdateBookmarkRequest) (model.Bookmark, error)
	Delete(ctx context.Context, id int) error
}
