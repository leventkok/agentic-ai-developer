ALTER TABLE users ADD COLUMN role TEXT NOT NULL DEFAULT 'member';

ALTER TABLE bookmarks ADD COLUMN user_id INTEGER REFERENCES users(id);

CREATE INDEX IF NOT EXISTS idx_bookmarks_user_id ON bookmarks (user_id);
