package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Config represents a configuration with time-related fields
type Config struct {
	StartTime time.Time     `json:"startTime"`
	Interval  time.Duration `json:"interval"`
}

func main() {
	// Creating a sample configuration
	config := Config{
		StartTime: time.Now(),
		Interval:  5 * time.Second,
	}

	// Convert the Config struct to JSON
	configJSON, err := json.Marshal(config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Config JSON:", string(configJSON))

	// Simulate reading the configuration from an external source (e.g., JSON)
	// In a real-world scenario, this might be reading from a file or an API response.
	var retrievedConfig Config
	if err := json.Unmarshal(configJSON, &retrievedConfig); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the retrieved configuration
	fmt.Println("Retrieved Config:")
	fmt.Println("Start Time:", retrievedConfig.StartTime)
	fmt.Println("Interval:", retrievedConfig.Interval)
}
