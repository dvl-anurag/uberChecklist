package main

import "fmt"

// DataService represents a service with internal state.
type DataService struct {
	data map[string]int
}

// NewDataService creates a new DataService instance.
func NewDataService() *DataService {
	return &DataService{
		data: map[string]int{"key1": 1, "key2": 2, "key3": 3},
	}
}

// GetData returns a copy of the internal data to avoid exposing the original map.
func (ds *DataService) GetData() map[string]int {
	dataCopy := make(map[string]int, len(ds.data))
	for key, value := range ds.data {
		dataCopy[key] = value
	}
	return dataCopy
}

func main() {
	dataService := NewDataService()

	// Attempt to modify the internal state directly (which is not allowed)
	// Uncommenting the next line would result in a compilation error:
	// dataService.data["key1"] = 100

	// Retrieve a copy of the data through a safe method
	dataCopy := dataService.GetData()

	// Modify the copy safely
	dataCopy["key1"] = 100

	// Original internal state remains unchanged
	fmt.Println("Original Data:", dataService.data)
	fmt.Println("Modified Copy:", dataCopy)
}
