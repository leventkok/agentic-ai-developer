package env_test

import (
	"testing"

	"learn/go/day-93/internal/db/testutil"
)

// Each parallel subtest gets its own temp database file via t.TempDir().
func TestParallel_IsolatedDatabases(t *testing.T) {
	for range 3 {
		t.Run("worker", func(t *testing.T) {
			t.Parallel()
			tdb := testutil.OpenTestDB(t)
			testutil.ResetTables(t, tdb)
			if tdb.Path == "" {
				t.Fatal("expected database path")
			}
		})
	}
}
