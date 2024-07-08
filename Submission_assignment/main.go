package main

import (
	"fmt"
)

func main() {
	// Define Phi (π)
	const Phi = 3.14

	// Define the radius and height of the tube
	var radius, height float64

	// Ask the user to input the radius and height
	fmt.Print("Enter the radius of the tube: ")
	fmt.Scan(&radius)
	fmt.Print("Enter the height of the tube: ")
	fmt.Scan(&height)

	// Calculate the volume of the tube
	volume := Phi * radius * radius * height

	// Print the result
	fmt.Printf("The volume of the tube is: %.2f\n", volume)
}
