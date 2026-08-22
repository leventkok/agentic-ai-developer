package main

import (
	"fmt"
	"net/http"

	"learn/go/day-25/cafe"
)

func main() {
	http.HandleFunc("/menu", cafe.MenuHandler)
	http.HandleFunc("/order", cafe.OrderHandler)
	fmt.Println("Cafe server on :8080")
	http.ListenAndServe(":8080", nil)
}
