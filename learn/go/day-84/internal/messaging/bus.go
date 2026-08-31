package messaging

import "context"

type Handler func(ctx context.Context, evt Event) error

type Bus interface {
	Publish(ctx context.Context, evt Event) error
	Subscribe(eventType EventType, handler Handler) error
	Close() error
}
