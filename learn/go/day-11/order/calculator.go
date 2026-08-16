package order

import "fmt"

var prices = map[string]int{
	"latte":    45,
	"espresso": 35,
	"filter":   30,
}

func Price(drink string) (int, error) {
	if drink == "" {
		return 0, fmt.Errorf("lookup price: %w", ErrUnknownDrink)
	}
	price, ok := prices[drink]
	if !ok {
		return 0, fmt.Errorf("lookup price for %q :%w", drink, ErrUnknownDrink)
	}

	return price, nil
}

func Total(items []string) (int, error) {
	if len(items) == 0 {
		return 0, ErrEmptyOrder
	}

	sum := 0

	for _, item := range items {
		price, err := Price(item)
		if err != nil {
			return 0, err
		}
		sum += price
	}
	return sum, nil

}
