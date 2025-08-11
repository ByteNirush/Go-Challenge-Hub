package main

import (
	"fmt"
	"encoding/json"
)

type Product struct {
	ID			int 	`json:"id"`
	Name		string 	`json:"name"`
	Price		float64 `json:"price"`
	Description string 	`json:"description,omitempty"`
}

func main() {
	product := Product{
		ID:		1,
		Name: 	"Laptop",
		Price:	999.99,
	}

	// Marshal (encode) the struct to JSON
	jsonData, err := json.Marshal(product)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	fmt.Println("JSON Data:", string(jsonData))

	// Unmarshal (decode) the JSON back to a struct
	var newProduct Product
	err = json.Unmarshal(jsonData, &newProduct)
	if err != nil {
		fmt.Println("Error unmarshaling from JSON:", err)
		return
	}
	fmt.Println(newProduct)
}
