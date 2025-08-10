// Map Literal: You can also initialize a map with key-value pairs directly using a map literal.

package main

import "fmt"

func main() {
	myMap := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 15,
	}
	fmt.Println(myMap)
}