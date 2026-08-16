package menu

type Order struct {
	Items []string
}

func NewOrder() Order {
	return Order{Items: []string{}}
}

func (o *Order) Add(item string) {
	o.Items = append(o.Items, item)
}

func (o Order) Total() int {
	sum := 0
	for _, item := range o.Items {
		if price, ok := Prices[item]; ok {
			sum += price
		}
	}
	return sum
}
