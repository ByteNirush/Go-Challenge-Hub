package main

import (
	"fmt"
)

func sumEvenOdd() {
	var n int
	fmt.Print("Enter a range (n): ")
	fmt.Scan(&n)

	evenSum := 0
	oddSum := 0

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}

	fmt.Println("Sum of even numbers:", evenSum)
	fmt.Println("Sum of odd numbers:", oddSum)
}

func main() {
	sumEvenOdd()
}
