package repository

import (
	"context"
	"learn/go/day-75/internal/domain"
)



type Auth interface {
	Register(ctx context.Context, email, password string) (domain.User, error)
	Login(ctx context.Context, email, password string) (domain.AuthResponse, error)
	UserFromToken(ctx context.Context, token string) (domain.User, error)
}
