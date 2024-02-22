package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a time representing the current moment
	now := time.Now()
	fmt.Println("Current Time:", now)

	// Create a specific time using time.Date
	specificTime := time.Date(2023, time.May, 15, 12, 30, 0, 0, time.UTC)
	fmt.Println("Specific Time:", specificTime)

	// Perform operations with time
	oneWeekLater := now.AddDate(0, 0, 7)
	fmt.Println("One week later:", oneWeekLater)

	// Compare two times
	isAfter := specificTime.After(now)
	fmt.Println("Is specificTime after now?", isAfter)

	// Format time as a string
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formattedTime)
}
