package main

import "fmt"

func main() {
	// Good: Variable 'message' is declared with the smallest necessary scope
	if true {
		message := "Hello, World!"
		fmt.Println(message)
	}

	// Uncommenting the line below would result in a compilation error
	// fmt.Println(message) // 'message' undefined

	// Bad: Variable 'greeting' has a larger scope than necessary
	var greeting string

	if true {
		greeting = "Hi there!"
	}

	fmt.Println(greeting)
}
