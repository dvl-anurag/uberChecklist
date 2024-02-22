package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Zero-value Mutex
	var mu sync.Mutex

	// Attempt to lock the zero-value Mutex
	mu.Lock() // This line will cause a run-time panic

	// Example usage of a valid Mutex
	var validMutex sync.Mutex
	validMutex.Lock()
	fmt.Println("Locked successfully")
	time.Sleep(2 * time.Second)
	validMutex.Unlock()
	fmt.Println("Unlocked successfully")
}
