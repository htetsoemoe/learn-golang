package main

import "fmt"

func main() {
	// for loop
	/*
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	*/

	// while loop
	/*
	n := 0
	for n < 5 {
		fmt.Println(n)
		n++
	}
	*/

	// for-range loop (Slices, Maps, Arrays)
	fruits := []string{"Apple", "Banana", "Cherry"}
	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}