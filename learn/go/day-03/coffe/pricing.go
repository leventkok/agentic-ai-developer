package coffe

// CalculateTotal returns the sum of all prices
func CalculateTotal(prices []int) int {
	total := 0
	for _, p := range prices {
		total += p
	}
	return total
}

func GetDrinkPrice(drink string) (int, bool) {
	switch drink {
	case "Latte":
		return 45, true
	case "Espresso":
		return 35, true
	default:
		return 0, false
	}
}
