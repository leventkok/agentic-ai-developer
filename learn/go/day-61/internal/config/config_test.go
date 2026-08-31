package config_test

import (
	"testing"

	"learn/go/day-61/internal/config"
)

func TestLoad_Defaults(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("ENV", "development")

	cfg, err := config.Load()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Port != "9090" {
		t.Fatalf("Port = %q, want 9090", cfg.Port)
	}
	if cfg.DBPath != "bookmarks.db" {
		t.Fatalf("DBPath = %q, want bookmarks.db", cfg.DBPath)
	}
	if cfg.DBMaxOpenConns != 1 {
		t.Fatalf("DBMaxOpenConns = %d, want 1", cfg.DBMaxOpenConns)
	}
}
