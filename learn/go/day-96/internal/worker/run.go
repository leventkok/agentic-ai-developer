package worker

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"learn/go/day-96/internal/app"
	"learn/go/day-96/internal/config"
	"learn/go/day-96/internal/db"
	"learn/go/day-96/internal/messaging"
	"learn/go/day-96/internal/messaging/dlq"
	"learn/go/day-96/internal/messaging/idempotency"
	applog "learn/go/day-96/internal/observability/log"
)

func Run(ctx context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	logger := applog.New(cfg.Env)

	database, err := db.Open(cfg.DBPath, cfg.PoolConfig())
	if err != nil {
		return err
	}
	defer database.Close()
	if err := db.RunMigrations(database, "migrations"); err != nil {
		return err
	}

	bus, err := app.OpenBus(cfg)
	if err != nil {
		return err
	}
	defer bus.Close()

	idempotencyStore := idempotency.NewStore(database)
	deadLetter := dlq.NewMemory()
	handler := NewBookmarkHandler(logger)
	wrapped := WithRetryAndDLQ(deadLetter, idempotency.Wrap(idempotencyStore, handler.Handle), 3)

	for _, eventType := range []messaging.EventType{
		messaging.EventBookmarkCreated,
		messaging.EventBookmarkUpdated,
	} {
		if err := bus.Subscribe(eventType, wrapped); err != nil {
			return fmt.Errorf("subscribe %s: %w", eventType, err)
		}
	}

	logger.Info("worker started", "nats_url", cfg.NATSURL)

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	logger.Info("worker stopped",
		"processed", handler.ProcessedCount(),
		"dlq", deadLetter.Len(),
	)
	return nil
}
