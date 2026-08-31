package relay

import (
	"context"
	"log/slog"
	"time"

	"learn/go/day-99/internal/concurrency"
	"learn/go/day-99/internal/messaging"
	"learn/go/day-99/internal/messaging/outbox"
)

type Relay struct {
	Outbox   *outbox.Store
	Bus      messaging.Bus
	Interval time.Duration
	Logger   *slog.Logger
	Workers  int
}

func (r *Relay) Run(ctx context.Context) {
	interval := r.Interval
	if interval <= 0 {
		interval = 500 * time.Millisecond
	}
	logger := r.Logger
	if logger == nil {
		logger = slog.Default()
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		r.flush(ctx, logger)
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

func (r *Relay) flush(ctx context.Context, logger *slog.Logger) {
	records, err := r.Outbox.FetchUnpublished(ctx, 20)
	if err != nil {
		logger.Error("outbox fetch failed", "err", err)
		return
	}
	if len(records) == 0 {
		return
	}

	workers := r.Workers
	if workers <= 0 {
		workers = 4
	}

	err = concurrency.RunLimited(ctx, workers, records, func(ctx context.Context, rec outbox.Record) error {
		evt, err := messaging.UnmarshalEvent(rec.Payload)
		if err != nil {
			logger.Error("outbox decode failed", "id", rec.ID, "err", err)
			return nil
		}
		if err := r.Bus.Publish(ctx, evt); err != nil {
			logger.Error("outbox publish failed", "id", rec.ID, "err", err)
			return err
		}
		if err := r.Outbox.MarkPublished(ctx, rec.ID); err != nil {
			logger.Error("outbox mark published failed", "id", rec.ID, "err", err)
			return err
		}
		return nil
	})
	if err != nil {
		logger.Error("outbox flush", "err", err)
	}
}
