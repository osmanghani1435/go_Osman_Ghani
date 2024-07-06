package main

import (
	"fmt"
)

func main() {
outerLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				break outerLoop
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}
