package main

import (
	"fmt"
)

//  basic for loop
func main() {
	for i := 0; i<5; i++ {
		fmt.Println(i)
	}
}

// for loop with only condition

func main() {
	i := 0
	for i <5 {
		fmt.Println(i)
		i++
	}
// }

// for loop with infinite loop

func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i >=4 {
			break
		}
	}
}

// range loop
// iterate over a slice
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

// If only need the vlaue, it can omit the index using the black identifier (_)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}
}