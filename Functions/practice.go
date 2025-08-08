package main

import (
	"fmt"
)

// Calculate Area: Write a function that takes the length and width of a rectangle as input and returns its area.
// func Area(length, width float64) float64 {
// 	return length * width
// }

// func main() {
// 	length := 5.0
// 	width := 3.0
// 	area := Area(length, width)
// 	fmt.Printf("Area of rectangle: %.2f\n", area)
// }

// Find Max: Write a function that takes a variable number of integers as input and returns the largest integer.

func Max(number ...int) int {
	if len(number) == 0 {
		return 0
	}
	max :=number[0]
	for _, n := range number {
		if n > max {
			max = n
		}
	}
	return max
}

func main() {
	numbers := []int{3, 15, 7, 2, 8}
	Max(numbers...)
	fmt.Println("The maximum number is:", Max(numbers...))
}
