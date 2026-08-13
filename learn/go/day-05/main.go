package main

import (
	"fmt"
	"learn/go/day-05/calc"
	"learn/go/day-05/convert"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . <num1> <operator> <num2>")
		fmt.Println("Example: go run . 1 + 2")
		return
	}

	a, err1 := strconv.Atoi(os.Args[1])
	op := os.Args[2]
	b, err2 := strconv.Atoi(os.Args[3])

	if err1 != nil || err2 != nil {
		fmt.Println("Error: both arguments must be integers")
		return
	}

	result, err := calc.Calculate(a, b, op)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%d %s %d = %d\n", a, op, b, result)

	demoTemperature()
}

func demoTemperature() {
	fmt.Println("===== Temperature Conversion =====")
	c := 100.0
	f := convert.ToFahrenheit(c)
	fmt.Printf("%f°C = %f°F\n", c, f)

	f2 := 32.0
	c2 := convert.ToCelsius(f2)
	fmt.Printf("%f°F = %f°C\n", f2, c2)
}
