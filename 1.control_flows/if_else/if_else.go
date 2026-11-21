package main

import "fmt"

func main() {
	// Short Variable Declaration (:=)
	// Important: You can use := inside functions, but not directly at package level.
	/*
	age := 30

	// Unlike JavaScript, Go does not require parentheses around conditions
	if age >= 18 {
		fmt.Println("Your are an adult.")
	} else {
		fmt.Println("Your are a minor.")
	}
	*/

	// We can use num variable in if/else block only
	if num := 9; num < 0 {
		fmt.Println(num, "is negative.")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit.")
	} else {
		fmt.Println(num, "has multiple digits.")
	}

}