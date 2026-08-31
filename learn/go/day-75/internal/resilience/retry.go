package resilience

import (
	"context"
	"time"
)

// Retry runs fn up to attempts times with exponential backoff.
// Only use for idempotent read paths.
func Retry(ctx context.Context, attempts int, baseDelay time.Duration, fn func() error) error {
	if attempts < 1 {
		attempts = 1
	}

	var err error
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return nil
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if i+1 >= attempts {
			break
		}
		delay := baseDelay * time.Duration(1<<i)
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
	}
	return err
}
