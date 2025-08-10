// Example 1: Counting Word Frequencies
// This example demonstrates how to use a map to count the frequency of each word in a given string.

package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "This is a sample text. This text is used to demonstrate word frequency counting."
    words := strings.Fields(text) // Split the text into words

    wordFrequencies := make(map[string]int)

    for _, word := range words {
        // Convert the word to lowercase to count case-insensitively
        word = strings.ToLower(word)
        wordFrequencies[word]++ // Increment the count for the word
    }

    // Print the word frequencies
    for word, frequency := range wordFrequencies {
        fmt.Printf("Word: %s, Frequency: %d\n", word, frequency)
    }
}
// In this example, we split the input text into words using strings.Fields. Then, we iterate over the words and use a map to store the frequency of each word. The strings.ToLower function ensures that the counting is case-insensitive.


