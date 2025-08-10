package main

import (
	"errors"
	"fmt"
)

func innerFunction() error {
	return errors.New("Original error")
}

func outerFunction() error {
	err := innerFunction()
	if err != nil {
		return fmt.Errorf("Outer Function Failed: %w", err)
	}
	return nil
}

func main() {
	err := outerFunction()
	if err != nil {
		fmt.Println(err)

		unwrappedErr := errors.Unwrap(err)
		if unwrappedErr != nil {
			fmt.Println("Unwrapped Error:", unwrappedErr)
		}

		if errors.Is(err, errors.New("Original error")) {
			fmt.Println("The error chain contains the original error")
		}
	}
}