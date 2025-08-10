// Checking for Key Existence: To determine if a key exists in a map, use the "comma ok" idiom.

package main

import "fmt"

func main() {
    myMap := map[string]int{
        "apple":  1,
        "banana": 2,
        "cherry": 3,
    }

    // Check if the key "banana" exists
    value, ok := myMap["banana"]
    if ok {
        fmt.Println("Key exists, value:", value) // Output: Key exists, value: 2
    } else {
        fmt.Println("Key does not exist")
    }

    // Check if the key "grape" exists
    value, ok = myMap["grape"]
    if ok {
        fmt.Println("Key exists, value:", value)
    } else {
        fmt.Println("Key does not exist") // Output: Key does not exist
        fmt.Println("Value:", value)       // Output: Value: 0 (zero value of int)
    }
}
