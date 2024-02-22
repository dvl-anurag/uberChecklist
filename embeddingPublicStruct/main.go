package main

import (
	"fmt"
)

// Embedding a type in a public struct
type User struct {
	ID      int
	Profile UserProfile
}

// Embedded type with its own fields
type UserProfile struct {
	Username string
	Email    string
}

func main() {
	// Creating a user with embedded profile
	user := User{
		ID: 1,
		Profile: UserProfile{
			Username: "john_doe",
			Email:    "john@example.com",
		},
	}

	// Accessing fields directly (not recommended)
	fmt.Println("User ID:", user.ID)
	fmt.Println("Username:", user.Username) // Compilation error

	// Accessing fields through embedded type (recommended)
	fmt.Println("User ID:", user.ID)
	fmt.Println("Username:", user.Profile.Username)
}
