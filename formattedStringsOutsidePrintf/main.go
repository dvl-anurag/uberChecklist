package main

import "fmt"

func main() {
	// Incorrect: string concatenation
	name := "John"
	age := 30
	// This is less readable and can be error-prone
	result := "Name: " + name + ", Age: " + fmt.Sprint(age)
	fmt.Println(result)

	// Correct: formatted string using fmt.Printf
	// This is more readable and idiomatic
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}
