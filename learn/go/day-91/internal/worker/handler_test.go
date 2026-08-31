package worker_test

import (
	"context"
	"sync/atomic"
	"testing"

	"learn/go/day-91/internal/db/testutil"
	"learn/go/day-91/internal/messaging"
	"learn/go/day-91/internal/messaging/dlq"
	"learn/go/day-91/internal/messaging/idempotency"
	"learn/go/day-91/internal/worker"
)

func TestIdempotentHandlerProcessesOnce(t *testing.T) {
	tdb := testutil.OpenTestDB(t)
	store := idempotency.NewStore(tdb.DB)

	var calls atomic.Int32
	base := func(ctx context.Context, evt messaging.Event) error {
		calls.Add(1)
		return nil
	}
	handler := idempotency.Wrap(store, base)

	evt := messaging.NewEvent(messaging.EventBookmarkCreated, 42, 1)
	if err := handler(context.Background(), evt); err != nil {
		t.Fatal(err)
	}
	if err := handler(context.Background(), evt); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 1 {
		t.Fatalf("calls = %d, want 1", calls.Load())
	}
}

func TestWithRetryAndDLQ(t *testing.T) {
	dlqMem := dlq.NewMemory()
	attempts := 0
	failing := func(ctx context.Context, evt messaging.Event) error {
		attempts++
		return context.DeadlineExceeded
	}
	handler := worker.WithRetryAndDLQ(dlqMem, failing, 2)
	evt := messaging.NewEvent(messaging.EventBookmarkCreated, 1, 1)
	_ = handler(context.Background(), evt)
	if attempts != 2 {
		t.Fatalf("attempts = %d, want 2", attempts)
	}
	if dlqMem.Len() != 1 {
		t.Fatalf("dlq len = %d, want 1", dlqMem.Len())
	}
}
