package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateOrderRequest struct {
	Item  string  `json:"item"`
	Total float64 `json:"total"`
}

type Order struct {
	ID    int     `json:"id"`
	Item  string  `json:"item"`
	Total float64 `json:"total"`
}

var nextOrderID = 1

// CreateOrderHandler handles POST /orders
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	requestID := r.Header.Get("X-Request-ID")
	if requestID == "" {
		requestID = "unknown"
	}
	fmt.Printf("requestID=%s POST /orders\n", requestID)

	w.Header().Set("X-Request-ID", requestID)
	w.Header().Set("Cache-Control", "no-store")

	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Item == "" || req.Total <= 0 {
		writeError(w, http.StatusBadRequest, "item and total are required")
		return
	}

	order := Order{ID: nextOrderID, Item: req.Item, Total: req.Total}
	nextOrderID++

	writeJSON(w, http.StatusCreated, order)
}
