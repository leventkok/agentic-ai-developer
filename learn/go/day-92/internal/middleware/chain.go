package middleware

import (
	"log/slog"
	"net/http"

	"learn/go/day-92/internal/config"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recover() != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Chain(h http.Handler, wrappers ...func(http.Handler) http.Handler) http.Handler {
	for i := len(wrappers) - 1; i >= 0; i-- {
		h = wrappers[i](h)
	}
	return h
}

func DefaultStack(cfg config.Config, logger *slog.Logger, mux http.Handler) http.Handler {
	return Chain(
		mux,
		Tracing,
		RequestID,
		Prometheus,
		func(next http.Handler) http.Handler { return Logging(logger, next) },
		Recovery,
	)
}
