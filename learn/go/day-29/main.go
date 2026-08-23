package main

import (
	"fmt"
	"log"
	"net/http"

	"learn/go/day-29/handlers"
	"learn/go/day-29/middleware"
)

func main() {
	mux := http.NewServeMux()

	menu := middleware.Chain(
		http.HandlerFunc(handlers.MenuHandler),
		middleware.AuthPlaceholder,
	)
	mux.Handle("GET /menu", menu)
	mux.HandleFunc("GET /panic", handlers.PanicHandler)

	root := middleware.Chain(mux, middleware.Logging, middleware.Recovery)

	fmt.Println("Server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", root))
}
