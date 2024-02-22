package main

import "fmt"

// Animal struct representing a basic animal
type Animal struct {
	Name string
}

// Dog struct embeds Animal
type Dog struct {
	Animal // Embedding Animal in Dog
	Breed  string
}

// Cat struct embeds Animal
type Cat struct {
	Animal // Embedding Animal in Cat
	Color  string
}

func main() {
	// Creating a Dog
	myDog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Labrador",
	}

	// Creating a Cat
	myCat := Cat{
		Animal: Animal{Name: "Whiskers"},
		Color:  "Orange",
	}

	// Accessing fields from embedded Animal
	fmt.Println("Dog's Name:", myDog.Name)
	fmt.Println("Cat's Name:", myCat.Name)
}
