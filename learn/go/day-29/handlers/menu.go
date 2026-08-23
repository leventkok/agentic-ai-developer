package handlers

import (
	"encoding/json"
	"net/http"
)

func MenuHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]map[string]any{
		{"name": "Latte", "price": 45},
	})
}

// PanicHandler demonstrates recovery middleware — GET /panic
func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("something went wrong")
}
