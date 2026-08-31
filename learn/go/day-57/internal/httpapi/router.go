package httpapi

import (
	"net/http"
	"time"

	"learn/go/day-57/internal/config"
	"learn/go/day-57/internal/domain"
	"learn/go/day-57/internal/middleware"
	"learn/go/day-57/internal/repository"
	"learn/go/day-57/internal/service"
)

type Deps struct {
	Cfg       config.Config
	Bookmarks repository.Bookmarks
	Auth      repository.Auth
}

func NewRouter(d Deps) http.Handler {
	bookmarkSvc := service.NewBookmarkService(d.Bookmarks, time.Duration(d.Cfg.ListTimeoutMS)*time.Millisecond)
	authSvc := service.NewAuthService(d.Auth)

	bookmarks := NewBookmarkHandler(bookmarkSvc)
	authH := NewAuthHandler(authSvc)

	requireAuth := middleware.RequireAuth(d.Auth)
	requireAdmin := middleware.RequireRole(domain.RoleAdmin)
	rateLimit := middleware.NewRateLimiter(d.Cfg.AuthRateLimitPerMinute, time.Minute)

	mux := http.NewServeMux()
	mux.Handle("POST /auth/register", rateLimit.Limit(http.HandlerFunc(authH.Register)))
	mux.Handle("POST /auth/login", rateLimit.Limit(http.HandlerFunc(authH.Login)))
	mux.Handle("GET /auth/me", requireAuth(http.HandlerFunc(authH.Me)))

	mux.HandleFunc("GET /bookmarks", bookmarks.ListBookmarks)
	mux.HandleFunc("GET /bookmarks/{id}", bookmarks.GetBookmark)
	mux.Handle("POST /bookmarks", requireAuth(http.HandlerFunc(bookmarks.CreateBookmark)))
	mux.Handle("POST /bookmarks/bulk", requireAuth(requireAdmin(http.HandlerFunc(bookmarks.BulkCreateBookmarks))))
	mux.Handle("PATCH /bookmarks/{id}", requireAuth(http.HandlerFunc(bookmarks.UpdateBookmark)))
	mux.Handle("DELETE /bookmarks/{id}", requireAuth(http.HandlerFunc(bookmarks.DeleteBookmark)))

	return middleware.DefaultStack(d.Cfg, d.Bookmarks, mux)
}
