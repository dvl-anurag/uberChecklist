package main

import "fmt"

// MyType is a custom type with a method that has a value receiver.
type MyType struct {
	Value int
}

// ValueMethod is a method with a value receiver.
func (m MyType) ValueMethod() {
	fmt.Println("ValueMethod called with value:", m.Value)
}

// PointerMethod is a method with a pointer receiver.
func (m *MyType) PointerMethod() {
	fmt.Println("PointerMethod called with pointer:", m.Value)
}

func main() {
	// Create an instance of MyType
	myValue := MyType{Value: 42}

	// Call ValueMethod with a value receiver using the value
	myValue.ValueMethod()

	// Call ValueMethod with a value receiver using a pointer
	(&myValue).ValueMethod()

	// Call PointerMethod with a pointer receiver using the value
	// This works because Go automatically takes the address of myValue
	myValue.PointerMethod()

	// Create a pointer to MyType
	myPointer := &MyType{Value: 99}

	// Call PointerMethod with a pointer receiver using the pointer
	myPointer.PointerMethod()

	// Call PointerMethod with a pointer receiver using the value
	// This works because Go automatically dereferences myPointer
	(*myPointer).PointerMethod()
}
