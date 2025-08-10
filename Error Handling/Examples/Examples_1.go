// Handling File I/O Errors

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("readFile: failed to read file %s: %w", filename, err)
	}
	return string(content), nil
}

func main() {
	filename := "Error_Handling.txt"
	content, err := readFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("File content:", content)
}