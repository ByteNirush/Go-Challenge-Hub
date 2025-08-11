package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// Method to calculate the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method to calculate the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	rectangele := Rectangle{
		Width:  5.0,
		Height: 3.0,
	}

	fmt.Println("Area:", rectangele.Area())
	fmt.Println("Perimeter:", rectangele.Perimeter())

}