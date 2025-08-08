// Swapping Values Using Pointers

package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 10
	y := 20

	fmt.Println("Before swap:", x, "y =", y)
	swap(&x, &y)

	fmt.Println("After swap: x =", x, "y =", y)
}