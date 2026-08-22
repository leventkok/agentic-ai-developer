package cafe

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

// MenuHandler serves GET /menu as JSON.
func MenuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}

// OrderHandler accepts POST /order with {"name","qty"} and returns line total.
func OrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Name string `json:"name"`
		Qty  int    `json:"qty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	price := findPrice(req.Name)
	total, err := LineTotal(price, req.Qty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"total": total})
}

func findPrice(name string) float64 {
	for _, item := range menu {
		if item.Name == name {
			return item.Price
		}
	}
	return 0
}
