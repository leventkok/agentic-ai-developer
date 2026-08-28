package model

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

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
