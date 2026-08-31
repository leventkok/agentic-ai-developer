package metrics

import (
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests handled",
	}, []string{"method", "route", "status"})

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "HTTP request latency in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "route"})
)

func RecordHTTP(method, path string, status int, elapsed time.Duration) {
	route := RouteLabel(path)
	code := strconv.Itoa(status)
	httpRequestsTotal.WithLabelValues(method, route, code).Inc()
	httpRequestDuration.WithLabelValues(method, route).Observe(elapsed.Seconds())
}

func RouteLabel(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	for i, part := range parts {
		if _, err := strconv.Atoi(part); err == nil {
			parts[i] = "{id}"
		}
	}
	if len(parts) == 0 || parts[0] == "" {
		return "/"
	}
	return "/" + strings.Join(parts, "/")
}
