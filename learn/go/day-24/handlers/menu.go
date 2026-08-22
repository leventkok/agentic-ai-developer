package handlers

import (
	"encoding/json"
	"net/http"
)

type MenuItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var menu = []MenuItem{
	{Name: "Latte", Price: 45},
	{Name: "Espresso", Price: 30},
}

// MenuHandler returns the cafe menu as JSON.
func MenuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}

// HealthHandler returns 200 OK for health checks.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
