package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	// Create a Point struct using a struct literal
	p := Point{x: 10, y: 20}
	fmt.Println(p)

	// Create a point struct without specifying field names (other matters)
	p2 := Point{30, 40}
	fmt.Println(p2)

	// Partial initialization
	p3 := Point{x: 50}
	fmt.Println(p3)
}