package main

import (
	"errors"
	"fmt"
)

func processData(data []int) error {
	if len(data) == 0 {
		return errors.New("input data is empty")
	}

	// Simulate processing data (here, we assume some processing logic)
	// In a real-world scenario, this could be more complex operations.
	processedData, err := performProcessing(data)
	if err != nil {
		// Wrap the error to provide more context
		return fmt.Errorf("error processing data: %w", err)
	}

	// Simulate using the processed data (not shown in detail)
	_ = useProcessedData(processedData)

	return nil
}

func performProcessing(data []int) (int, error) {
	// Simulate some processing logic
	if len(data) == 1 {
		return 0, errors.New("insufficient data for processing")
	}

	// In a real scenario, this function would perform more complex operations.
	// Here, we'll just return the sum for simplicity.
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum, nil
}

func useProcessedData(result int) error {
	// Simulate using the processed data
	fmt.Println("Processed Result:", result)
	return nil
}

func main() {
	// Example 1: Processing non-empty data
	data1 := []int{1, 2, 3}
	if err := processData(data1); err != nil {
		fmt.Println("Error:", err)
	}

	// Example 2: Processing empty data
	data2 := []int{}
	if err := processData(data2); err != nil {
		fmt.Println("Error:", err)
	}
}
