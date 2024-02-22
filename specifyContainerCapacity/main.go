package main

import "fmt"

func main() {
	// Creating a slice with a specified capacity
	capacity := 5
	slice := make([]int, 0, capacity)

	// Appending elements to the slice
	for i := 1; i <= capacity; i++ {
		slice = append(slice, i)
	}

	// Printing the elements of the slice
	fmt.Println("Slice Elements:", slice)
	fmt.Println("Slice Length:", len(slice))
	fmt.Println("Slice Capacity:", cap(slice))
}
