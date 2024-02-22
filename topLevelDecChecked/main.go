package main

import "fmt"

// Top-level variable declaration
var globalVariable = "I am a global variable"

func main() {
	// Using the global variable
	fmt.Println(globalVariable)

	// Calling a function that uses a local variable
	localVariableFunction()
}

// Function that uses a local variable
func localVariableFunction() {
	// Local variable
	localVariable := "I am a local variable"

	// Using the local variable
	fmt.Println(localVariable)
}
