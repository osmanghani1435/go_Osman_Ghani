package main

import "fmt"

func main() {
	fmt.Println(spinString("alta"))    // Output: aalt
	fmt.Println(spinString("alterra")) // Output: aalrtre
	fmt.Println(spinString("sepulsa")) // Output: saesplu
}

func spinString(word string) string {
	// Take the last character and add it to the front
	if len(word) > 1 {
		return word[len(word)-1:] + word[:len(word)-1]
	}
	return word
}
