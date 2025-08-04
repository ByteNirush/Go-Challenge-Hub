// Write a program that takes a student's score as input and prints their grade based on the following criteria:
// 90-100: A
// 80-89: B
// 70-79: C
// 60-69: D
// Below 60:

package main

import (
	"fmt"
)

func main() {
	score := 85
	if score >= 90 && score <= 100 {
		fmt.Println("Grade: A")
	} else if score >= 80 && score < 90 {
		fmt.Println("Grade: B")
	} else if score >= 70 && score < 80 {
		fmt.Println("Grade: C")
	} else if score >= 60 && score < 70 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
}

// Calculating the Sum of Numbers

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum:", sum)
}