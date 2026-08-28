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

	"learn/go/day-42/internal/config"
	"learn/go/day-42/internal/db"
	"learn/go/day-42/internal/handler"
	"learn/go/day-42/internal/middleware"
	"learn/go/day-42/internal/store"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.Open("bookmarks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err := db.RunMigrations(database, "migrations"); err != nil {
		log.Fatal(err)
	}

	bookmarks, err := db.ListBookmarks(context.Background(), database)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB seed count:", len(bookmarks))

	s := store.NewMemoryStore()
	h := &handler.BookmarkHandler{
		Store:       s,
		ListTimeout: time.Duration(cfg.ListTimeoutMS) * time.Millisecond,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /bookmarks", h.ListBookmarks)
	mux.HandleFunc("POST /bookmarks", h.CreateBookmark)
	mux.HandleFunc("GET /bookmarks/{id}", h.GetBookmark)
	mux.HandleFunc("PATCH /bookmarks/{id}", h.UpdateBookmark)
	mux.HandleFunc("DELETE /bookmarks/{id}", h.DeleteBookmark)

	root := middleware.DefaultStack(cfg, s, mux)
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
	log.Println("server stopped cleanly")
}