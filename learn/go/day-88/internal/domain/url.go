package domain

import (
	"net/url"
)

type BookmarkURL struct {
	value string
}

func NewBookmarkURL(raw string) (BookmarkURL, error) {
	if len(raw) > MaxURLLen {
		return BookmarkURL{}, ValidationError{Field: "url", Message: "is too long"}
	}
	u, err := url.ParseRequestURI(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return BookmarkURL{}, ValidationError{Field: "url", Message: "must be a valid http or https URL"}
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return BookmarkURL{}, ValidationError{Field: "url", Message: "must start with http:// or https://"}
	}
	return BookmarkURL{value: raw}, nil
}

func (u BookmarkURL) String() string {
	return u.value
}
