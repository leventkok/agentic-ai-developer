package sqlite_test

import (
	"context"
	"path/filepath"
	"testing"

	"learn/go/day-90/internal/db"
	"learn/go/day-90/internal/repository/sqlite"
)

func openBenchStore(b *testing.B) *sqlite.Store {
	b.Helper()
	path := filepath.Join(b.TempDir(), "bench.db")
	database, err := db.Open(path, db.DefaultPoolConfig())
	if err != nil {
		b.Fatal(err)
	}
	b.Cleanup(func() { database.Close() })
	if err := db.RunMigrations(database, "../../../migrations"); err != nil {
		b.Fatal(err)
	}
	store, err := sqlite.New(database)
	if err != nil {
		b.Fatal(err)
	}
	b.Cleanup(func() { store.Close() })
	return store
}

func BenchmarkStoreList(b *testing.B) {
	store := openBenchStore(b)
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := store.List(ctx); err != nil {
			b.Fatal(err)
		}
	}
}

// Baseline guard: list on empty DB should stay under 1ms/op on typical dev hardware.
func TestListBenchmarkRegressionGuard(t *testing.T) {
	result := testing.Benchmark(BenchmarkStoreList)
	if result.NsPerOp() > 1_000_000 {
		t.Fatalf("BenchmarkStoreList regression: %d ns/op (max 1000000)", result.NsPerOp())
	}
}
