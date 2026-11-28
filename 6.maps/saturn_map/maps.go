package main
import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"Alice": 25,
		"Bob": 27,
	}

	fmt.Println("Initial Map:", ages)

	// Delete a element from Map
	delete(ages, "Alice")
	fmt.Println("After delete:", ages)

	// Check there is no key exist or not
	// 'val' gets value and 'ok' gets boolean value 
	val, ok := ages["Cherry"]
	if ok {
		fmt.Println("Cherry's age is", val)
	} else {
		fmt.Println("User not found.")
	}

	// Map Iteration
	// Unordered iteration (Default feature for Maps)
	testScores := map[string]int {
		"Charlie": 88,
		"Alice": 95,
		"Bob": 80,
	}
	fmt.Println("\nUnordered Iteration:")
	for key, score := range testScores {
		fmt.Printf("\nName: %s, Score: %d", key, score)
	}
	fmt.Println("\n_____________________")

	// Sorted Iteration
	// 1. Need to keep keys in slice
	var names []string
	for key := range testScores {
		names = append(names, key)
	}

	// 2. Sort the slice
	sort.Strings(names)

	fmt.Println("\nOrdered Iteration:")

	// 3. Loop map with sorted slice
	for _, name := range names {
		fmt.Printf("\nName: %s, Score: %d", name, testScores[name])
	}
}