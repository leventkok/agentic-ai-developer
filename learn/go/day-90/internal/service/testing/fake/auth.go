package fake

import (
	"context"
	"strings"
	"sync"

	"learn/go/day-90/internal/domain"
)

type Auth struct {
	mu           sync.Mutex
	users        map[string]domain.User
	registerCalls int
	lastEmail    string
	lastPassword string
}

func NewAuth() *Auth {
	return &Auth{users: make(map[string]domain.User)}
}

func (a *Auth) Register(ctx context.Context, email, password string) (domain.User, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.registerCalls++
	a.lastEmail = email
	a.lastPassword = password
	email = strings.TrimSpace(strings.ToLower(email))
	if _, exists := a.users[email]; exists {
		return domain.User{}, domain.ErrDuplicateEmail
	}
	user := domain.User{ID: len(a.users) + 1, Email: email, Role: domain.RoleMember}
	a.users[email] = user
	return user, nil
}

func (a *Auth) Login(ctx context.Context, email, password string) (domain.AuthResponse, error) {
	return domain.AuthResponse{Token: "fake-token", User: domain.User{Email: email}}, nil
}

func (a *Auth) UserFromToken(ctx context.Context, token string) (domain.User, error) {
	if token == "" {
		return domain.User{}, domain.ErrUnauthorized
	}
	return domain.User{ID: 1, Email: "a@b.com", Role: domain.RoleMember}, nil
}

func (a *Auth) RegisterCalls() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.registerCalls
}
