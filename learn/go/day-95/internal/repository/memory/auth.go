package memory

import (
	"context"
	"strings"
	"sync"
	"time"

	"learn/go/day-95/internal/auth"
	"learn/go/day-95/internal/domain"
)

type Auth struct {
	mu     sync.Mutex
	users  map[string]domain.User
	byID   map[int]domain.User
	nextID int
	tokens *auth.TokenService
}

func NewAuth(tokens *auth.TokenService) *Auth {
	return &Auth{
		users:  make(map[string]domain.User),
		byID:   make(map[int]domain.User),
		tokens: tokens,
	}
}

func (a *Auth) Register(ctx context.Context, email, password string) (domain.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	hash, err := auth.HashPassword(password)
	if err != nil {
		return domain.User{}, err
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	if _, exists := a.users[email]; exists {
		return domain.User{}, domain.ErrDuplicateEmail
	}
	a.nextID++
	user := domain.User{
		ID:        a.nextID,
		Email:     email,
		Role:      domain.RoleMember,
		CreatedAt: time.Now().UTC(),
	}
	stored := domain.User{
		ID:           user.ID,
		Email:        user.Email,
		Role:         user.Role,
		PasswordHash: hash,
		CreatedAt:    user.CreatedAt,
	}
	a.users[email] = stored
	a.byID[user.ID] = stored
	return user, nil
}

func (a *Auth) Login(ctx context.Context, email, password string) (domain.AuthResponse, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	a.mu.Lock()
	defer a.mu.Unlock()

	user, ok := a.users[email]
	if !ok {
		return domain.AuthResponse{}, domain.ErrInvalidCredentials
	}
	if err := auth.CheckPassword(user.PasswordHash, password); err != nil {
		return domain.AuthResponse{}, domain.ErrInvalidCredentials
	}

	token, err := a.tokens.Issue(user.ID, user.Email, user.Role)
	if err != nil {
		return domain.AuthResponse{}, err
	}

	user.PasswordHash = ""
	return domain.AuthResponse{Token: token, User: user}, nil
}

func (a *Auth) UserFromToken(ctx context.Context, token string) (domain.User, error) {
	claims, err := a.tokens.Parse(token)
	if err != nil {
		return domain.User{}, domain.ErrUnauthorized
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	user, ok := a.byID[claims.UserID]
	if !ok {
		return domain.User{}, domain.ErrUnauthorized
	}
	user.PasswordHash = ""
	return user, nil
}

func (a *Auth) RegisterAndLogin(ctx context.Context, email, password string) (domain.User, string, error) {
	user, err := a.Register(ctx, email, password)
	if err != nil {
		return domain.User{}, "", err
	}
	resp, err := a.Login(ctx, email, password)
	if err != nil {
		return domain.User{}, "", err
	}
	return user, resp.Token, nil
}
