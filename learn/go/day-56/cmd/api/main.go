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

	"learn/go/day-56/internal/auth"
	"learn/go/day-56/internal/config"
	"learn/go/day-56/internal/db"
	"learn/go/day-56/internal/httpapi"
	"learn/go/day-56/internal/repository/sqlite"
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

	bookmarks, err := sqlite.New(database)
	if err != nil {
		database.Close()
		log.Fatal(err)
	}

	tokens, err := auth.NewTokenService(cfg.JWTSecret, cfg.JWTTTL())
	if err != nil {
		bookmarks.Close()
		database.Close()
		log.Fatal(err)
	}
	authStore := sqlite.NewAuthStore(database, tokens)

	handler := httpapi.NewRouter(httpapi.Deps{
		Cfg:       cfg,
		Bookmarks: bookmarks,
		Auth:      authStore,
	})

	addr := ":" + cfg.Port
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeoutSec)*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	bookmarks.Close()
	database.Close()
	log.Println("server stopped cleanly")
}
