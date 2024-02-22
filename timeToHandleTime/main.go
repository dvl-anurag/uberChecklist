package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	now := time.Now()
	fmt.Println("Current Time:", now)

	// Format time as a string
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formattedTime)

	// Add 2 hours to the current time
	futureTime := now.Add(2 * time.Hour)
	fmt.Println("Time after 2 hours:", futureTime)

	// Calculate the duration between two times
	duration := futureTime.Sub(now)
	fmt.Println("Duration between two times:", duration)

	// Sleep for 3 seconds
	fmt.Println("Sleeping for 3 seconds...")
	time.Sleep(3 * time.Second)
	fmt.Println("Awake!")

	// Timer example
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("Timer expired!")
}
