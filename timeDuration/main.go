package main

import (
	"fmt"
	"time"
)

func main() {
	// Define a duration of 2 hours
	duration := 2 * time.Hour

	fmt.Println("Duration:", duration)

	// Sleep for the specified duration
	fmt.Println("Sleeping for 2 seconds...")
	time.Sleep(duration)
	fmt.Println("Awake!")

	// Convert seconds to duration
	seconds := 120
	secondsDuration := time.Duration(seconds) * time.Second
	fmt.Println("Duration from seconds:", secondsDuration)
}
