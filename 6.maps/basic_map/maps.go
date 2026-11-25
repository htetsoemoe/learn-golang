package main
import "fmt"

func main() {
	// map[string]string means Key type: string, Value type: string
	user := map[string]string{
		"name": "soe",
		"city": "Yangon",
	}

	fmt.Println("Name:", user["name"]) // Accessing value

	// Adding and updating values
	user["email"] = "soe@example.com"
	user["city"] = "Mandalay"

	fmt.Println("Email:", user["email"], "\nCity:", user["city"])

	// Go uses the comma-ok idiom to check whether a key exists.
	// A standard Go pattern where an operation returns two values: the result and a boolean (ok)
	// indicating whether the operation succeeded.
	value, exists := user["name"]

	if exists {
		fmt.Println("Name:", value)
	} else {
		fmt.Println("Not found")
	}

	// Iterating over maps
	for key, value := range user {
		fmt.Println(key, value)
	}
}