package logger

type NoopLogger struct{}

func (NoopLogger) Log(message string) {
	// Do nothing
}
