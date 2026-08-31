package resilient

import (
	"context"
	"time"

	"github.com/sony/gobreaker"

	"learn/go/day-92/internal/domain"
	"learn/go/day-92/internal/repository"
	"learn/go/day-92/internal/resilience"
)

// Bookmarks wraps a repository with retry + circuit breaker for read paths.
type Bookmarks struct {
	inner   repository.Bookmarks
	breaker *gobreaker.CircuitBreaker
}

func NewBookmarks(inner repository.Bookmarks) *Bookmarks {
	return &Bookmarks{
		inner: inner,
		breaker: gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "bookmarks-db",
			Timeout: 5 * time.Second,
		}),
	}
}

func (b *Bookmarks) List(ctx context.Context) ([]domain.Bookmark, error) {
	result, err := b.breaker.Execute(func() (any, error) {
		var list []domain.Bookmark
		err := resilience.Retry(ctx, 3, 10*time.Millisecond, func() error {
			var innerErr error
			list, innerErr = b.inner.List(ctx)
			return innerErr
		})
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]domain.Bookmark), nil
}

func (b *Bookmarks) Get(ctx context.Context, id int) (domain.Bookmark, error) {
	result, err := b.breaker.Execute(func() (any, error) {
		var bookmark domain.Bookmark
		err := resilience.Retry(ctx, 3, 10*time.Millisecond, func() error {
			var innerErr error
			bookmark, innerErr = b.inner.Get(ctx, id)
			return innerErr
		})
		if err != nil {
			return nil, err
		}
		return bookmark, nil
	})
	if err != nil {
		return domain.Bookmark{}, err
	}
	return result.(domain.Bookmark), nil
}

func (b *Bookmarks) Create(ctx context.Context, actor domain.User, in domain.CreateBookmarkInput) (domain.Bookmark, error) {
	return b.inner.Create(ctx, actor, in)
}

func (b *Bookmarks) BulkCreate(ctx context.Context, actor domain.User, inputs []domain.CreateBookmarkInput) ([]domain.Bookmark, error) {
	return b.inner.BulkCreate(ctx, actor, inputs)
}

func (b *Bookmarks) Update(ctx context.Context, actor domain.User, id int, in domain.UpdateBookmarkInput) (domain.Bookmark, error) {
	return b.inner.Update(ctx, actor, id, in)
}

func (b *Bookmarks) Delete(ctx context.Context, actor domain.User, id int) error {
	return b.inner.Delete(ctx, actor, id)
}

func (b *Bookmarks) Close() error {
	if c, ok := b.inner.(interface{ Close() error }); ok {
		return c.Close()
	}
	return nil
}
