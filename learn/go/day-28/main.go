package main

import (
	"fmt"
	"log"
	"net/http"

	"learn/go/day-28/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /orders", handlers.CreateOrderHandler)

	fmt.Println("Server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
