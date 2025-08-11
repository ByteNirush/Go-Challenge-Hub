package main

import "fmt"

type Circle struct {
	Radius float64
}

// Method with a value receiver
func (c Circle) DoubleRadius() {
	c.Radius *= 2 			// Modifies the copy of the Circle struct
	fmt.Println("Radius inside DoubleRadius:", c.Radius)
}

func main(){
	circle := Circle {
		Radius: 5.0,
	}
	fmt.Println("Initial Radius:", circle.Radius)
	circle.DoubleRadius()		// Calls the DoubleRadius method
	fmt.Println("Radius after DoubleRadius:", circle.Radius)
}