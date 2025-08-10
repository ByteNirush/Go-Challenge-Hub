// Deleting Elements: Remove key-value pairs from the map using the delete function.


package main

import "fmt"

func main() {
    myMap := map[string]int{
        "apple":  1,
        "banana": 2,
        "cherry": 3,
    }

    // Delete the key "banana"
    delete(myMap, "banana")
    fmt.Println(myMap) // Output: map[apple:1 cherry:3]

    // Deleting a non-existent key has no effect
    delete(myMap, "grape")
    fmt.Println(myMap) // Output: map[apple:1 cherry:3]
}
