package pricing

import "fmt"

func Total(price float64, qty int) float64 {
	return price * float64(qty)
}

func ApplyDiscount(total float64, percent int) float64 {
	return total * (1 - float64(percent)/100)
}

func DiscountedTotal(price float64, qty int, percent int) (float64, error) {
	if qty <= 0 {
		return 0, fmt.Errorf("invalid qty: %d", qty)
	}
	if percent < 0 || percent > 100 {
		return 0, fmt.Errorf("invalid percent: %d", percent)
	}
	return ApplyDiscount(Total(price, qty), percent), nil
}
