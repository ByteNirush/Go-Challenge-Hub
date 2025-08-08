// Modifying an Array Element Using Pointers

package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	p := &arr[0]

	*p = 100

	fmt.Println("Modified array:", arr)
}