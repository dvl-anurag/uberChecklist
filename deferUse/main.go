package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "example.txt"

	// Open a file
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer closeFile(file) // Defer the closeFile function call

	// Writing data to the file
	data := []byte("Hello, Defer!")
	_, writeErr := file.Write(data)
	if writeErr != nil {
		fmt.Println("Error writing to file:", writeErr)
		return
	}

	// Rest of the program
	fmt.Println("Data written to file successfully.")
}

// closeFile is a helper function to close the file.
func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
	}
}
