package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"learn/go/day-37/internal/config"
	"learn/go/day-37/internal/handler"
	"learn/go/day-37/internal/middleware"
	"learn/go/day-37/internal/store"
)

func main() {

	cfg, err := config.Load()
	if err != nil{
		log.Fatal(err)
	}

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

	root := middleware.Chain(mux, middleware.Logging, middleware.Recovery)


	addr := ":" + cfg.Port

	fmt.Printf("Bookmarks API [%s] on http://localhost%s\n", cfg.Env, addr)
	log.Fatal(http.ListenAndServe(addr, root))
}
