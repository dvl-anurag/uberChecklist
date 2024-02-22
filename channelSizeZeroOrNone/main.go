package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create an unbuffered channel
	ch := make(chan int)

	// Use a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Start a goroutine to send data to the channel
	wg.Add(1)
	go sendData(ch, &wg)

	// Start a goroutine to receive data from the channel
	wg.Add(1)
	go receiveData(ch, &wg)

	// Wait for goroutines to finish
	wg.Wait()
}

// sendData sends data to the channel
func sendData(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	data := 42
	fmt.Println("Sending data to channel:", data)
	ch <- data
}

// receiveData receives data from the channel
func receiveData(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	data := <-ch
	fmt.Println("Received data from channel:", data)
}
