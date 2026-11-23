package main
import "fmt"

func main() {
	// numbers := []int{10, 20, 30, 40, 50}

	// len() returns the number of elements currently in the slice
	// cap() returns the capacity of the underlying array from the slice's starting position
	numbers := make([]int, 3, 5)
	fmt.Println("Length of slice:", len((numbers)))
	fmt.Println("Capacity of slice", cap(numbers))

	numbers = append(numbers, 10, 20, 30)
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Slice Expression
	fmt.Println("Slice Expression", numbers[3:5])


	// Copying Slices
	source := []int{10, 20, 30}
	destination := make([]int, len(source))

	/*
	Important: Assigning one slice to another does not necessarily create an independent copy.
	Both slice headers can refer to the same underlying array.
	*/
	copy(destination, source)

	destination[0] = 100

	fmt.Println("Source", source)
	fmt.Println("Destination", destination)
}