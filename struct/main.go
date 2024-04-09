package main

import "fmt"

// Define a struct named Person
type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {
	// Create a new instance of the Person struct
	person1 := Person{
		"Alice",
		30,
		"USA",
	}

	// Accessing struct fields
	fmt.Printf("Person 1: %+v\n", person1)

    // Print the struct using %v to display only values
    fmt.Printf("Person 1 (values only): %v\n", person1.Age)

    // Update struct fields
    person1.Age = 31

    // Print the updated struct
    fmt.Printf("Updated Person 1: %+v\n", person1)

    // Declare and initialize struct without specifying field names (order matters)
    person2 := Person{"Bob", 25, "Canada"}

    // Print the struct using %+v to display field names
    fmt.Printf("Person 2: %+v\n", person2)

    // Print the struct using %v to display only values
    fmt.Printf("Person 2 (values only): %v\n", person2)
}
