package main

import "fmt"

func main() {
	// Specifying the capacity hint for the slice
	capacity := 5
	mySlice := make([]int, 0, capacity)

	// Adding elements to the slice
	for i := 1; i <= capacity; i++ {
		mySlice = append(mySlice, i)
	}

	// Printing the elements of the slice
	fmt.Println("Slice Elements:", mySlice)
}
