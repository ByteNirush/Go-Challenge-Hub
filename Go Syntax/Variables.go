// 1. Using the 'var' keyword: this is the mst explicit way to declare a variable.

package main

import "fmt"

func main() {
	var age int
	var name string
	var isStudent bool
	
	fmt.Println(age, name, isStudent)
}
// If don't initialize the variable, Go assigns them zero values. int is 0, string is "", and bool is false.


// 2. Using the 'var' keyword with initialization: It can declare and initialize a variable in the same line.

package main
import "fmt"

func main() {
	var age int = 33
	var name string = "Man"
	var isStudent bool = false

	fmt.Println(age, name, isStudent)
}

// 3. Type inference with 'var': Go can infer the type of the variable based  on the initial value.


// package main

import "fmt"

func main(){
	var age = 33  // age is the inferred to be an int 
	var name = "Man" //name is inferred to be a string
	var isStudent = true // isStudent is inferred to be a bool
	 fmt.Println(age, name, isStudent)
}

// 4. Short variable declaration: this is the most concise way to declare and initialize variables, using the ':=' operator. It only be used inside a function.

package main

import "fmt"

func main(){
	age := 33
	name := "Man"
	isStudent := true
	fmt.Println(age, name, isStudent)
}