package main

import "fmt"

// Function parameters
func greet(name string) {
	fmt.Println("Hello ", name)
}

// You can declare multiple parameters with the same type and return
func add(a, b int) int {
	return  a + b
}

// Multiple return values
func calculate(a, b int) (int, int) {
	sum := a + b
	difference := a - b

	return  sum, difference
}

// Named return values
func namedCalculate(a, b int) (sum int, different int) {
	sum = a + b
	different = a - b

	return
}

// Variadic functions: A variadic function accepts a variable number of arguments.
func adder(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return  total
}

// Closures
func counter() func() int {
	count := 0

	return  func() int {
		count ++
		return  count
	}
}

func main() {
	greet("Ko Htet")

	result := add(30, 6)
	fmt.Println(result)

	sum, difference := calculate(10, 5)
	fmt.Println("Sum", sum, "Difference", difference)

	sumOne, different := namedCalculate(10, 5)
	fmt.Println(sumOne, different)

	rstOne := adder(1, 2, 3)
	rstTwo := adder(10, 10, 10, 10, 10)
	fmt.Println("ResultOne", rstOne, "ResultTwo", rstTwo)

	// Anonymous function
	sayHello := func(name string) {
		fmt.Println("Hello", name)
	}
	sayHello("Ko Htet")

	// Closures
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

}
