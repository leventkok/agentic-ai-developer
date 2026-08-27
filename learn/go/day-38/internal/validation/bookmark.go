package validation

import (
	"net/url"
	"strconv"
	"strings"

	"learn/go/day-38/internal/model"
)

const MaxTitleLen = 100
const MaxURLLen = 2048

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func ParseID(raw string) (int, error) {
	if strings.TrimSpace(raw) == "" {
		return 0, ValidationError{Message: "invalid id"}
	}
	id, err := strconv.Atoi(raw)
	if err != nil || id <= 0 {
		return 0, ValidationError{Message: "invalid id"}
	}
	return id, nil
}

func ValidateCreate(req model.CreateBookmarkRequest) error {
	if err := requireNonEmpty("title", req.Title); err != nil {
		return err
	}
	if err := requireNonEmpty("url", req.URL); err != nil {
		return err
	}
	if err := requireMaxLen("title", req.Title, MaxTitleLen); err != nil {
		return err
	}
	if err := requireMaxLen("url", req.URL, MaxURLLen); err != nil {
		return err
	}
	return requireURL(req.URL)
}

func ValidateUpdate(req model.UpdateBookmarkRequest) error {
	if req.Title != nil {
		if err := requireNonEmpty("title", *req.Title); err != nil {
			return err
		}
		if err := requireMaxLen("title", *req.Title, MaxTitleLen); err != nil {
			return err
		}
	}
	if req.URL != nil {
		if err := requireMaxLen("url", *req.URL, MaxURLLen); err != nil {
			return err
		}
		if err := requireURL(*req.URL); err != nil {
			return err
		}
	}
	return nil
}

func requireNonEmpty(field, value string) error {
	if strings.TrimSpace(value) == "" {
		return ValidationError{Message: field + " is required"}
	}
	return nil
}

func requireMaxLen(field, value string, max int) error {
	if len(value) > max {
		return ValidationError{Message: field + " is too long"}
	}
	return nil
}

func requireURL(raw string) error {
	u, err := url.ParseRequestURI(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return ValidationError{Message: "url must be a valid http or https URL"}
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return ValidationError{Message: "url must start with http:// or https://"}
	}
	return nil
}
