package memory

import (
	"context"
	"strings"
	"sync"
	"time"

	"learn/go/day-55/internal/auth"
	"learn/go/day-55/internal/model"
	"learn/go/day-55/internal/repository"
)

type Auth struct {
	mu     sync.Mutex
	users  map[string]model.User
	byID   map[int]model.User
	nextID int
	tokens *auth.TokenService
}

func NewAuth(tokens *auth.TokenService) *Auth {
	return &Auth{
		users:  make(map[string]model.User),
		byID:   make(map[int]model.User),
		tokens: tokens,
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
		Role:      auth.RoleMember,
		CreatedAt: time.Now().UTC(),
	}
	stored := model.User{
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

	token, err := a.tokens.Issue(user.ID, user.Email, user.Role)
	if err != nil {
		return model.AuthResponse{}, err
	}

	user.PasswordHash = ""
	return model.AuthResponse{Token: token, User: user}, nil
}

func (a *Auth) UserFromToken(ctx context.Context, token string) (model.User, error) {
	claims, err := a.tokens.Parse(token)
	if err != nil {
		return model.User{}, repository.ErrUnauthorized
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	user, ok := a.byID[claims.UserID]
	if !ok {
		return model.User{}, repository.ErrUnauthorized
	}
	user.PasswordHash = ""
	return user, nil
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

func (a *Auth) SeedAdmin(email, password string) (model.User, error) {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return model.User{}, err
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.nextID++
	user := model.User{
		ID:           a.nextID,
		Email:        email,
		Role:         auth.RoleAdmin,
		PasswordHash: hash,
		CreatedAt:    time.Now().UTC(),
	}
	a.users[email] = user
	a.byID[user.ID] = user
	user.PasswordHash = ""
	return user, nil
}
