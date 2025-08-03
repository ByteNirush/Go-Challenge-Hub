package main

import (
	"fmt"
)

// if statement
// func main() {
// 	age := 20
// 	if age >= 18 {
// 		fmt.Println("You are adult.")
// 	}
// }


// if-else statement

// func main() {
// 	age := 20
// 	if age >=18 {
// 		fmt.Println("You are adult.")
// 	} else {
// 		fmt.Println("You are not adult.")
// 	}
// }


// if-else if-else statement

func main(){
	age := 35
	if age >= 30 {
		fmt.Println("You are middle-aged.")
	} else if age >= 18 {
		fmt.Println("You are adult.")
	} else if age >= 13 {
		fmt.Println("You are teenager.")
	} else {
		fmt.Println("You are child.")
	}
}