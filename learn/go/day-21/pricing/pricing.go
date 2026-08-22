package pricing

// Total returns the line total for a menu item.
func Total(price float64, qty int) float64 {
	return price * float64(qty)
}

// ApplyDiscount returns total after percent discount (0–100).
func ApplyDiscount(total float64, percent int) float64 {
	return total * (1 - float64(percent)/100)
}
