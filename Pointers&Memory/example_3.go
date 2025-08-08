// Using Pointers with slices


package main

import "fmt"

func modifySlice(s []int) {
	s[0] = 150
	s[2] = 300
}

func main() {
	mySlice := []int{1,2,3}
	fmt.Println("Original Slice:", mySlice)

	modifySlice(mySlice)
	fmt.Println("Modified Slice:", mySlice)
}