package main

import "fmt"

func main() {
	var today int = 1 // Assign an appropriate value to `today`

	switch today {
	case 1:
		fmt.Println("Today is Monday")
	case 2:
		fmt.Println("Today is Tuesday")
	default:
		fmt.Println("Invalid Date")
	}
}
