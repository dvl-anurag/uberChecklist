package main

import "fmt"

func main() {
	// Good: Group similar variable declarations
	var (
		name   string
		age    int
		height float64
	)

	// Initializing values
	name = "John"
	age = 25
	height = 175.5

	// Printing values
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)

	// Bad: Avoid separate variable declarations unless they are unrelated
	var country string
	var weight float64

	// Initializing values
	country = "USA"
	weight = 70.5

	// Printing values
	fmt.Println("Country:", country)
	fmt.Println("Weight:", weight)
}
