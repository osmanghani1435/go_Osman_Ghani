package main

import "fmt"

func main() {
	n := 3 // Define the value of n

	// Print pairs (i, j) where 1 <= i <= n and 1 <= j <= i
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Println(i, j)
		}
	}
}
