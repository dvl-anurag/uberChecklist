package main

import "fmt"

// Enum for days of the week
const (
	Sunday    = 1
	Monday    = 2
	Tuesday   = 3
	Wednesday = 4
	Thursday  = 5
	Friday    = 6
	Saturday  = 7
)

func main() {
	// Example usage
	day := Tuesday

	switch day {
	case Sunday:
		fmt.Println("It's Sunday!")
	case Monday:
		fmt.Println("It's Monday!")
	case Tuesday:
		fmt.Println("It's Tuesday!")
		// ... handle other days
	default:
		fmt.Println("It's another day.")
	}
}
