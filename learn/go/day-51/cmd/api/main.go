package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"learn/go/day-51/internal/config"
	"learn/go/day-51/internal/db"
	"learn/go/day-51/internal/handler"
	"learn/go/day-51/internal/middleware"
	"learn/go/day-51/internal/repository/sqlite"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.Open(cfg.DBPath, cfg.PoolConfig())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.RunMigrations(database, "migrations"); err != nil {
		database.Close()
		log.Fatal(err)
	}

	repo, err := sqlite.New(database)
	if err != nil {
		database.Close()
		log.Fatal(err)
	}

	authStore := sqlite.NewAuthStore(database, cfg.SessionTTL())

	seeded, err := repo.List(context.Background())
	if err != nil {
		repo.Close()
		database.Close()
		log.Fatal(err)
	}
	fmt.Println("DB seed count:", len(seeded))

	h := &handler.BookmarkHandler{
		Repo:        repo,
		ListTimeout: time.Duration(cfg.ListTimeoutMS) * time.Millisecond,
	}
	authHandler := &handler.AuthHandler{Auth: authStore}
	requireAuth := middleware.RequireAuth(authStore)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth/register", authHandler.Register)
	mux.HandleFunc("POST /auth/login", authHandler.Login)
	mux.Handle("GET /auth/me", requireAuth(http.HandlerFunc(authHandler.Me)))

	mux.HandleFunc("GET /bookmarks", h.ListBookmarks)
	mux.HandleFunc("GET /bookmarks/{id}", h.GetBookmark)
	mux.Handle("POST /bookmarks", requireAuth(http.HandlerFunc(h.CreateBookmark)))
	mux.Handle("POST /bookmarks/bulk", requireAuth(http.HandlerFunc(h.BulkCreateBookmarks)))
	mux.Handle("PATCH /bookmarks/{id}", requireAuth(http.HandlerFunc(h.UpdateBookmark)))
	mux.Handle("DELETE /bookmarks/{id}", requireAuth(http.HandlerFunc(h.DeleteBookmark)))

	root := middleware.DefaultStack(cfg, repo, mux)
	addr := ":" + cfg.Port

	server := &http.Server{
		Addr:         addr,
		Handler:      root,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSec) * time.Second,
	}

	go func() {
		fmt.Printf("Bookmarks API [%s] on http://localhost%s\n", cfg.Env, addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeoutSec)*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown failed: %v", err)
	}

	stats := db.CollectPoolStats(database)
	log.Printf("db pool on shutdown: open=%d in_use=%d idle=%d waits=%d wait_duration=%s",
		stats.OpenConnections, stats.InUse, stats.Idle, stats.WaitCount, stats.WaitDuration)

	if err := repo.Close(); err != nil {
		log.Printf("repo close: %v", err)
	}
	if err := database.Close(); err != nil {
		log.Fatalf("db close failed: %v", err)
	}
	log.Println("server stopped cleanly")
}
