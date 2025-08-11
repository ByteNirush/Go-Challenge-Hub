package main

import "fmt"

type Address struct {
	City   string
}

type Person struct {
	Name   string
	City   string
	Address Address // shadowing the city field 
}

func main() {
	p := Person{
		Name: "Alice",
		City: "New York",
		Address: Address{
			City: "Los Angeles", // This will shadow the Person's City field
		},
	}
	fmt.Println("Person Name:", p.Name)
	fmt.Println("Person City:", p.City)
}

// the Person struct has a City field that shadows the City field from the embedded Address struct. To access the City field of the Address struct, you must explicitly use p.Address.City.

