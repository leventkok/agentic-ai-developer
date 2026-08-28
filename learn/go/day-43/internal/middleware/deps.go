package middleware

import (
	"net/http"
	"context"
	"learn/go/day-43/internal/config"
	"learn/go/day-43/internal/store"
)


type depsKey struct{}

type Deps struct {

	Config config.Config
	Store store.BookmarkRepository
}


func InjectDeps(cfg config.Config, s store.BookmarkRepository) func(http.Handler) http.Handler{
	deps := Deps{
		Config: cfg,
		Store: s,
	}
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			ctx := context.WithValue(r.Context(), depsKey{}, deps)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}


func DepsFromContext(ctx context.Context) (Deps, bool){
	deps, ok := ctx.Value(depsKey{}).(Deps)
	return deps, ok
}

