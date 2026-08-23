package main

import (
	"fmt"
	"log"
	"net/http"

	"learn/go/day-30/handlers"
	"learn/go/day-30/middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /notes", handlers.ListNotesHandler)
	mux.HandleFunc("POST /notes", handlers.CreateNoteHandler)
	mux.HandleFunc("DELETE /notes/{id}", handlers.DeleteNoteHandler)

	root := middleware.Chain(mux, middleware.Logging, middleware.Recovery)

	fmt.Println("Notes API on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", root))
}
