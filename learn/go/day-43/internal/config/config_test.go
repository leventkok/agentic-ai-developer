package config_test

import (
	"testing"

	"learn/go/day-43/internal/config"
)

func TestLoad_Defaults(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("LIST_TIMEOUT_MS", "")
	t.Setenv("ENV", "")

	cfg, err := config.Load()
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Port != "8080" || cfg.ListTimeoutMS != 100 || cfg.Env != "development" {
		t.Fatalf("unexpected defaults: %+v", cfg)
	}
}

func TestLoad_InvalidEnv(t *testing.T) {
	t.Setenv("ENV", "banana")
	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error for invalid ENV")
	}
}
