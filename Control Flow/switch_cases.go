package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	day := "Saturday"
// 	switch day {
// 		case "Monday":
// 			fmt.Println("Start of the work week!")
// 		case "Tuesday":
// 			fmt.Println("It's Tuesday, keep going!")
// 		case "Wednesday":
// 			fmt.Println("Midweek already!")
// 		case "Thursday":
// 			fmt.Println("Almost the weekend!")
// 		case "Friday":
// 			fmt.Println("It's Friday, the weekend is near!")
// 		case "Saturday", "Sunday":
// 			fmt.Println("It's the weekend, relax!")
// 		default:
// 			fmt.Println("Not a valid day of the week.")	
// 	}
// }

// fallthrough 

// func main() {
// 	number := 1
// 	switch number {
// 	case 1:
// 		fmt.Println("Number is 1")
// 		fallthrough
// 	case 2:
// 		fmt.Println("Number is 2")
// 		fallthrough
// 	case 3:
// 		fmt.Println("Number is 3")
// 	default:
// 		fmt.Println("Number is greater than 3 or not recognized")
// 	}
// }


// switch with No Expression

// func main() {
// 	age := 14
// 	switch {
// 	case age < 13:
// 		fmt.Println("You are a child.")
// 	case age >= 13 && age < 20:
// 		fmt.Println("You are a teenager.")
// 	default:
// 		fmt.Println("You are an adult.")
// 	}
// }


// Short Variable Declaration in switch Statements

func main() {
	str := "123"
	switch num, err := strconv.Atoi(str); err {
	case nil:
		fmt.Println("passed number:", num)
	default:
		fmt.Println("Error parsing string:", err)
	}
	
}

