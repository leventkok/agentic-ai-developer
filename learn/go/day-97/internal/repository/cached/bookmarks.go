package cached

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"learn/go/day-97/internal/cache"
	"learn/go/day-97/internal/domain"
	"learn/go/day-97/internal/repository"
)

const listKey = "bookmarks:list"

// Bookmarks adds cache-aside for List/Get and invalidates on writes.
type Bookmarks struct {
	inner repository.Bookmarks
	cache cache.Store
	ttl   time.Duration
}

func New(inner repository.Bookmarks, c cache.Store, ttl time.Duration) *Bookmarks {
	if ttl <= 0 {
		ttl = time.Minute
	}
	return &Bookmarks{inner: inner, cache: c, ttl: ttl}
}

func (b *Bookmarks) List(ctx context.Context) ([]domain.Bookmark, error) {
	if raw, ok, err := b.cache.Get(ctx, listKey); err == nil && ok {
		var list []domain.Bookmark
		if json.Unmarshal(raw, &list) == nil {
			return list, nil
		}
	}
	list, err := b.inner.List(ctx)
	if err != nil {
		return nil, err
	}
	if raw, err := json.Marshal(list); err == nil {
		_ = b.cache.Set(ctx, listKey, raw, b.ttl)
	}
	return list, nil
}

func (b *Bookmarks) Get(ctx context.Context, id int) (domain.Bookmark, error) {
	key := idKey(id)
	if raw, ok, err := b.cache.Get(ctx, key); err == nil && ok {
		var bookmark domain.Bookmark
		if json.Unmarshal(raw, &bookmark) == nil {
			return bookmark, nil
		}
	}
	bookmark, err := b.inner.Get(ctx, id)
	if err != nil {
		return domain.Bookmark{}, err
	}
	if raw, err := json.Marshal(bookmark); err == nil {
		_ = b.cache.Set(ctx, key, raw, b.ttl)
	}
	return bookmark, nil
}

func (b *Bookmarks) Create(ctx context.Context, actor domain.User, in domain.CreateBookmarkInput) (domain.Bookmark, error) {
	bookmark, err := b.inner.Create(ctx, actor, in)
	if err != nil {
		return domain.Bookmark{}, err
	}
	_ = b.invalidate(ctx, bookmark.ID)
	return bookmark, nil
}

func (b *Bookmarks) BulkCreate(ctx context.Context, actor domain.User, inputs []domain.CreateBookmarkInput) ([]domain.Bookmark, error) {
	bookmarks, err := b.inner.BulkCreate(ctx, actor, inputs)
	if err != nil {
		return nil, err
	}
	_ = b.invalidate(ctx)
	return bookmarks, nil
}

func (b *Bookmarks) Update(ctx context.Context, actor domain.User, id int, in domain.UpdateBookmarkInput) (domain.Bookmark, error) {
	bookmark, err := b.inner.Update(ctx, actor, id, in)
	if err != nil {
		return domain.Bookmark{}, err
	}
	_ = b.invalidate(ctx, id)
	return bookmark, nil
}

func (b *Bookmarks) Delete(ctx context.Context, actor domain.User, id int) error {
	if err := b.inner.Delete(ctx, actor, id); err != nil {
		return err
	}
	return b.invalidate(ctx, id)
}

func (b *Bookmarks) Close() error {
	if c, ok := b.inner.(interface{ Close() error }); ok {
		return c.Close()
	}
	return nil
}

func (b *Bookmarks) invalidate(ctx context.Context, ids ...int) error {
	keys := []string{listKey}
	for _, id := range ids {
		if id > 0 {
			keys = append(keys, idKey(id))
		}
	}
	return b.cache.Delete(ctx, keys...)
}

func idKey(id int) string {
	return fmt.Sprintf("bookmarks:id:%d", id)
}
