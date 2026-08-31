package validation

import (
	"net/url"
	"strconv"
	"strings"

	"learn/go/day-58/internal/domain"
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

func ValidateCreateInput(in domain.CreateBookmarkInput) error {
	return validateCreateFields(in.Title, in.URL)
}

func ValidateUpdateInput(in domain.UpdateBookmarkInput) error {
	if in.Title != nil {
		if err := requireNonEmpty("title", *in.Title); err != nil {
			return err
		}
		if err := requireMaxLen("title", *in.Title, MaxTitleLen); err != nil {
			return err
		}
	}
	if in.URL != nil {
		if err := requireMaxLen("url", *in.URL, MaxURLLen); err != nil {
			return err
		}
		if err := requireURL(*in.URL); err != nil {
			return err
		}
	}
	return nil
}

func ValidateBulkCreateInputs(inputs []domain.CreateBookmarkInput) error {
	if len(inputs) == 0 {
		return ValidationError{Message: "bookmarks must not be empty"}
	}
	for i, item := range inputs {
		if err := ValidateCreateInput(item); err != nil {
			return ValidationError{Message: "bookmarks[" + strconv.Itoa(i) + "]: " + err.Error()}
		}
	}
	return nil
}

func validateCreateFields(title, url string) error {
	if err := requireNonEmpty("title", title); err != nil {
		return err
	}
	if err := requireNonEmpty("url", url); err != nil {
		return err
	}
	if err := requireMaxLen("title", title, MaxTitleLen); err != nil {
		return err
	}
	if err := requireMaxLen("url", url, MaxURLLen); err != nil {
		return err
	}
	return requireURL(url)
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
