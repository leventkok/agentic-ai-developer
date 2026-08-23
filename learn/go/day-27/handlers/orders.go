package handlers
import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Order struct {
	ID    int     `json:"id"`
	Item  string  `json:"item"`
	Total float64 `json:"total"`
}

var orders = map[string]Order{
	"1": {ID: 1, Item: "Latte", Total: 45},
	"2": {ID: 2, Item: "Espresso", Total: 30},
}

// OrderGetHandler handles GET /orders/{id}
func OrderGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	order, ok := orders[id]
	if !ok {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// OrderCreateHandler handles POST /orders
func OrderCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Item  string  `json:"item"`
		Total float64 `json:"total"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	id := len(orders) + 1
	order := Order{ID: id, Item: req.Item, Total: req.Total}
	orders[strconv.Itoa(id)] = order

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
