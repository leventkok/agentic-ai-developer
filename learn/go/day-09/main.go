package main

import (
	"fmt"
	"learn/go/day-09/staff"
)

func main() {
	fmt.Println(" === RotaCoffee Staff Management === ")

	b := staff.Barista{
		Staff: staff.Staff{Name: "Levent", Email: "levent@rota.coffee"},
		Shift: "Morning",
	}

	fmt.Println(b.String())

	fmt.Println(b.Name)

	m := staff.Manager{
		Staff: staff.Staff{Name: "John", Email: "john@rota.coffee"},
		Level: 2,
	}

	fmt.Println(m.String())

	fmt.Println(m.Staff.String())

}
