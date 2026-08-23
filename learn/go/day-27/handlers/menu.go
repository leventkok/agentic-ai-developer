package handlers

import (
	"encoding/json"
	"net/http"
)

type MenuItem struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var menu = []MenuItem{
	{ID: 1, Name: "Latte", Price: 45},
	{ID: 2, Name: "Espresso", Price: 30},
}

// MenuHandler handles GET /menu
func MenuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}
