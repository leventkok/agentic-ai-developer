package domain

import "errors"

var (
	ErrInvalidTitle = errors.New("invalid title")
	ErrInvalidURL   = errors.New("invalid url")
	ErrInvalidEmail = errors.New("invalid email")
)

const (
	MaxTitleLen = 100
	MaxURLLen   = 2048
)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	if e.Field == "" {
		return e.Message
	}
	return e.Field + ": " + e.Message
}

func IsValidation(err error) bool {
	var ve ValidationError
	return errors.As(err, &ve) ||
		errors.Is(err, ErrInvalidTitle) ||
		errors.Is(err, ErrInvalidURL) ||
		errors.Is(err, ErrInvalidEmail)
}
