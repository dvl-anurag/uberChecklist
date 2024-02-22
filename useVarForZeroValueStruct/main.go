package main

import "fmt"

// Define a simple struct
type Point struct {
	X int
	Y int
}

func main() {
	// Declare a variable of type Point without assigning a value
	var p Point

	// Print the zero values of the struct fields
	fmt.Println("X:", p.X) // Zero value for int is 0
	fmt.Println("Y:", p.Y) // Zero value for int is 0
}
