package httpapi

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateBookmarkRequest struct {
	Title string   `json:"title"`
	URL   string   `json:"url"`
	Tags  []string `json:"tags,omitempty"`
}

type UpdateBookmarkRequest struct {
	Title *string  `json:"title,omitempty"`
	URL   *string  `json:"url,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

type BulkCreateBookmarksRequest struct {
	Bookmarks []CreateBookmarkRequest `json:"bookmarks"`
}