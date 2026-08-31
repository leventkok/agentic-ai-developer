package redis

import (
	"context"
	"time"

	goredis "github.com/redis/go-redis/v9"

	"learn/go/day-95/internal/cache"
)

// Store wraps a Redis client for cache-aside reads.
type Store struct {
	client *goredis.Client
}

func New(url string) (*Store, error) {
	opts, err := goredis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	client := goredis.NewClient(opts)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		return nil, err
	}
	return &Store{client: client}, nil
}

func (s *Store) Get(ctx context.Context, key string) ([]byte, bool, error) {
	val, err := s.client.Get(ctx, key).Bytes()
	if err == goredis.Nil {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return val, true, nil
}

func (s *Store) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return s.client.Set(ctx, key, value, ttl).Err()
}

func (s *Store) Delete(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	return s.client.Del(ctx, keys...).Err()
}

func (s *Store) Close() error {
	return s.client.Close()
}

var _ cache.Store = (*Store)(nil)
