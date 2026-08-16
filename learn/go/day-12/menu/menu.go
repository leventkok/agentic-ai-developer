package menu

import "fmt"

var Prices = map[string]int{
	"latte":    45,
	"espresso": 35,
	"filter":   30,
}

func AddItem(name string, price int) {
	Prices[name] = price
}

func GetPrice(name string) (int, bool) {
	price, ok := Prices[name]
	return price, ok
}

func PrintMenu() {
	for name, price := range Prices {
		fmt.Printf("%s: %d TL\n", name, price)
	}
}
