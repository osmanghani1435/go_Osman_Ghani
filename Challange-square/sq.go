package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Enter a number: ")
	fmt.Scan(&N)

	fmt.Printf("Squares up to %d:\n", N)
	for i := 1; i <= N; i++ {
		fmt.Printf("%d ", i*i)
	}
	fmt.Println()
}
