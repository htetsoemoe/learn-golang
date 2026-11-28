// Two Receivers in Go method
// Value Receiver and Pointer Receiver
// Value Receiver: gets struct copy and changes from this method can't change original struct.
// Pointer Receiver: gets original struct memory address (Pointer) and change original struct.

package main
import "fmt"

type Person struct {
	FirstName string
	Age int
}

// (p Person) -> Value Receiver
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.FirstName, p.Age)
	p.Age = 99 // this change can affect Person struct copy and can't effect original struct
}

// (p *Person) -> Pointer Receiver
func (p *Person) SetAge(newAge int) {
	p.Age = newAge
}

func main() {
	p := Person{
		FirstName: "John",
		Age: 25,
	}

	p.Greet()
	fmt.Println("Age after Greet():", p.Age) // can't change original Person struct

	p.SetAge(30)
	fmt.Println("Age after SetAge():", p.Age) // change original Person struct
}