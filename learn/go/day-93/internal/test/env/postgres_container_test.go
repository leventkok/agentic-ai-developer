//go:build integration

package env_test

import (
	"os/exec"
	"runtime"
	"testing"

	"learn/go/day-93/internal/test/env"
)

func TestIntegration_PostgresContainerLifecycle(t *testing.T) {
	if runtime.GOOS == "windows" {
		if err := exec.Command("docker", "info").Run(); err != nil {
			t.Skip("docker not available on this Windows host")
		}
	}

	connStr := env.PostgresContainer(t)
	env.PingPostgres(t, connStr)
}
