package worker

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"learn/go/day-83/internal/app"
	"learn/go/day-83/internal/config"
	"learn/go/day-83/internal/messaging"
	applog "learn/go/day-83/internal/observability/log"
)

func Run(ctx context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	logger := applog.New(cfg.Env)

	bus, err := app.OpenBus(cfg)
	if err != nil {
		return err
	}
	defer bus.Close()

	handler := NewBookmarkHandler(logger)
	wrapped := WithRetry(handler.Handle, 3)

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
	logger.Info("worker stopped", "processed", handler.ProcessedCount())
	return nil
}
