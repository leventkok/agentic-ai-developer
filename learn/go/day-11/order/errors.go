package order

import "errors"

var ErrEmptyOrder = errors.New("order is empty")
var ErrUnknownDrink = errors.New("drink not on menu")
