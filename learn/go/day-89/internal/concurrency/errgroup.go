package concurrency

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// RunLimited runs fn for each item with at most limit concurrent goroutines.
func RunLimited[T any](ctx context.Context, limit int, items []T, fn func(context.Context, T) error) error {
	if limit <= 0 {
		limit = 4
	}
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(limit)
	for _, item := range items {
		item := item
		g.Go(func() error {
			return fn(ctx, item)
		})
	}
	return g.Wait()
}
