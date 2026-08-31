package cache

import (
	"context"
	"time"
)

// Store is a minimal cache-aside backend.
type Store interface {
	Get(ctx context.Context, key string) ([]byte, bool, error)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Delete(ctx context.Context, keys ...string) error
	Close() error
}
