package ctxkey

import (
	"context"

	"learn/go/day-59/internal/domain"
)

type userKey struct{}

func WithUser(ctx context.Context, user domain.User) context.Context {
	return context.WithValue(ctx, userKey{}, user)
}

func UserFromContext(ctx context.Context) (domain.User, bool) {
	user, ok := ctx.Value(userKey{}).(domain.User)
	return user, ok
}
