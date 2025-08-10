// Using make: This is the most common way to create a map.

package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	// add key value pairs
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3

	// print the map
	fmt.Println(myMap)
}