package main

import (
	"fmt"
	"io"
)

// MyWriter is a custom type that implements the io.Writer interface.
type MyWriter struct{}

// Write is the implementation of the io.Writer interface.
func (w MyWriter) Write(p []byte) (n int, err error) {
	// Implement the Write method logic
	return len(p), nil
}

// Ensure MyWriter complies with the io.Writer interface at compile time.
var _ io.Writer = MyWriter{}

func main() {
	// Create an instance of MyWriter
	myWriter := MyWriter{}

	// Use MyWriter as an io.Writer
	bytesWritten, err := myWriter.Write([]byte("Hello, world!"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Bytes written:", bytesWritten)
}
