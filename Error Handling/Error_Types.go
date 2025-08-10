package main

import "fmt"

type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Item with ID %d not found", e.ID)
}

func Find(id int) (string, error) {
	if id < 0 {
		return "", &NotFoundError{ID: id}
	}
	// ...logic to find the item ...
	return "item", nil
}

func main() {
	_, err := Find(-1)
	if err != nil {
		if notFoundErr, ok := err.(*NotFoundError); ok {
			fmt.Printf("Item not found with ID: %d\n", notFoundErr.ID)
		}
	}
}