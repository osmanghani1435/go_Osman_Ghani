package main

import (
	"fmt"
)

// fizzBuzz function takes an integer n and prints the appropriate response for each value in the range {1, 2, ..., n}
func fizzBuzz(n int32) {
	for i := int32(1); i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	var n int32
	fmt.Print("Enter the value of n: ")
	fmt.Scan(&n)
	fizzBuzz(n)
}
