package main

import (
	"fmt"
)

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		}
		if i%3 == 0 && i%5 != 0{
			fmt.Println("Fizz")
		}
		if i%3 != 0 && i%5 == 0{
			fmt.Println("Buzz")
		} else if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		}

	}
}

func main() {
	FizzBuzz(30) //  prints FizzBuzz for numbers 1 to 30
}
