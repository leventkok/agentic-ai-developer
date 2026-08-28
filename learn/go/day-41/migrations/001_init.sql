CREATE TABLE IF NOT EXISTS bookmarks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    tags TEXT NOT NULL DEFAULT '[]'
);

INSERT OR IGNORE INTO bookmarks (id, title, url, tags) VALUES
    (1, 'Go Docs', 'https://go.dev', '["lang","docs"]'),
    (2, 'SQLite', 'https://sqlite.org', '["database"]');
