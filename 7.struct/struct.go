package main
import "fmt"

// Create 'Person' struct
type Person struct {
	FirstName string
	LastName string
	Age int
}

func main() {
	// Create a person from Person struct
	p1 := Person{
		FirstName: "John",
		LastName: "Doe",
		Age: 30,
	}

	fmt.Println("Person 1:", p1)
	fmt.Println("First Name:", p1.FirstName)

	// Change Person struct field value
	p1.Age = 35
	fmt.Println("Updated Age:", p1.Age)
}