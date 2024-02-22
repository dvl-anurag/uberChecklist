package main

import (
	"fmt"
)

// Avoiding init() and using main()
func main() {
	initMessage()
}

func initMessage() {
	fmt.Println("This message is part of initialization.")
}
