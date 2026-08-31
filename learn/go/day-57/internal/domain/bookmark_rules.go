package domain

func CanModifyBookmark(actor User, bookmark Bookmark) bool {
	if actor.Role == RoleAdmin {
		return true
	}
	if bookmark.UserID == nil {
		return false
	}
	return *bookmark.UserID == actor.ID
}

func CanBulkCreate(actor User) bool {
	return actor.Role == RoleAdmin
}
