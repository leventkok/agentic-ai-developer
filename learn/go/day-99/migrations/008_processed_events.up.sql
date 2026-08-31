CREATE TABLE IF NOT EXISTS processed_events (
    dedup_key TEXT PRIMARY KEY,
    processed_at TEXT NOT NULL DEFAULT (datetime('now'))
);
