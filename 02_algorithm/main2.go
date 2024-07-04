package main

import "fmt"

func main() {
	num := 12 // Example value

	// Loop from 1 to 5
	for i := 1; i <= 5; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}
}
