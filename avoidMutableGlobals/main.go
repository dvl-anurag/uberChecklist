package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	globalCounter int
	counterMutex  sync.Mutex
)

func incrementGlobalCounter() {
	counterMutex.Lock()
	defer counterMutex.Unlock()

	globalCounter++
}

func printGlobalCounter() {
	counterMutex.Lock()
	defer counterMutex.Unlock()

	fmt.Println("Global Counter:", globalCounter)
}

func main() {
	// Incrementing global counter in a goroutine
	go func() {
		for i := 0; i < 5; i++ {
			incrementGlobalCounter()
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// Printing global counter
	printGlobalCounter()

	// Wait for the goroutine to finish
	time.Sleep(time.Second)
}
