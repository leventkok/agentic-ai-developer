package middleware

import (
	"net/http"
	"learn/go/day-40/internal/config"
	"learn/go/day-40/internal/store"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recover() != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Chain(h http.Handler, wrappers ...func(http.Handler) http.Handler) http.Handler {
	for i := len(wrappers) - 1; i >= 0; i-- {
		h = wrappers[i](h)
	}
	return h
}



func DefaultStack(cfg config.Config, s store.BookmarkRepository, mux http.Handler) http.Handler{
	return Chain(
		mux,
		InjectDeps(cfg, s),
		RequestID,
		Logging,
		Recovery,
		
	)

}