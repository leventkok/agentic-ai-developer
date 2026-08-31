package domain

import "time"

const (
	RoleAdmin  = "admin"
	RoleMember = "member"
)

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
