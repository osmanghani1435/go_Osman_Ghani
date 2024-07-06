package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// First input
	fmt.Println("Enter love:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "love" {
		fmt.Println("Love is a Pandemic")
		fmt.Println("that has two types")
		fmt.Println("Negative and Positive")
		fmt.Println("Originated by Osman Ghani")
		fmt.Println("if you enjoy something which is really hard, it means you're really strong")
	} else {
		fmt.Println("Love is like your typo mistake")
		fmt.Println("Just write 'love' in lowercase")
		fmt.Println("or just type in lowercase")
		fmt.Println("what is love")
	}

	// Second input
	fmt.Println("Enter what is love:")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "what is love" {
		fmt.Println("Love is an unexplainable feeling")
		fmt.Println("do love, feel love")
		fmt.Println("be positive, be long lasting")
		fmt.Println("no one is trusted")
	} else {
		fmt.Println("Love is like your typo mistake")
		fmt.Println("In lowercase, just write 'what is love'")
		fmt.Println("or in lowercase, just type 'love'")
		fmt.Println("Originated by Osman just for")
		fmt.Println("if you enjoy something that is really hard, it means you're really strong")
	}
}
