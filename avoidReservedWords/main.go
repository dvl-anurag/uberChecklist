package main

import (
	"fmt"
)

// Avoiding built-in names
type string int // Not recommended

func main() {
	// Using the custom type named string
	var value string = 42 // Using the custom type
	fmt.Println(value)
}
