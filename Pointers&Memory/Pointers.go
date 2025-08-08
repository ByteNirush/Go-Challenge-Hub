// package main

// import (
// 	"fmt"
// 	"time"
// )

// Declaring pointers
// func main() {
// 	var x int = 10
// 	var p *int // Declare a pointer to an integer

//     p = &x // Assign the address of x to the pointer p

//     fmt.Println("Value of x:", x)   // Output: Value of x: 10
//     fmt.Println("Address of x:", &x) // Output: Address of x: 0x... (some memory address)
//     fmt.Println("Value of p:", p)   // Output: Value of p: 0x... (same memory address as &x)
//     fmt.Println("Value pointed to by p:", *p) // Output: Value pointed to by p: 10
// }


// Dereferencing Pointers
// func main() {
// 	x := 20
// 	p := &x

// 	fmt.Println("Original value of x:", x)

// 	*p = 70 // Dereferencing p and change the value at that address

// 	fmt.Println("New value of x:", x) 
// 	fmt.Println("Value pointed to by p:", *p)
// }


// Zero Value of a Pointer

// func main() {
// 	var p *int // Declare a pointer to an integer without initializing it

// 	fmt.Println("Value of p:", p)

	// The following line would cause a panic:
    // fmt.Println(*p) // Dereferencing a nil pointer
// }


// Pointers and Function Arguments
// func increment(x *int) {
// 	*x++
// }

// func main() {
// 	y := 5
// 	fmt.Println("Original value of y:", y)

// 	increment(&y) // Pass the address of y to the function

// 	fmt.Println("New value of y:", y)
// }


// Pointers to Structs
// type Person struct {
// 	Name string
// 	Age  int
// }

// func updateAge(p *Person, newAge int) {
// 	p.Age = newAge // Accessing the Age field of the Person struct
// }

// func main() {
// 	Person := Person{Name: "Nirush", Age: 20}
// 	fmt.Println("Original Person:", Person)

// 	updateAge(&Person, 25) // Pass a pointer to the person struct

// 	fmt.Println("Updated Person:", Person)
// }



// Memory Management

// The new Function
// func main() {
//     p := new(int) // Allocate memory for an int on the heap and return a pointer to it

//     fmt.Println("Value of p:", p)
//     fmt.Println("Value pointed to by p:", *p) // Output: Value pointed to by p: 0 (zero value of int)

//     *p = 42 // Assign a value to the memory location pointed to by p

//     fmt.Println("New value pointed to by p:", *p)
// }


// Memory Leaks
// var globalSlice []int

// func main() {
// 	for i := 0; i < 1000; i++ {
// 		// Simulate adding data to a slice
// 		data := make([]int, 10) // Allocate a small slice
// 		globalSlice = append(globalSlice, data[0]) // Keep a reference to the first element
// 		time.Sleep(10 * time.Millisecond)
// 		fmt.Println("Iteration:", i)
// 	}
// }
