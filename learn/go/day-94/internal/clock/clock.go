package clock

import "time"

// Clock is an injectable time source for testable time-dependent logic.
type Clock interface {
	Now() time.Time
}

// RealClock uses the system clock in production.
type RealClock struct{}

func (RealClock) Now() time.Time {
	return time.Now().UTC()
}
