package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func main() {
	// Create an instance of the Person struct
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}

	// Marshal the struct to JSON with field tags
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))

	// Unmarshal the JSON data back to a struct
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the values of the new struct
	fmt.Printf("Decoded Struct: %+v\n", newPerson)
}
