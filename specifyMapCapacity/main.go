package main

import "fmt"

func main() {
	// Specifying the capacity hint for the map
	capacity := 5
	myMap := make(map[string]int, capacity)

	// Adding elements to the map
	for i := 1; i <= capacity; i++ {
		key := fmt.Sprintf("Key%d", i)
		myMap[key] = i
	}

	// Printing the elements of the map
	fmt.Println("Map Elements:", myMap)
}
