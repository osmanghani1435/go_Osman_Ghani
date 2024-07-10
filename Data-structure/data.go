package main

import "fmt"

func main() {
	// arey store many items size .. fixed size..
	// how we create arrey
	// osman Ghani
	// approach 1 empty arrey

	var numbers [5]int

	// insert items
	numbers[0] = 12
	numbers[1] = 14
	// numbers [2] = 46
	// numbers [3] = 50

	// approachg for accessing arrey regular for loop
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])

	}
	// approach 2

	for idx, val := range numbers {
		fmt.Println("the index: ", idx, "the value: ", val)

	}
}
