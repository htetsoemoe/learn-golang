package main

import "fmt"

func main() {
	day := "Monday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println(day, "It's the weekend!")
	case "Monday":
		fmt.Println(day, "It's the start of the week.")
	default:
		fmt.Println(day, "It's a weekday.")
	}

	// fallthrough (keyword): if one case block is true and proceed following cases
	role := "admin"

	switch role {
	case "admin":
		fmt.Println(role, "-- Can Delete Users")
		fallthrough // if user role is admin, can edit content
	case "editor":
		fmt.Println(role, "-- Can Edit Content")
		fallthrough // if user role is editor, can view content
	case "viewer":
		fmt.Println(role, "-- Can View Content")
	default:
		fmt.Println(role, "-- No Access")
	}
}