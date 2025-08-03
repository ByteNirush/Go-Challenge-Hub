package main

import "fmt"

// integer type
func main() {
	var age int =30
	fmt.Println(age)
}

// floating point type
func main() {
	var pi float64 = 3.14
	var e float32 = 2.718

	fmt.Println(pi, e)
}


// complex number

func main() {
	var a complex64 = complex(1, 2) // 1 is the real part, 2 is the imaginary part
	var b complex128 = complex(3, 4) // 3 is the real part, 4 is the imaginary part

	fmt.Println(a,b)
}


// boolean type

func main() {
	var isTrue bool = true
	var isFalse bool = false

	fmt.Println(isTrue, isFalse)
}


// string type
func main () {
	var name string = "Man"
	var message string = "Hello!"
	fmt.Println(name, message)
}