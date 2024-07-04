package main

import "fmt"

func main() {
	// Declare and initialize variables
	var length int
	var width int

	// Prompt user for input
	fmt.Print("Enter Length: ")
	fmt.Scan(&length) // Read user input for length

	fmt.Print("Enter Width: ")
	fmt.Scan(&width) // Read user input for width

	// Calculate area
	area := length * width

	// Print area
	fmt.Println("Rectangle Area:", area)
}
