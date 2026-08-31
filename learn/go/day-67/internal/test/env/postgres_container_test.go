//go:build integration

package env_test

import (
	"testing"

	"learn/go/day-67/internal/test/env"
)

func TestIntegration_PostgresContainerLifecycle(t *testing.T) {
	connStr := env.PostgresContainer(t)
	env.PingPostgres(t, connStr)
}
