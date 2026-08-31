package domain

// NewBookmark validates input and returns a bookmark ready for persistence (no ID/timestamps yet).
// TODO (Day 59): Use NewTitle + NewBookmarkURL; enforce invariants before repo.Create.
func NewBookmark(in CreateBookmarkInput) (Bookmark, error) {
	panic("TODO: implement — construct value objects, return populated Bookmark without ID")
}
