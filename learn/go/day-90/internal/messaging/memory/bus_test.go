package memory_test

import (
	"context"
	"sync/atomic"
	"testing"

	"learn/go/day-90/internal/messaging"
	"learn/go/day-90/internal/messaging/memory"
)

func TestBusPublishSubscribe(t *testing.T) {
	bus := memory.NewBus()
	var count atomic.Int32

	if err := bus.Subscribe(messaging.EventBookmarkCreated, func(ctx context.Context, evt messaging.Event) error {
		count.Add(1)
		if evt.BookmarkID != 7 {
			t.Fatalf("bookmark id = %d", evt.BookmarkID)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	evt := messaging.NewEvent(messaging.EventBookmarkCreated, 7, 1)
	if err := bus.Publish(context.Background(), evt); err != nil {
		t.Fatal(err)
	}
	if count.Load() != 1 {
		t.Fatalf("handler calls = %d", count.Load())
	}
}
