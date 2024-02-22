package main

import (
	"fmt"
)

func processInterfaceValue(value interface{}) {
	// Attempting type assertion
	if str, ok := value.(string); ok {
		// Type assertion successful
		fmt.Println("String Value:", str)
	} else {
		// Type assertion failed
		fmt.Println("Not a string value")
	}
}

func main() {
	// Example 1: Type assertion succeeds
	processInterfaceValue("Hello, World!")

	// Example 2: Type assertion fails
	processInterfaceValue(42)
}
