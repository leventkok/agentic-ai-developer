package db

// Named SQL queries — single source of truth for the data layer.
const (
	SQLListBookmarks = `
		SELECT id, title, url, tags, created_at, updated_at
		FROM bookmarks ORDER BY id`

	SQLGetBookmarkByID = `
		SELECT id, title, url, tags, created_at, updated_at
		FROM bookmarks WHERE id = ?`

	SQLInsertBookmark = `
		INSERT INTO bookmarks (title, url, tags, created_at, updated_at)
		VALUES (?, ?, ?, datetime('now'), datetime('now'))
		RETURNING id, title, url, tags, created_at, updated_at`

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
		SELECT b.id, b.title, b.url, b.tags, b.created_at, b.updated_at,
		       COALESCE(a.cnt, 0) AS audit_count
		FROM bookmarks b
		LEFT JOIN (
			SELECT bookmark_id, COUNT(*) AS cnt
			FROM bookmark_audit GROUP BY bookmark_id
		) a ON a.bookmark_id = b.id
		ORDER BY b.id`

	SQLInsertUser = `
		INSERT INTO users (email, password_hash, created_at)
		VALUES (?, ?, datetime('now'))
		RETURNING id, email, password_hash, created_at`

	SQLGetUserByEmail = `
		SELECT id, email, password_hash, created_at
		FROM users WHERE email = ?`

	SQLInsertSession = `
		INSERT INTO sessions (token, user_id, expires_at)
		VALUES (?, ?, ?)`

	SQLGetUserBySession = `
		SELECT u.id, u.email, u.password_hash, u.created_at
		FROM users u
		INNER JOIN sessions s ON s.user_id = u.id
		WHERE s.token = ? AND s.expires_at > datetime('now')`
)
