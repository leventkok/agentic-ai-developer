package domain

type CreateBookmarkInput struct {
	Title string
	URL   string
	Tags  []string
}

type UpdateBookmarkInput struct {
	Title *string
	URL   *string
	Tags  []string
}
