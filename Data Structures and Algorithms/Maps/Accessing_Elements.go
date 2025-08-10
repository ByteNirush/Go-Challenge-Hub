// Accessing Elements: Use the key to access the corresponding value.

package main

import "fmt"

func main() {
    myMap := map[string]int{
        "apple":  1,
        "banana": 2,
        "cherry": 3,
    }

    // Access the value associated with the key "banana"
    value := myMap["banana"]
    fmt.Println(value) // Output: 2

    // Accessing a non-existent key returns the zero value of the value type
    nonExistentValue := myMap["grape"]
    fmt.Println(nonExistentValue) // Output: 0
}
