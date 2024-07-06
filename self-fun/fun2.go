package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "love" {
		fmt.Println("Love is a pandemic")
	} else {
		fmt.Println("Try again, typo mistake")
	}
}
