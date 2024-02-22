package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Converting int to string using strconv.Itoa
	number := 42
	strNumber := strconv.Itoa(number)
	fmt.Printf("Converted to string: %s\n", strNumber)

	// Converting string to int using strconv.Atoi
	strValue := "123"
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Converted to int: %d\n", intValue)

	// Parsing boolean from string using strconv.ParseBool
	boolValue, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Parsed boolean: %t\n", boolValue)
}
