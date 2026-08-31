package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

// ExplainQueryPlan returns SQLite EXPLAIN QUERY PLAN output for debugging slow queries.
func ExplainQueryPlan(ctx context.Context, database *sql.DB, query string, args ...any) (string, error) {
	rows, err := database.QueryContext(ctx, "EXPLAIN QUERY PLAN "+query, args...)
	if err != nil {
		return "", fmt.Errorf("explain query plan: %w", err)
	}
	defer rows.Close()

	var lines []string
	for rows.Next() {
		var id, parent, notused int
		var detail string
		if err := rows.Scan(&id, &parent, &notused, &detail); err != nil {
			return "", fmt.Errorf("scan explain row: %w", err)
		}
		lines = append(lines, detail)
	}
	if err := rows.Err(); err != nil {
		return "", err
	}
	return strings.Join(lines, "\n"), nil
}
