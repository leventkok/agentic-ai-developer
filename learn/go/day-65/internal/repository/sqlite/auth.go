package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"learn/go/day-65/internal/auth"
	"learn/go/day-65/internal/db"
	"learn/go/day-65/internal/domain"
)

type AuthStore struct {
	db     *sql.DB
	tokens *auth.TokenService
}

func NewAuthStore(database *sql.DB, tokens *auth.TokenService) *AuthStore {
	return &AuthStore{db: database, tokens: tokens}
}

func (s *AuthStore) Register(ctx context.Context, email, password string) (domain.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	hash, err := auth.HashPassword(password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}

	row := s.db.QueryRowContext(ctx, db.SQLInsertUser, email, hash, domain.RoleMember)
	user, err := scanUser(row)
	if err != nil {
		if isUniqueViolation(err) {
			return domain.User{}, domain.ErrDuplicateEmail
		}
		return domain.User{}, fmt.Errorf("insert user: %w", err)
	}
	user.PasswordHash = ""
	return user, nil
}

func (s *AuthStore) Login(ctx context.Context, email, password string) (domain.AuthResponse, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	row := s.db.QueryRowContext(ctx, db.SQLGetUserByEmail, email)
	user, err := scanUser(row)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.AuthResponse{}, domain.ErrInvalidCredentials
	}
	if err != nil {
		return domain.AuthResponse{}, fmt.Errorf("get user by email: %w", err)
	}

	if err := auth.CheckPassword(user.PasswordHash, password); err != nil {
		return domain.AuthResponse{}, domain.ErrInvalidCredentials
	}

	token, err := s.tokens.Issue(user.ID, user.Email, user.Role)
	if err != nil {
		return domain.AuthResponse{}, fmt.Errorf("issue token: %w", err)
	}

	user.PasswordHash = ""
	return domain.AuthResponse{Token: token, User: user}, nil
}

func (s *AuthStore) UserFromToken(ctx context.Context, token string) (domain.User, error) {
	claims, err := s.tokens.Parse(token)
	if err != nil {
		return domain.User{}, domain.ErrUnauthorized
	}

	row := s.db.QueryRowContext(ctx, db.SQLGetUserByID, claims.UserID)
	user, err := scanUser(row)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, domain.ErrUnauthorized
	}
	if err != nil {
		return domain.User{}, fmt.Errorf("get user by id: %w", err)
	}
	user.PasswordHash = ""
	return user, nil
}

func scanUser(scanner rowScanner) (domain.User, error) {
	var u domain.User
	var createdAt string
	if err := scanner.Scan(&u.ID, &u.Email, &u.Role, &u.PasswordHash, &createdAt); err != nil {
		return domain.User{}, err
	}
	u.CreatedAt = parseSQLiteTime(createdAt)
	return u, nil
}

func isUniqueViolation(err error) bool {
	return strings.Contains(err.Error(), "UNIQUE constraint failed")
}
