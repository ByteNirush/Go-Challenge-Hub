package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var person Person
	person.FirstName = "Nirush"
	person.LastName = "Man"
	person.Age = 20

	fmt.Println(person)
}