package main

import (
	"testing"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		inputA   int
		inputB   int
		expected int
	}{
		{"Positive numbers", 3, 5, 8},
		{"Negative numbers", -2, -4, -6},
		{"Zero values", 0, 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := add(test.inputA, test.inputB)
			if result != test.expected {
				t.Errorf("For inputs %d and %d, expected %d, but got %d", test.inputA, test.inputB, test.expected, result)
			}
		})
	}
}
