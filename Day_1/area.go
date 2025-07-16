package main

import(
	"fmt"
)

// Define the Shape interface
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Define the main function
func main() {

	// Create slices of Circle and Rectangle structs
	shapes := []Shape{
		Circle{Radius: 5},
		Circle{Radius: 10},
		Circle{Radius: 15},
		Rectangle{Width: 4, Height: 6},
		Rectangle{Width: 5, Height: 10},
		Rectangle{Width: 6, Height: 12},
	}


	// Iterate over the slices and call the Area method
	for _, shape := range shapes {
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("Circle with Radius %.2f has Area: %.2f\n", s.Radius, s.Area())
		case Rectangle:
			fmt.Printf("Rectangle with Width %.2f and Height %.2f has Area: %.2f\n", s.Width, s.Height, s.Area())
		}
	}

}
