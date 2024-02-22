package main

import (
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Program executed successfully.")
}

func run() error {
	// Simulating an error condition
	return fmt.Errorf("something went wrong")
}
