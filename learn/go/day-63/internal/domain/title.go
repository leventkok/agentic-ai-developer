package domain

import "strings"

type Title struct {
	value string
}

func NewTitle(raw string) (Title, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return Title{}, ValidationError{Field: "title", Message: "is required"}
	}
	if len(trimmed) > MaxTitleLen {
		return Title{}, ValidationError{Field: "title", Message: "is too long"}
	}
	return Title{value: trimmed}, nil
}

func (t Title) String() string {
	return t.value
}
