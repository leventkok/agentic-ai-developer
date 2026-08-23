package main

import (
	"fmt"
	"log"
	"net/http"

	"learn/go/day-27/handlers"
)

func main() {
	mux := http.NewServeMux()

	// TODO Task 1: register routes
	// mux.HandleFunc("GET /menu", handlers.MenuHandler)
	// mux.HandleFunc("GET /orders/{id}", handlers.OrderGetHandler)
	// mux.HandleFunc("POST /orders", handlers.OrderCreateHandler)

	fmt.Println("Server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
