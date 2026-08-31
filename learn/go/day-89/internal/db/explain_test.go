package db_test

import (
	"context"
	"strings"
	"testing"

	"learn/go/day-89/internal/db"
	"learn/go/day-89/internal/db/testutil"
)

func TestExplainListBookmarksUsesTableScanOrIndex(t *testing.T) {
	tdb := testutil.OpenTestDB(t)
	plan, err := db.ExplainQueryPlan(context.Background(), tdb.DB, `
		SELECT id, user_id, title, url, tags, created_at, updated_at
		FROM bookmarks ORDER BY id`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(plan, "bookmarks") {
		t.Fatalf("expected bookmarks in plan: %q", plan)
	}
}
