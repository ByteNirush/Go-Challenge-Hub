// The dot operator allows you to both read and modify the values of struct fields.

package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages   int
}

func main() {
	book := Book {
		Title: "This Go Programming Language",
		Author: "John Doe",
		Pages: 203,
	}

	fmt.Println("Title:", book.Title)
	fmt.Println("Author:", book.Author)
	fmt.Println("Pages:", book.Pages)

	// Modifying a field
	book.Pages = 300
	fmt.Println("Updated Book:", book.Pages)
}