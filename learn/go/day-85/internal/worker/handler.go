package worker

import (
	"context"
	"log/slog"
	"sync/atomic"
	"time"

	"learn/go/day-85/internal/messaging"
	"learn/go/day-85/internal/messaging/dlq"
)

type BookmarkHandler struct {
	Logger *slog.Logger
	seen   atomic.Int64
}

func NewBookmarkHandler(logger *slog.Logger) *BookmarkHandler {
	if logger == nil {
		logger = slog.Default()
	}
	return &BookmarkHandler{Logger: logger}
}

func (h *BookmarkHandler) Handle(ctx context.Context, evt messaging.Event) error {
	h.seen.Add(1)
	h.Logger.Info("bookmark event",
		"event_id", evt.ID,
		"type", evt.Type,
		"bookmark_id", evt.BookmarkID,
		"user_id", evt.UserID,
		"occurred_at", evt.OccurredAt.Format(time.RFC3339),
	)
	return nil
}

func (h *BookmarkHandler) ProcessedCount() int64 {
	return h.seen.Load()
}

func WithDLQ(queue *dlq.Memory, handler messaging.Handler) messaging.Handler {
	return func(ctx context.Context, evt messaging.Event) error {
		if err := handler(ctx, evt); err != nil {
			queue.Push(evt)
			return err
		}
		return nil
	}
}

func WithRetryAndDLQ(queue *dlq.Memory, handler messaging.Handler, maxAttempts int) messaging.Handler {
	return WithDLQ(queue, WithRetry(handler, maxAttempts))
}

func WithRetry(handler messaging.Handler, maxAttempts int) messaging.Handler {
	if maxAttempts <= 0 {
		maxAttempts = 3
	}
	return func(ctx context.Context, evt messaging.Event) error {
		var last error
		for attempt := 1; attempt <= maxAttempts; attempt++ {
			if err := handler(ctx, evt); err == nil {
				return nil
			} else {
				last = err
				time.Sleep(time.Duration(attempt*attempt) * 50 * time.Millisecond)
			}
		}
		return last
	}
}

