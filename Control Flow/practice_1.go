// Write a program that takes a month number (1-12) as input and prints the corresponding month name using a switch statement.

// package main

// import "fmt"

// func main() {
// 	var month int
// 	fmt.Print("Enter month number (1-12): ")
// 	fmt.Scan(&month)
// 	switch month {
// 	case 1:
// 		fmt.Println("January")
// 	case 2:
// 		fmt.Println("February")
// 	case 3:
// 		fmt.Println("March")
// 	case 4:
// 		fmt.Println("April")
// 	case 5:
// 		fmt.Println("May")
// 	case 6:
// 		fmt.Println("June")
// 	case 7:
// 		fmt.Println("July")
// 	case 8:
// 		fmt.Println("August")
// 	case 9:
// 		fmt.Println("September")
// 	case 10:
// 		fmt.Println("October")
// 	case 11:
// 		fmt.Println("November")
// 	case 12:
// 		fmt.Println("December")
// 	default:
// 		fmt.Println("Invalid month number")
// 	}
// }

// Write a program that takes a character as input and checks if it is a vowel or a consonant using a switch statement.

// package main

// import "fmt"

// func main() {
// 	var character string
// 	fmt.Print("Enter a character: ")
// 	fmt.Scan(&character)
// 	switch character {
// 		case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
// 			fmt.Println(character, "is a vowel")
// 		default:
// 			fmt.Println(character, "is a consonant")

// 	}
// }

// Write a program that simulates a simple calculator. It should take two numbers and an operator (+, -, *, /) as input and perform the corresponding operation using a switch statement. Handle division by zero errors.

package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)
	switch operator {
		case "+":
			fmt.Println("Result:", num1+num2)
		case "-":
			fmt.Println("Result:", num1-num2)
		case "*":
			fmt.Println("Result:", num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero is not allowed")
			} else {
				fmt.Println("Result:", num1/num2)
			}
		default:
			fmt.Println("Error: Invalid operator")
	}
}