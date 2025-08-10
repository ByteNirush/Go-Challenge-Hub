package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func unreliableFunction() error {
	// Simulate an unreliable function that sometimes returns an error
	if rand.Intn(3) == 0 {
		return errors.New("unreliable function failed")
	}
	return nil
}

func retry(attempts int, sleep time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		fmt.Println("Attempt", i+1, "failed:", err)
		time.Sleep(sleep)
	}
	return fmt.Errorf("after %d attempts, the function failed: %w", attempts, err)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	err := retry(3, time.Second, unreliableFunction)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Function succeeded after retries")
	}
}
