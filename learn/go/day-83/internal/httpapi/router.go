package httpapi

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"learn/go/day-83/internal/config"
	"learn/go/day-83/internal/domain"
	"learn/go/day-83/internal/middleware"
	applog "learn/go/day-83/internal/observability/log"
	"learn/go/day-83/internal/repository"
	"learn/go/day-83/internal/service"
)

type Deps struct {
	Cfg         config.Config
	Bookmarks   repository.Bookmarks
	BookmarkSvc *service.BookmarkService
	Auth        repository.Auth
}

func NewRouter(d Deps) http.Handler {
	bookmarkSvc := d.BookmarkSvc
	if bookmarkSvc == nil {
		bookmarkSvc = service.NewBookmarkService(d.Bookmarks, time.Duration(d.Cfg.ListTimeoutMS)*time.Millisecond, nil)
	}
	authSvc := service.NewAuthService(d.Auth)

	bookmarks := NewBookmarkHandler(bookmarkSvc)
	authH := NewAuthHandler(authSvc)

	requireAuth := middleware.RequireAuth(d.Auth)
	requireAdmin := middleware.RequireRole(domain.RoleAdmin)
	rateLimit := middleware.NewRateLimiter(d.Cfg.AuthRateLimitPerMinute, time.Minute)

	mux := http.NewServeMux()
	mux.Handle("GET /metrics", promhttp.Handler())

	mux.Handle("POST /auth/register", rateLimit.Limit(http.HandlerFunc(authH.Register)))
	mux.Handle("POST /auth/login", rateLimit.Limit(http.HandlerFunc(authH.Login)))
	mux.Handle("GET /auth/me", requireAuth(http.HandlerFunc(authH.Me)))

	mux.HandleFunc("GET /bookmarks", bookmarks.ListBookmarks)
	mux.HandleFunc("GET /bookmarks/{id}", bookmarks.GetBookmark)
	mux.Handle("POST /bookmarks", requireAuth(http.HandlerFunc(bookmarks.CreateBookmark)))
	mux.Handle("POST /bookmarks/bulk", requireAuth(requireAdmin(http.HandlerFunc(bookmarks.BulkCreateBookmarks))))
	mux.Handle("PATCH /bookmarks/{id}", requireAuth(http.HandlerFunc(bookmarks.UpdateBookmark)))
	mux.Handle("DELETE /bookmarks/{id}", requireAuth(http.HandlerFunc(bookmarks.DeleteBookmark)))

	return middleware.DefaultStack(d.Cfg, applog.New(d.Cfg.Env), mux)
}
