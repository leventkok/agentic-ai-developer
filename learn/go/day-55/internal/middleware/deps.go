package middleware

import (
	"context"
	"net/http"

	"learn/go/day-55/internal/config"
	"learn/go/day-55/internal/repository"
)

type depsKey struct{}

type Deps struct {
	Config config.Config
	Repo   repository.Bookmarks
}

func InjectDeps(cfg config.Config, repo repository.Bookmarks) func(http.Handler) http.Handler {
	deps := Deps{Config: cfg, Repo: repo}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), depsKey{}, deps)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func DepsFromContext(ctx context.Context) (Deps, bool) {
	deps, ok := ctx.Value(depsKey{}).(Deps)
	return deps, ok
}
