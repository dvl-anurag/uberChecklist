package main

import "fmt"

// Correct: naming the function with 'f' suffix
func printDetailsF(name string, age int) {
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

func main() {
	// Calling the correctly named function
	printDetailsF("John", 30)
}
