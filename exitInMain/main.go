package main

import (
	"fmt"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("Error:", err)
		// Don't use os.Exit() here; handle the error gracefully
		return
	}

	fmt.Println("Program executed successfully.")
}

func run() error {
	// Simulating an error condition
	return fmt.Errorf("something went wrong")
}
