package menu

import "fmt"

type Drink struct {
	Name  string
	Price int
}

func (d Drink) String() string {
	return fmt.Sprintf("%s: %d TL", d.Name, d.Price)
}

func (d Drink) IsPremium() bool {
	if d.Price >= 40 {
		return true
	}
}

func (d *Drink) ApplyDiscount(percent int) {
	d.Price = d.Price * (100 - percent) / 100
}

func (d *Drink) SetPrice(price int) {
	d.Price = price
}
