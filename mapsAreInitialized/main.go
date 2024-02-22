package main

import "fmt"

func main() {
	// Incorrect: uninitialized map will panic if used
	// var myMap map[string]int

	// Correct: initialize the map before using it
	myMap := make(map[string]int)

	// Now you can safely use the map
	myMap["one"] = 1
	myMap["two"] = 2

	// Print the values
	fmt.Println("Value of 'one':", myMap["one"])
	fmt.Println("Value of 'two':", myMap["two"])
}
