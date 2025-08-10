package main

import "errors"

var ErrNotFound = errors.New("Not Found")

func Find(id int) (string, error) {
	if id < 0 {
		return "", ErrNotFound
	}

	// ...logic to find the item ...
	return "item", nil
}

func main() {
	_, err := Find(-1)
	if err == ErrNotFound {
		println("Item not found")
	}
}


