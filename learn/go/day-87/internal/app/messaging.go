package app

import (
	"fmt"

	"learn/go/day-87/internal/config"
	"learn/go/day-87/internal/messaging"
	"learn/go/day-87/internal/messaging/memory"
	natsbus "learn/go/day-87/internal/messaging/nats"
)

func OpenBus(cfg config.Config) (messaging.Bus, error) {
	if cfg.NATSURL != "" {
		return natsbus.New(cfg.NATSURL, cfg.NATSQueueGroup)
	}
	return memory.NewBus(), nil
}

func CloseBus(bus messaging.Bus) error {
	if bus == nil {
		return nil
	}
	return bus.Close()
}

func busCloseErr(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("bus close: %w", err)
}
