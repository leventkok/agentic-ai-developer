package sqlite

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"learn/go/day-51/internal/auth"
	"learn/go/day-51/internal/db"
	"learn/go/day-51/internal/model"
	"learn/go/day-51/internal/repository"
)

type AuthStore struct {
	db         *sql.DB
	sessionTTL time.Duration
}

func NewAuthStore(database *sql.DB, sessionTTL time.Duration) *AuthStore {
	return &AuthStore{db: database, sessionTTL: sessionTTL}
}

func (s *AuthStore) Register(ctx context.Context, email, password string) (model.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	hash, err := auth.HashPassword(password)
	if err != nil {
		return model.User{}, fmt.Errorf("hash password: %w", err)
	}

	row := s.db.QueryRowContext(ctx, db.SQLInsertUser, email, hash)
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

	token, err := s.CreateSession(ctx, user.ID)
	if err != nil {
		return model.AuthResponse{}, err
	}

	user.PasswordHash = ""
	return model.AuthResponse{Token: token, User: user}, nil
}

func (s *AuthStore) CreateSession(ctx context.Context, userID int) (string, error) {
	token, err := newSessionToken()
	if err != nil {
		return "", err
	}
	expiresAt := time.Now().UTC().Add(s.sessionTTL).Format("2006-01-02 15:04:05")
	if _, err := s.db.ExecContext(ctx, db.SQLInsertSession, token, userID, expiresAt); err != nil {
		return "", fmt.Errorf("insert session: %w", err)
	}
	return token, nil
}

func (s *AuthStore) UserFromToken(ctx context.Context, token string) (model.User, error) {
	row := s.db.QueryRowContext(ctx, db.SQLGetUserBySession, token)
	user, err := scanUser(row)
	if errors.Is(err, sql.ErrNoRows) {
		return model.User{}, repository.ErrUnauthorized
	}
	if err != nil {
		return model.User{}, fmt.Errorf("get user by session: %w", err)
	}
	user.PasswordHash = ""
	return user, nil
}

func scanUser(scanner rowScanner) (model.User, error) {
	var u model.User
	var createdAt string
	if err := scanner.Scan(&u.ID, &u.Email, &u.PasswordHash, &createdAt); err != nil {
		return model.User{}, err
	}
	u.CreatedAt = parseSQLiteTime(createdAt)
	return u, nil
}

func newSessionToken() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}
	return hex.EncodeToString(buf), nil
}

func isUniqueViolation(err error) bool {
	return strings.Contains(err.Error(), "UNIQUE constraint failed")
}
