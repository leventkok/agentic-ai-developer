package middleware

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// Tracing wraps handlers with OpenTelemetry HTTP spans.
func Tracing(next http.Handler) http.Handler {
	return otelhttp.NewHandler(next, "bookmarks-http")
}
