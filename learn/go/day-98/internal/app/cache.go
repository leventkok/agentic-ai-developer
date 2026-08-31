package app

import (
	"fmt"

	"learn/go/day-98/internal/cache"
	cacheredis "learn/go/day-98/internal/cache/redis"
	cachememory "learn/go/day-98/internal/cache/memory"
	"learn/go/day-98/internal/config"
)

func openCache(cfg config.Config) (cache.Store, error) {
	if cfg.RedisURL != "" {
		return cacheredis.New(cfg.RedisURL)
	}
	return cachememory.New(), nil
}

func closeCache(store cache.Store) error {
	if store == nil {
		return nil
	}
	return store.Close()
}

func cacheCloseErr(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("cache close: %w", err)
}
