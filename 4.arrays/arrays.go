package main
import "fmt"

// Declare fix-sized array
// var numbers [5]int


func main() {
	// numbers[0] = 10
	// numbers[1] = 20
	// numbers[2] = 30
	// numbers[3] = 40
	// numbers[4] = 50

	numbers := [3]int {1, 3, 5}

	for i := 0; i < 3; i++ {
		fmt.Println(numbers[i])
	}

	// Using like while loop and there is no while loop in Go
	n := 0
	for n < 5 {
		fmt.Println(n)
		n++
	}


	// for-range loop (For Slices, Maps, Arrays)
	heros := [3]string {"Superman", "Batman", "Ironman"}
	for _, hero := range heros{
		fmt.Println(hero)
	}
}