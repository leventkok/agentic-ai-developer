package memory

import (
	"context"
	"sync"
	"time"

	"learn/go/day-99/internal/cache"
)

type entry struct {
	value     []byte
	expiresAt time.Time
}

// Store is a process-local TTL cache for single-instance deployments.
type Store struct {
	mu    sync.Mutex
	items map[string]entry
}

func New() *Store {
	return &Store{items: make(map[string]entry)}
}

func (s *Store) Get(ctx context.Context, key string) ([]byte, bool, error) {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.items[key]
	if !ok || time.Now().After(e.expiresAt) {
		delete(s.items, key)
		return nil, false, nil
	}
	out := make([]byte, len(e.value))
	copy(out, e.value)
	return out, true, nil
}

func (s *Store) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	_ = ctx
	if ttl <= 0 {
		ttl = time.Minute
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	dup := make([]byte, len(value))
	copy(dup, value)
	s.items[key] = entry{value: dup, expiresAt: time.Now().Add(ttl)}
	return nil
}

func (s *Store) Delete(ctx context.Context, keys ...string) error {
	_ = ctx
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, key := range keys {
		delete(s.items, key)
	}
	return nil
}

func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = make(map[string]entry)
	return nil
}

var _ cache.Store = (*Store)(nil)
