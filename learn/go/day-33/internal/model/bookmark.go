package model

type Bookmark struct {
	ID    int      `json:"id"`
	Title string   `json:"title"`
	URL   string   `json:"url"`
	Tags  []string `json:"tags,omitempty"`
}
