package service

import (
	"context"
	"learn/go/day-88/internal/repository"
	"learn/go/day-88/internal/domain"
)



type AuthService struct {
	Auth repository.Auth
}

func NewAuthService(auth repository.Auth) *AuthService {
	return &AuthService{Auth: auth}
}

func (s *AuthService) Register(ctx context.Context, email, password string) (domain.User, error) {
	return s.Auth.Register(ctx, email, password)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (domain.AuthResponse, error) {
	return s.Auth.Login(ctx, email, password)
}


func (s *AuthService) UserFromToken(ctx context.Context, token string) (domain.User, error) {
	return s.Auth.UserFromToken(ctx, token)
}
