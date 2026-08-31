DROP INDEX IF EXISTS idx_bookmarks_user_id;

-- SQLite does not support DROP COLUMN before 3.35; recreate would be needed for full rollback.
-- For learning migrations we document the limitation.
-- bookmarks.user_id and users.role remain if rolling back only this file.
