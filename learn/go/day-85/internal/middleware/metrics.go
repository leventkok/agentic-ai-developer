package middleware

import (
	"net/http"
	"time"

	"learn/go/day-85/internal/observability/metrics"
)

func Prometheus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		metrics.RecordHTTP(r.Method, r.URL.Path, rec.status, time.Since(start))
	})
}
