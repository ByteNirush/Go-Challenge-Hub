// Filtering Even Numbers

package main

import (
	"fmt"
)
func filterEvenNumbers(numbers []int) []int {
	var evenNumbers []int
	for _, number := range numbers {
		if number%2 == 0 {
			evenNumbers = append(evenNumbers, number)
		}
	}
	return evenNumbers
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := filterEvenNumbers(numbers)
	fmt.Println("Even numbers:", evenNumbers) // Output: Even numbers: [2 4 6 8 10]
}
