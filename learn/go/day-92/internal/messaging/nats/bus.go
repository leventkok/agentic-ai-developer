package nats

import (
	"context"
	"fmt"
	"sync"

	natspkg "github.com/nats-io/nats.go"

	"learn/go/day-92/internal/messaging"
)

type Bus struct {
	conn     *natspkg.Conn
	queue    string
	mu       sync.Mutex
	handlers []messaging.Handler
}

func New(url, queueGroup string) (*Bus, error) {
	conn, err := natspkg.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("nats bus: connect: %w", err)
	}
	if queueGroup == "" {
		queueGroup = "bookmarks-worker"
	}
	return &Bus{conn: conn, queue: queueGroup}, nil
}

func (b *Bus) Publish(_ context.Context, evt messaging.Event) error {
	data, err := evt.Marshal()
	if err != nil {
		return err
	}
	return b.conn.Publish(string(evt.Type), data)
}

func (b *Bus) Subscribe(eventType messaging.EventType, handler messaging.Handler) error {
	if handler == nil {
		return fmt.Errorf("nats bus: nil handler")
	}
	sub, err := b.conn.QueueSubscribe(string(eventType), b.queue, func(msg *natspkg.Msg) {
		evt, err := messaging.UnmarshalEvent(msg.Data)
		if err != nil {
			return
		}
		_ = handler(context.Background(), evt)
	})
	if err != nil {
		return fmt.Errorf("nats bus subscribe: %w", err)
	}
	b.mu.Lock()
	b.handlers = append(b.handlers, handler)
	b.mu.Unlock()
	_ = sub
	return nil
}

func (b *Bus) Close() error {
	if b.conn != nil {
		b.conn.Close()
	}
	return nil
}
