package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{45, 6 - 99, 12, 13}

	sort.Ints(numbers)
	fmt.Println("sorted", numbers)

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[j]
	})

	fmt.Println("reversed", numbers)

	fmt.Println("item", item)

	otherNuns := []int{4, 3, 2, 77, 28, 1}
	sorte.SearchInt(numbers, 4)

}

// {
// 	func factorial(n int) {
// 		if n == 1 {
// 			return 1

// 		}
// 		return n *
// 	}
// }
