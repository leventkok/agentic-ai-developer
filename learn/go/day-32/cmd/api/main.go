package main

import (
	"fmt"
	"log"
	"net/http"

	"learn/go/day-32/internal/handler"
	"learn/go/day-32/internal/store"
)

func main() {
	s := store.NewBookmarkStore()
	h := &handler.BookmarkHandler{Store: s}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /bookmarks", h.ListBookmarks)
	mux.HandleFunc("POST /bookmarks", h.CreateBookmark)
	mux.HandleFunc("GET /bookmarks/{id}", h.GetBookmark)
	mux.HandleFunc("PATCH /bookmarks/{id}", h.UpdateBookmark)
	mux.HandleFunc("DELETE /bookmarks/{id}", h.DeleteBookmark)

	fmt.Println("Bookmarks API on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
