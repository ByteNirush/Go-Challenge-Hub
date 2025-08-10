// Declaring a nil map: You can declare a map without initializing it. This creates a nil map.

package main

import "fmt"

func main() {
	var myMap map[string]int
	
	if myMap == nil {
		fmt.Println("myMap is nil")
		myMap = make(map[string]int)
	}

	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 15

	fmt.Println(myMap)
}