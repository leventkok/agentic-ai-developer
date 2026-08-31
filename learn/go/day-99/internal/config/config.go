package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"learn/go/day-99/internal/db"
)

type Config struct {
	Port                   string
	GRPCPort               string
	ListTimeoutMS          int
	Env                    string
	ReadTimeoutSec         int
	WriteTimeoutSec        int
	ShutdownTimeoutSec     int
	DBPath                 string
	DBMaxOpenConns         int
	DBMaxIdleConns         int
	DBConnMaxLifetimeMin   int
	JWTSecret              string
	JWTTTLHours            int
	AuthRateLimitPerMinute int
	RedisURL               string
	CacheTTLSeconds        int
	NATSURL                string
	NATSQueueGroup         string
	PPROFPort              string
	OutboxWorkers          int
}

func Load() (Config, error) {
	_ = loadDotEnv(".env")

	cfg := Config{
		Port:                   getEnv("PORT", "8080"),
		GRPCPort:               getEnv("GRPC_PORT", "9090"),
		ListTimeoutMS:          getEnvInt("LIST_TIMEOUT_MS", 100),
		Env:                    getEnv("ENV", "development"),
		ReadTimeoutSec:         getEnvInt("READ_TIMEOUT_SEC", 5),
		WriteTimeoutSec:        getEnvInt("WRITE_TIMEOUT_SEC", 10),
		ShutdownTimeoutSec:     getEnvInt("SHUTDOWN_TIMEOUT_SEC", 15),
		DBPath:                 getEnv("DB_PATH", "bookmarks.db"),
		DBMaxOpenConns:         getEnvInt("DB_MAX_OPEN_CONNS", 1),
		DBMaxIdleConns:         getEnvInt("DB_MAX_IDLE_CONNS", 1),
		DBConnMaxLifetimeMin:   getEnvInt("DB_CONN_MAX_LIFETIME_MIN", 0),
		JWTSecret:              getEnv("JWT_SECRET", "dev-only-change-me-in-production"),
		JWTTTLHours:            getEnvInt("JWT_TTL_HOURS", 24),
		AuthRateLimitPerMinute: getEnvInt("AUTH_RATE_LIMIT_PER_MIN", 10),
		RedisURL:               getEnv("REDIS_URL", ""),
		CacheTTLSeconds:        getEnvInt("CACHE_TTL_SEC", 60),
		NATSURL:                getEnv("NATS_URL", ""),
		NATSQueueGroup:         getEnv("NATS_QUEUE_GROUP", "bookmarks-worker"),
		PPROFPort:              getEnv("PPROF_PORT", "6060"),
		OutboxWorkers:          getEnvInt("OUTBOX_WORKERS", 4),
	}
	return cfg, cfg.validate()
}

func (c Config) PoolConfig() db.PoolConfig {
	return db.PoolConfig{
		MaxOpenConns:    c.DBMaxOpenConns,
		MaxIdleConns:    c.DBMaxIdleConns,
		ConnMaxLifetime: time.Duration(c.DBConnMaxLifetimeMin) * time.Minute,
	}
}

func (c Config) JWTTTL() time.Duration {
	return time.Duration(c.JWTTTLHours) * time.Hour
}

func (c Config) CacheTTL() time.Duration {
	if c.CacheTTLSeconds <= 0 {
		return time.Minute
	}
	return time.Duration(c.CacheTTLSeconds) * time.Second
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	raw := os.Getenv(key)
	if raw == "" {
		return fallback
	}
	n, err := strconv.Atoi(raw)
	if err != nil {
		return fallback
	}
	return n
}

func loadDotEnv(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		if os.Getenv(key) == "" {
			os.Setenv(key, val)
		}
	}
	return scanner.Err()
}

func (c Config) validate() error {
	if strings.TrimSpace(c.Port) == "" {
		return fmt.Errorf("config: PORT is required")
	}
	if strings.TrimSpace(c.GRPCPort) == "" {
		return fmt.Errorf("config: GRPC_PORT is required")
	}
	if c.ListTimeoutMS <= 0 {
		return fmt.Errorf("config: LIST_TIMEOUT_MS must be positive, got %d", c.ListTimeoutMS)
	}
	switch c.Env {
	case "development", "staging", "production":
	default:
		return fmt.Errorf("config: ENV must be one of development, staging, production, got %q", c.Env)
	}
	if c.ReadTimeoutSec <= 0 {
		return fmt.Errorf("config: READ_TIMEOUT_SEC must be positive, got %d", c.ReadTimeoutSec)
	}
	if c.WriteTimeoutSec <= 0 {
		return fmt.Errorf("config: WRITE_TIMEOUT_SEC must be positive, got %d", c.WriteTimeoutSec)
	}
	if c.ShutdownTimeoutSec <= 0 {
		return fmt.Errorf("config: SHUTDOWN_TIMEOUT_SEC must be positive, got %d", c.ShutdownTimeoutSec)
	}
	if strings.TrimSpace(c.DBPath) == "" {
		return fmt.Errorf("config: DB_PATH is required")
	}
	if c.DBMaxOpenConns <= 0 {
		return fmt.Errorf("config: DB_MAX_OPEN_CONNS must be positive, got %d", c.DBMaxOpenConns)
	}
	if c.DBMaxIdleConns < 0 {
		return fmt.Errorf("config: DB_MAX_IDLE_CONNS must be >= 0, got %d", c.DBMaxIdleConns)
	}
	if strings.TrimSpace(c.JWTSecret) == "" {
		return fmt.Errorf("config: JWT_SECRET is required")
	}
	if c.Env == "production" && c.JWTSecret == "dev-only-change-me-in-production" {
		return fmt.Errorf("config: JWT_SECRET must be changed in production")
	}
	if c.JWTTTLHours <= 0 {
		return fmt.Errorf("config: JWT_TTL_HOURS must be positive, got %d", c.JWTTTLHours)
	}
	if c.AuthRateLimitPerMinute <= 0 {
		return fmt.Errorf("config: AUTH_RATE_LIMIT_PER_MIN must be positive, got %d", c.AuthRateLimitPerMinute)
	}
	return nil
}
