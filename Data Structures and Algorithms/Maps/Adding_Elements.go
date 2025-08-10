// Adding Elements: Add new key-value pairs to the map.


package main

import "fmt"

func main() {
    myMap := map[string]int{
        "apple":  1,
        "banana": 2,
        "cherry": 3,
    }

    // Add a new key-value pair
    myMap["grape"] = 4
    fmt.Println(myMap) // Output: map[apple:1 banana:2 cherry:3 grape:4]
}
