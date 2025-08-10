// Calculating the Average of Numbers

package main

import (
	"fmt"
)

func calculateAverage (numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0 // Avoid division by zero
	}
	Sum := 0.0
	for _, number := range numbers {
		Sum += number
	}
	return Sum / float64(len(numbers))
}

func main() {
	numbers := []float64{10.5, 20.7, 30.5, 40.3}
	average := calculateAverage (numbers)
	fmt.Printf("Average: %.2f\n", average)
}