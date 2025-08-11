package main

import "fmt"

func main() {
	// Defining and initializing an anonymous struct
	
	person := struct {
		Name string
		Age  int
	}{
		Name: "Nirush",
		Age:  20,
	}

	fmt.Println(person)
}