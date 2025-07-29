// Function scope: Variables declared inside a function are only accessible within that function.

package main

import "fmt"

func myfunction() {
	age := 22
	fmt.Println(age)
}

func main() {
	myfunction()
	// fmt.Println(age) // This line would cause an error because 'age' is not defined in this scope.
	fmt.Println()
}


// Block scope: Variables declared inside a block (like an if statement ir a for loop) are only accessible within that block.

package main

import "fmt"

func main() {
	if true {
		name := "Man"
		fmt.Println(name)
	}
	// fmt.Println(name) // This would cause a compile error: undefined name 'name'

	for i := 0; i < 5; i++ {
		fmt.Println(i) // i is accessible here
	}
	// fmt.Println(i) // This would cause a compile error: undefined i
}

// The variable name is declared inside the if block and i inside the for loop; they are only accessible within their respective blocks.


