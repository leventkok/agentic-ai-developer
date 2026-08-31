package domain

// CanModifyBookmark is a pure domain rule — no HTTP or SQL imports.
func CanModifyBookmark(actor User, bookmark Bookmark) bool {
	if actor.Role == RoleAdmin {
		return true
	}
	if bookmark.UserID == nil {
		return false
	}
	return *bookmark.UserID == actor.ID
}

// CanBulkCreate returns whether the actor may bulk-create bookmarks.
func CanBulkCreate(actor User) bool {
	return actor.Role == RoleAdmin
}
