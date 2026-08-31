package memory

import (
	"context"
	"fmt"
	"sync"

	"learn/go/day-94/internal/messaging"
)

type Bus struct {
	mu       sync.RWMutex
	handlers map[messaging.EventType][]messaging.Handler
}

func NewBus() *Bus {
	return &Bus{handlers: make(map[messaging.EventType][]messaging.Handler)}
}

func (b *Bus) Publish(ctx context.Context, evt messaging.Event) error {
	b.mu.RLock()
	handlers := append([]messaging.Handler(nil), b.handlers[evt.Type]...)
	b.mu.RUnlock()

	for _, h := range handlers {
		if err := h(ctx, evt); err != nil {
			return fmt.Errorf("memory bus handler: %w", err)
		}
	}
	return nil
}

func (b *Bus) Subscribe(eventType messaging.EventType, handler messaging.Handler) error {
	if handler == nil {
		return fmt.Errorf("memory bus: nil handler")
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[eventType] = append(b.handlers[eventType], handler)
	return nil
}

func (b *Bus) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers = make(map[messaging.EventType][]messaging.Handler)
	return nil
}
