package validation

import (
	"net/mail"
	"strings"
)

func ValidateRegister(email, password string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return ValidationError{Message: "email is required"}
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return ValidationError{Message: "email must be valid"}
	}
	if len(password) < 8 {
		return ValidationError{Message: "password must be at least 8 characters"}
	}
	return nil
}

func ValidateLogin(email, password string) error {
	if strings.TrimSpace(email) == "" {
		return ValidationError{Message: "email is required"}
	}
	if password == "" {
		return ValidationError{Message: "password is required"}
	}
	return nil
}
