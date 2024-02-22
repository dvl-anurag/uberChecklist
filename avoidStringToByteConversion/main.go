package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// String to byte slice without conversion
	str := "Hello, 世界!"
	byteSlice := []byte(str)

	// Iterate over bytes in the slice
	for _, b := range byteSlice {
		fmt.Printf("%c ", b)
	}
	fmt.Println()

	// Byte slice to string without conversion
	newStr := string(byteSlice)

	// Print the original and reconstructed strings
	fmt.Println("Original String:", str)
	fmt.Println("Reconstructed String:", newStr)

	// Counting characters in a string without converting to bytes
	charCount := utf8.RuneCountInString(str)
	fmt.Println("Number of characters in the string:", charCount)
}
