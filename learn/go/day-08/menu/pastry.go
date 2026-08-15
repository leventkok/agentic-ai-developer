package menu

import "fmt"

type Pastry struct {
	Name  string
	Price int
}

func (p Pastry) String() string {
	return fmt.Sprintf("%s: %d TL", p.Name, p.Price)
}

func (p Pastry) PriceTL() int {
	return p.Price
}
