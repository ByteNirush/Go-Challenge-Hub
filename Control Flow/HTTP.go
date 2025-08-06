package main

import (
	"fmt"
)

// 1. Handling HTTP status codes using switch-case

// func main() {
// 	statusCode := 50
// 	switch statusCode {
// 	case 200:
// 		fmt.Println("OK: The request has succeeded.")
// 	case 404:
// 		fmt.Println("Not Found")
// 	case 500:
// 		fmt.Println("Internal Server Error")
// 	default:
// 		fmt.Println("Unknown status code")
// 	}
// }

// 2. checking File Types

func main() {
	fileType := "image/png"
	switch fileType {
	case "image/jpeg", "image/png":
		fmt.Println("This is an image file.")
	case "application/pdf":
		fmt.Println("This is a PDF file.")
	default:
		fmt.Println("Unknown file type.")
	}
}