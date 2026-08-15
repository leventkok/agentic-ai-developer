package main

import (
	"fmt"
	"learn/go/day-10/logger"
	"learn/go/day-10/shapes"
)

func printShapeInfo(s shapes.Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func testShapes() {
	tests := []struct {
		name      string
		shape     shapes.Shape
		wantArea  float64
		wantPerim float64
	}{
		{"Circle r=5", shapes.Circle{Radius: 5}, 78.54, 31.42},
		{"Rectangle 4x6", shapes.Rectangle{Width: 4, Height: 6}, 24.0, 20.0},
	}

	for _, tt := range tests {
		area := tt.shape.Area()
		perim := tt.shape.Perimeter()
		areaOK := int(area*100) == int(tt.wantArea*100) // rough compare
		perimOK := int(perim*100) == int(tt.wantPerim*100)
		if areaOK && perimOK {
			fmt.Printf("PASS | %s\n", tt.name)
		} else {
			fmt.Printf("FAIL | %s (got area=%.2f perim=%.2f)\n", tt.name, area, perim)
		}
	}
}

func demoLogger(log logger.Logger) {
	log.Log("RotaCoffee opened")
	log.Log("Order recevied: 2 lattes")
}

func main() {
	fmt.Println(" === Shape Tests === ")

	testShapes()

	printShapeInfo(shapes.Circle{Radius: 5})

	fmt.Println(" === Logger Demo === ")

	demoLogger(logger.ConsoleLogger{})
	demoLogger(logger.NoopLogger{})

}
