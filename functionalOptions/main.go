package main

import "fmt"

// Config defines the configuration structure
type Config struct {
	Timeout    int
	RetryCount int
	Debug      bool
}

// OptionFunc defines the functional option type
type OptionFunc func(*Config)

// WithTimeout sets the timeout option
func WithTimeout(timeout int) OptionFunc {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithRetryCount sets the retry count option
func WithRetryCount(retryCount int) OptionFunc {
	return func(c *Config) {
		c.RetryCount = retryCount
	}
}

// WithDebug enables or disables debugging
func WithDebug(debug bool) OptionFunc {
	return func(c *Config) {
		c.Debug = debug
	}
}

// NewConfig creates a new Config instance with functional options
func NewConfig(options ...OptionFunc) *Config {
	// Default configuration
	config := &Config{
		Timeout:    10,
		RetryCount: 3,
		Debug:      false,
	}

	// Apply functional options
	for _, option := range options {
		option(config)
	}

	return config
}

func main() {
	// Create a new Config instance with functional options
	config := NewConfig(
		WithTimeout(20),
		WithRetryCount(5),
		WithDebug(true),
	)

	// Print the configured values
	fmt.Printf("Timeout: %d, RetryCount: %d, Debug: %t\n", config.Timeout, config.RetryCount, config.Debug)
}
