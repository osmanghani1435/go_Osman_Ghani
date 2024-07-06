package main

import (
	"fmt"
)

func main() {
	sentence := "hello"

	// Print each character in the string using a traditional for loop
	for i := 0; i < len(sentence); i++ {
		fmt.Println(string(sentence[i]))
	}
	fmt.Println("---------")

	// Print each character and its position using a range-based for loop
	for pos, char := range sentence {
		fmt.Printf("Character: %c at position: %d\n", char, pos)
	}
}
