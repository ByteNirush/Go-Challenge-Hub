package main

import (
	"fmt"
	"errors"
)

// shorthand parameter declaration
// func add(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	result := add(3,5)
// 	fmt.Println(result)
// }

// Positional parameters

// func greet(name string, age int) {
// 	fmt.Printf("Hello, %s %d!\n", name, age)
// }

// func main() {
// 	greet("Alice", 30)

// }

// Variadic Parameters

// func sum (numbers ...int) int {
// 	total := 0
// 	for _, n := range numbers {
// 		total += n
// 	}
// 	return total
// }

// func main() {
// 	result := sum(1, 2, 3, 4, 5)
// 	fmt.Println(result)

// 	result2 := sum(10, 20, 30)
// 	fmt.Println(result2)

// }


// slice to variadic parameters
// func sum(numbers ...int) int {
// 	total := 0
// 	for _, n := range numbers {
// 		total += n
// 	}
// 	return total
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5}
// 	result := sum(nums...)
// 	fmt.Println(result)
// }

// Named Return Values

// func calculate(x, y int) (sum int, difference int) {
// 	sum = x + y
// 	difference = x - y
// 	return
// }

// func main() {
// 	sum, diff := calculate(10, 5)
// 	fmt.Println("Sum:", sum)
// 	fmt.Println("Difference:", diff)
// }



// Single Return Value

// func square(x int) int {
// 	return x * x
// }

// func main() {
// 	result := square(4)
// 	fmt.Println("Square:", result)
// }


// Multiple Return Values

func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }
    return x / y, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result) 

    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
