package validation

import (
	"strconv"
	"strings"
)

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
