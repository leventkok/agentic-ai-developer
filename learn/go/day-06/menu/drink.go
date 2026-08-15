package menu

import "fmt"

type Drink struct {
	Name  string
	Price int
}

func (d Drink) String() string {
	return fmt.Sprintf("%s: %d TL", d.Name, d.Price)
}
