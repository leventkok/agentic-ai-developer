package model

type CreateBookmarkRequest struct {
	Title string `json:"title"`
	URL string `json:"url"`
	Tags []string `json:"tags,omitempty"`
}


type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}



