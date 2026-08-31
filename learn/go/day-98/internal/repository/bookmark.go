package repository

import (
	"context"

	"learn/go/day-98/internal/domain"
)

type Bookmarks interface {
	List(ctx context.Context) ([]domain.Bookmark, error)
	Get(ctx context.Context, id int) (domain.Bookmark, error)
	Create(ctx context.Context, actor domain.User, in domain.CreateBookmarkInput) (domain.Bookmark, error)
	BulkCreate(ctx context.Context, actor domain.User, inputs []domain.CreateBookmarkInput) ([]domain.Bookmark, error)
	Update(ctx context.Context, actor domain.User, id int, in domain.UpdateBookmarkInput) (domain.Bookmark, error)
	Delete(ctx context.Context, actor domain.User, id int) error
}
