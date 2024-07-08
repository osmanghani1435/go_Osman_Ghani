package main

import (
	"fmt"
	"math"
)

func main() {
	var radius, height float64

	// Prompt user for radius
	fmt.Print("Enter the radius of the tube: ")
	fmt.Scan(&radius)

	// Prompt user for height
	fmt.Print("Enter the height of the tube: ")
	fmt.Scan(&height)

	// Calculate the volume of the tube
	volume := math.Pi * math.Pow(radius, 2) * height

	// Print the calculated volume
	fmt.Printf("The volume of the tube is: %.2f cubic units\n", volume)
}
