package main

import "fmt"

// Grouping similar variable declarations
var (
	username string
	email    string
	age      int
	isAdmin  bool
)

func main() {
	// Initializing the variables
	username = "JohnDoe"
	email = "john.doe@example.com"
	age = 30
	isAdmin = true

	// Displaying user information
	fmt.Println("Username:", username)
	fmt.Println("Email:", email)
	fmt.Println("Age:", age)
	fmt.Println("Is Admin:", isAdmin)
}
