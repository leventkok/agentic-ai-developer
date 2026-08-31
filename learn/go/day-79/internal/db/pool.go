package db

import (
	"database/sql"
	"time"
)

type PoolConfig struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func DefaultPoolConfig() PoolConfig {
	return PoolConfig{
		MaxOpenConns:    1, // SQLite single-writer
		MaxIdleConns:    1,
		ConnMaxLifetime: 0,
	}
}

func ConfigurePool(database *sql.DB, cfg PoolConfig) {
	if cfg.MaxOpenConns > 0 {
		database.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.MaxIdleConns >= 0 {
		database.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.ConnMaxLifetime > 0 {
		database.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}
}

// PoolStats summarizes connection pool usage for logging.
type PoolStats struct {
	OpenConnections int
	InUse           int
	Idle            int
	WaitCount       int64
	WaitDuration    time.Duration
}

func CollectPoolStats(database *sql.DB) PoolStats {
	s := database.Stats()
	return PoolStats{
		OpenConnections: s.OpenConnections,
		InUse:           s.InUse,
		Idle:            s.Idle,
		WaitCount:       s.WaitCount,
		WaitDuration:    s.WaitDuration,
	}
}
