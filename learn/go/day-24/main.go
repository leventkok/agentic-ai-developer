package main

import (
	"fmt"
	"net/http"

	"learn/go/day-24/handlers"
)

func main() {
	http.HandleFunc("/menu", handlers.MenuHandler)
	http.HandleFunc("/health", handlers.HealthHandler)

	fmt.Println("Server on :8080")
	http.ListenAndServe(":8080", nil)
}
