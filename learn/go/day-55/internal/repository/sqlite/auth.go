package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"learn/go/day-55/internal/auth"
	"learn/go/day-55/internal/db"
	"learn/go/day-55/internal/model"
	"learn/go/day-55/internal/repository"
)

type AuthStore struct {
	db    *sql.DB
	tokens *auth.TokenService
}

func NewAuthStore(database *sql.DB, tokens *auth.TokenService) *AuthStore {
	return &AuthStore{db: database, tokens: tokens}
}

func (s *AuthStore) Register(ctx context.Context, email, password string) (model.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	hash, err := auth.HashPassword(password)
	if err != nil {
		return model.User{}, fmt.Errorf("hash password: %w", err)
	}

	row := s.db.QueryRowContext(ctx, db.SQLInsertUser, email, hash, auth.RoleMember)
	user, err := scanUser(row)
	if err != nil {
		if isUniqueViolation(err) {
			return model.User{}, repository.ErrDuplicateEmail
		}
		return model.User{}, fmt.Errorf("insert user: %w", err)
	}
	user.PasswordHash = ""
	return user, nil
}

func (s *AuthStore) Login(ctx context.Context, email, password string) (model.AuthResponse, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	row := s.db.QueryRowContext(ctx, db.SQLGetUserByEmail, email)
	user, err := scanUser(row)
	if errors.Is(err, sql.ErrNoRows) {
		return model.AuthResponse{}, repository.ErrInvalidCredentials
	}
	if err != nil {
		return model.AuthResponse{}, fmt.Errorf("get user by email: %w", err)
	}

	if err := auth.CheckPassword(user.PasswordHash, password); err != nil {
		return model.AuthResponse{}, repository.ErrInvalidCredentials
	}

	token, err := s.tokens.Issue(user.ID, user.Email, user.Role)
	if err != nil {
		return model.AuthResponse{}, fmt.Errorf("issue token: %w", err)
	}

	user.PasswordHash = ""
	return model.AuthResponse{Token: token, User: user}, nil
}

func (s *AuthStore) UserFromToken(ctx context.Context, token string) (model.User, error) {
	claims, err := s.tokens.Parse(token)
	if err != nil {
		return model.User{}, repository.ErrUnauthorized
	}

	row := s.db.QueryRowContext(ctx, db.SQLGetUserByID, claims.UserID)
	user, err := scanUser(row)
	if errors.Is(err, sql.ErrNoRows) {
		return model.User{}, repository.ErrUnauthorized
	}
	if err != nil {
		return model.User{}, fmt.Errorf("get user by id: %w", err)
	}
	user.PasswordHash = ""
	return user, nil
}

func scanUser(scanner rowScanner) (model.User, error) {
	var u model.User
	var createdAt string
	if err := scanner.Scan(&u.ID, &u.Email, &u.Role, &u.PasswordHash, &createdAt); err != nil {
		return model.User{}, err
	}
	u.CreatedAt = parseSQLiteTime(createdAt)
	return u, nil
}

func isUniqueViolation(err error) bool {
	return strings.Contains(err.Error(), "UNIQUE constraint failed")
}
