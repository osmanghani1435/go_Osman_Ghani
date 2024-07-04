package main

import (
	"fmt"
)

func main() {
	var side int

	fmt.Print("Enter the side length of the square: ")
	fmt.Scan(&side)

	area := side * side // Calculate the area of the square

	fmt.Println("Area of the square:", area)

	if area > 20 {
		fmt.Println("Cool square!")
	} else {
		fmt.Println("Uncool square!")
	}
}
