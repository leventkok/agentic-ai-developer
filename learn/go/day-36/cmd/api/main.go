package main

import (
	"fmt"
	"log"
	"net/http"

	"learn/go/day-36/internal/handler"
	"learn/go/day-36/internal/middleware"
	"learn/go/day-36/internal/store"
)

func main() {
	s := store.NewMemoryStore()
	h := &handler.BookmarkHandler{Store: s}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /bookmarks", h.ListBookmarks)
	mux.HandleFunc("POST /bookmarks", h.CreateBookmark)
	mux.HandleFunc("GET /bookmarks/{id}", h.GetBookmark)
	mux.HandleFunc("PATCH /bookmarks/{id}", h.UpdateBookmark)
	mux.HandleFunc("DELETE /bookmarks/{id}", h.DeleteBookmark)

	root := middleware.Chain(mux, middleware.Logging, middleware.Recovery)

	fmt.Println("Bookmarks API MVP on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", root))
}
