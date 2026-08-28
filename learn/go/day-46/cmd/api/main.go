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

	"learn/go/day-46/internal/config"
	"learn/go/day-46/internal/db"
	"learn/go/day-46/internal/handler"
	"learn/go/day-46/internal/middleware"
	"learn/go/day-46/internal/repository/sqlite"
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

	mux := http.NewServeMux()
	mux.HandleFunc("GET /bookmarks", h.ListBookmarks)
	mux.HandleFunc("POST /bookmarks", h.CreateBookmark)
	mux.HandleFunc("POST /bookmarks/bulk", h.BulkCreateBookmarks)
	mux.HandleFunc("GET /bookmarks/{id}", h.GetBookmark)
	mux.HandleFunc("PATCH /bookmarks/{id}", h.UpdateBookmark)
	mux.HandleFunc("DELETE /bookmarks/{id}", h.DeleteBookmark)

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
