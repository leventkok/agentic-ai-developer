package memory

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"strings"
	"sync"
	"time"

	"learn/go/day-51/internal/auth"
	"learn/go/day-51/internal/model"
	"learn/go/day-51/internal/repository"
)

type Auth struct {
	mu       sync.Mutex
	users    map[string]model.User
	sessions map[string]session
	nextID   int
}

type session struct {
	userID    int
	expiresAt time.Time
}

func NewAuth() *Auth {
	return &Auth{
		users:    make(map[string]model.User),
		sessions: make(map[string]session),
	}
}

func (a *Auth) Register(ctx context.Context, email, password string) (model.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	hash, err := auth.HashPassword(password)
	if err != nil {
		return model.User{}, err
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	if _, exists := a.users[email]; exists {
		return model.User{}, repository.ErrDuplicateEmail
	}
	a.nextID++
	user := model.User{
		ID:        a.nextID,
		Email:     email,
		CreatedAt: time.Now().UTC(),
	}
	a.users[email] = user
	a.users[email] = model.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: hash,
		CreatedAt:    user.CreatedAt,
	}
	user.PasswordHash = ""
	return user, nil
}

func (a *Auth) Login(ctx context.Context, email, password string) (model.AuthResponse, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	a.mu.Lock()
	defer a.mu.Unlock()

	user, ok := a.users[email]
	if !ok {
		return model.AuthResponse{}, repository.ErrInvalidCredentials
	}
	if err := auth.CheckPassword(user.PasswordHash, password); err != nil {
		return model.AuthResponse{}, repository.ErrInvalidCredentials
	}

	token, err := newToken()
	if err != nil {
		return model.AuthResponse{}, err
	}
	a.sessions[token] = session{
		userID:    user.ID,
		expiresAt: time.Now().UTC().Add(24 * time.Hour),
	}

	user.PasswordHash = ""
	return model.AuthResponse{Token: token, User: user}, nil
}

func (a *Auth) UserFromToken(ctx context.Context, token string) (model.User, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	sess, ok := a.sessions[token]
	if !ok || time.Now().UTC().After(sess.expiresAt) {
		return model.User{}, repository.ErrUnauthorized
	}
	for _, user := range a.users {
		if user.ID == sess.userID {
			user.PasswordHash = ""
			return user, nil
		}
	}
	return model.User{}, repository.ErrUnauthorized
}

func (a *Auth) RegisterAndLogin(ctx context.Context, email, password string) (model.User, string, error) {
	user, err := a.Register(ctx, email, password)
	if err != nil {
		return model.User{}, "", err
	}
	resp, err := a.Login(ctx, email, password)
	if err != nil {
		return model.User{}, "", err
	}
	return user, resp.Token, nil
}

func newToken() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}
