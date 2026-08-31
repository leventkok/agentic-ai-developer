package repository

import (
	"context"
	"errors"

	"learn/go/day-55/internal/model"
)

var (
	ErrDuplicateEmail     = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
)

type Auth interface {
	Register(ctx context.Context, email, password string) (model.User, error)
	Login(ctx context.Context, email, password string) (model.AuthResponse, error)
	UserFromToken(ctx context.Context, token string) (model.User, error)
}
