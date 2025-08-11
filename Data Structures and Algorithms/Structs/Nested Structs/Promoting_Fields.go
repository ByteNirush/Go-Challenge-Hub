package main

import "fmt"

type Address struct {
	Street 		string
	City   		string
	Country  	string
}

type Employee struct {
	ID 		int
	Name 	string
	Address Address // Embedding the Address struct(anonymous field)
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
	fmt.Println("Employee Name:", employee.Name)
}


// In this, Address struct is embedded as an anonymous field in the Employee struct. This promotes the fields of the Address struct to the Employee struct's level, allowing you to access them directly using employee.City instead of employee.Address.City.

