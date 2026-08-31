package handler

import (
	"context"

	"learn/go/day-55/internal/model"
)

type userContextKey struct{}

func WithUser(ctx context.Context, user model.User) context.Context {
	return context.WithValue(ctx, userContextKey{}, user)
}

func UserFromContext(ctx context.Context) (model.User, bool) {
	user, ok := ctx.Value(userContextKey{}).(model.User)
	return user, ok
}
