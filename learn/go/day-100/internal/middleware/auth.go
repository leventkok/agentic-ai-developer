package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"learn/go/day-100/internal/ctxkey"
	"learn/go/day-100/internal/repository"
)

func RequireAuth(auth repository.Auth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := bearerToken(r.Header.Get("Authorization"))
			if token == "" {
				writeAuthProblem(w, http.StatusUnauthorized, "authentication required")
				return
			}

			user, err := auth.UserFromToken(r.Context(), token)
			if err != nil {
				writeAuthProblem(w, http.StatusUnauthorized, "invalid or expired token")
				return
			}

			ctx := ctxkey.WithUser(r.Context(), user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func bearerToken(header string) string {
	const prefix = "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return ""
	}
	return strings.TrimSpace(strings.TrimPrefix(header, prefix))
}

func writeAuthProblem(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"code": status, "message": message})
}
