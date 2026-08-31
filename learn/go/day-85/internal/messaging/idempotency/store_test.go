package idempotency_test

import (
	"context"
	"sync/atomic"
	"testing"

	"learn/go/day-85/internal/db/testutil"
	"learn/go/day-85/internal/messaging"
	"learn/go/day-85/internal/messaging/idempotency"
)

func TestWrapSkipsDuplicateDelivery(t *testing.T) {
	tdb := testutil.OpenTestDB(t)
	store := idempotency.NewStore(tdb.DB)
	var calls atomic.Int32
	handler := idempotency.Wrap(store, func(ctx context.Context, evt messaging.Event) error {
		calls.Add(1)
		return nil
	})

	evt := messaging.NewEvent(messaging.EventBookmarkCreated, 1, 1)
	if err := handler(context.Background(), evt); err != nil {
		t.Fatal(err)
	}
	if err := handler(context.Background(), evt); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 1 {
		t.Fatalf("handler calls = %d, want 1", calls.Load())
	}
}
