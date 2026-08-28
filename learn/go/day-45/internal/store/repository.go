package store

import (
	"context"
	"errors"

	"learn/go/day-45/internal/model"
)

var ErrNotFound = errors.New("bookmark not found")

type BookmarkRepository interface {
	List(ctx context.Context) ([]model.Bookmark, error)
	Get(ctx context.Context, id int) (model.Bookmark, error)
	Create(ctx context.Context, req model.CreateBookmarkRequest) (model.Bookmark, error)
	BulkCreate(ctx context.Context, reqs []model.CreateBookmarkRequest) ([]model.Bookmark, error)
	Update(ctx context.Context, id int, req model.UpdateBookmarkRequest) (model.Bookmark, error)
	Delete(ctx context.Context, id int) error
}
