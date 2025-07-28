package main

import(
	"fmt"
)

// Define the Book struct
type Book struct {
	Title string
	Author string
	Pages int
}

// Define the PrintBookDetails method for the Book struct
func (b Book) PrintBookDetails(){
	fmt.Printf("Title: %s\n Author: %s\n Pages: %d\n", b.Title, b.Author, b.Pages)
}

func main(){
	// Create a slice of Book structs
	Books := []Book{
		{
			Title: "Go Programming",
			Author: "GO Author",
			Pages: 140,
		},

		{
			Title: "Python Programming",
			Author: "Python Author",
			Pages: 333,
		},

		{
			Title: "Java Programming",
			Author: "Java Author",
			Pages: 203,
		},
	}
	// Call the PrintBookDetails method
	for _, Book := range Books {
		Book.PrintBookDetails()
	}	
}
