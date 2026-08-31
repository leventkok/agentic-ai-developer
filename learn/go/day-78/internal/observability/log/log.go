package log

import (
	"log/slog"
	"os"
	"strings"
)

// New returns a JSON structured logger for the given environment.
func New(env string) *slog.Logger {
	level := slog.LevelInfo
	if strings.EqualFold(env, "development") {
		level = slog.LevelDebug
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
}
