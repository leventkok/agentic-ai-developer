package metrics_test

import (
	"testing"

	"learn/go/day-87/internal/observability/metrics"
)

func TestRouteLabel_NormalizesNumericIDs(t *testing.T) {
	got := metrics.RouteLabel("/bookmarks/42")
	if got != "/bookmarks/{id}" {
		t.Fatalf("got %q", got)
	}
}
