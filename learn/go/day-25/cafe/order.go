package cafe

import "fmt"

// LineTotal returns price * qty; errors if qty <= 0.
func LineTotal(price float64, qty int) (float64, error) {
	if qty <= 0 {
		return 0, fmt.Errorf("invalid qty: %d", qty)
	}
	return price * float64(qty), nil
}

// ApplyDiscount returns total after percent off (0–100).
func ApplyDiscount(total float64, percent int) (float64, error) {
	if percent < 0 || percent > 100 {
		return 0, fmt.Errorf("invalid percent: %d", percent)
	}
	return total * (1 - float64(percent)/100), nil
}
