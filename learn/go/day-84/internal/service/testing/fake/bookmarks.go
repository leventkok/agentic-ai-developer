package fake

import (
	"context"
	"sync"
	"time"

	"learn/go/day-84/internal/domain"
)

type Bookmarks struct {
	mu        sync.Mutex
	next      int
	byID      map[int]domain.Bookmark
	listDelay time.Duration
	listHook  func(context.Context) ([]domain.Bookmark, error)
}

func NewBookmarks(seed ...domain.Bookmark) *Bookmarks {
	f := &Bookmarks{byID: make(map[int]domain.Bookmark), next: 1}
	for _, b := range seed {
		copy := b
		if copy.ID >= f.next {
			f.next = copy.ID + 1
		}
		f.byID[copy.ID] = copy
	}
	return f
}

func (f *Bookmarks) SetListDelay(d time.Duration) {
	f.listDelay = d
}

func (f *Bookmarks) SetListHook(hook func(context.Context) ([]domain.Bookmark, error)) {
	f.listHook = hook
}

func copyBookmark(b domain.Bookmark) domain.Bookmark {
	if b.Tags != nil {
		b.Tags = append([]string(nil), b.Tags...)
	}
	if b.UserID != nil {
		id := *b.UserID
		b.UserID = &id
	}
	return b
}

func (f *Bookmarks) List(ctx context.Context) ([]domain.Bookmark, error) {
	if f.listHook != nil {
		return f.listHook(ctx)
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.listDelay > 0 {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(f.listDelay):
		}
	}
	list := make([]domain.Bookmark, 0, len(f.byID))
	for _, b := range f.byID {
		list = append(list, copyBookmark(b))
	}
	return list, nil
}

func (f *Bookmarks) Get(ctx context.Context, id int) (domain.Bookmark, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	b, ok := f.byID[id]
	if !ok {
		return domain.Bookmark{}, domain.ErrNotFound
	}
	return copyBookmark(b), nil
}

func (f *Bookmarks) Create(ctx context.Context, actor domain.User, in domain.CreateBookmarkInput) (domain.Bookmark, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	ownerID := actor.ID
	now := time.Now().UTC()
	b := domain.Bookmark{
		ID:        f.next,
		UserID:    &ownerID,
		Title:     in.Title,
		URL:       in.URL,
		Tags:      append([]string(nil), in.Tags...),
		CreatedAt: now,
		UpdatedAt: now,
	}
	f.byID[f.next] = b
	f.next++
	return copyBookmark(b), nil
}

func (f *Bookmarks) BulkCreate(ctx context.Context, actor domain.User, inputs []domain.CreateBookmarkInput) ([]domain.Bookmark, error) {
	out := make([]domain.Bookmark, 0, len(inputs))
	for _, in := range inputs {
		b, err := f.Create(ctx, actor, in)
		if err != nil {
			return nil, err
		}
		out = append(out, b)
	}
	return out, nil
}

func (f *Bookmarks) Update(ctx context.Context, actor domain.User, id int, in domain.UpdateBookmarkInput) (domain.Bookmark, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	b, ok := f.byID[id]
	if !ok {
		return domain.Bookmark{}, domain.ErrNotFound
	}
	if in.Title != nil {
		b.Title = *in.Title
	}
	if in.URL != nil {
		b.URL = *in.URL
	}
	if in.Tags != nil {
		b.Tags = append([]string(nil), in.Tags...)
	}
	b.UpdatedAt = time.Now().UTC()
	f.byID[id] = b
	return copyBookmark(b), nil
}

func (f *Bookmarks) Delete(ctx context.Context, actor domain.User, id int) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, ok := f.byID[id]; !ok {
		return domain.ErrNotFound
	}
	delete(f.byID, id)
	return nil
}
