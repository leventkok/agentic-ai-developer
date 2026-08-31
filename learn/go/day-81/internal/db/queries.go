package db

// Named SQL queries — single source of truth for the data layer.
const (
	SQLListBookmarks = `
		SELECT id, user_id, title, url, tags, created_at, updated_at
		FROM bookmarks ORDER BY id`

	SQLGetBookmarkByID = `
		SELECT id, user_id, title, url, tags, created_at, updated_at
		FROM bookmarks WHERE id = ?`

	SQLInsertBookmark = `
		INSERT INTO bookmarks (user_id, title, url, tags, created_at, updated_at)
		VALUES (?, ?, ?, ?, datetime('now'), datetime('now'))
		RETURNING id, user_id, title, url, tags, created_at, updated_at`

	SQLUpdateBookmark = `
		UPDATE bookmarks
		SET title = ?, url = ?, tags = ?, updated_at = datetime('now')
		WHERE id = ?`

	SQLDeleteBookmark = `DELETE FROM bookmarks WHERE id = ?`

	SQLInsertAudit = `
		INSERT INTO bookmark_audit (bookmark_id, action, detail)
		VALUES (?, 'updated', ?)`

	SQLCountAuditByBookmark = `
		SELECT COUNT(*) FROM bookmark_audit WHERE bookmark_id = ?`

	SQLListBookmarksWithAuditCount = `
		SELECT b.id, b.user_id, b.title, b.url, b.tags, b.created_at, b.updated_at,
		       COALESCE(a.cnt, 0) AS audit_count
		FROM bookmarks b
		LEFT JOIN (
			SELECT bookmark_id, COUNT(*) AS cnt
			FROM bookmark_audit GROUP BY bookmark_id
		) a ON a.bookmark_id = b.id
		ORDER BY b.id`

	SQLInsertUser = `
		INSERT INTO users (email, password_hash, role, created_at)
		VALUES (?, ?, ?, datetime('now'))
		RETURNING id, email, role, password_hash, created_at`

	SQLGetUserByEmail = `
		SELECT id, email, role, password_hash, created_at
		FROM users WHERE email = ?`

	SQLGetUserByID = `
		SELECT id, email, role, password_hash, created_at
		FROM users WHERE id = ?`
)
