package middleware

import (
	"net/http"

	"learn/go/day-87/internal/auth"
	"learn/go/day-87/internal/ctxkey"
)

func RequireRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := ctxkey.UserFromContext(r.Context())
			if !ok {
				writeAuthProblem(w, http.StatusUnauthorized, "authentication required")
				return
			}
			if !auth.HasRole(user.Role, roles...) {
				writeAuthProblem(w, http.StatusForbidden, "forbidden")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
