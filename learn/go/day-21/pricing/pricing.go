package pricing

func Total(price float64, qty int) float64 {
	return price * float64(qty)
}

func ApplyDiscount(total float64, percent int) float64 {
	return total * (1 - float64(percent)/100)
}
