package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port          string
	ListTimeoutMS int
	Env           string
}

func Load() (Config, error) {
	_ = loadDotEnv(".env")

	cfg := Config{
		Port:          getEnv("PORT", "8080"),
		ListTimeoutMS: getEnvInt("LIST_TIMEOUT_MS", 100),
		Env:           getEnv("ENV", "development"),
	}
	return cfg, cfg.validate()
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
	if c.ListTimeoutMS <= 0 {
		return fmt.Errorf("config: LIST_TIMEOUT_MS must be positive, got %d", c.ListTimeoutMS)
	}
	switch c.Env {
	case "development", "staging", "production":
		return nil
	default:
		return fmt.Errorf("config: ENV must be one of development, staging, production, got %q", c.Env)
	}
}
