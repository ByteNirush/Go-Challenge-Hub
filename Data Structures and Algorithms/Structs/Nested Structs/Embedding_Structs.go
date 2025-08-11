package main

import "fmt"

type Address struct {
	Street	string
	City  	string
	Country string
}

type Employee struct {
	ID 		int
	Name 	string
	Address Address //Embedding the Address struct
}

func main() {
	employee := Employee{
		ID: 1,
		Name: "John Doe",
		Address: Address{
			Street:  "123 Elm St",
			City:    "Springfield",
			Country: "USA",
		},
	}

	fmt.Println("Employe Name:", employee.Name)
	fmt.Println("Employe Street:", employee.Address.Street)	
}