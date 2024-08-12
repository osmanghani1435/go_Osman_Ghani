// package main

// import "fmt"

//	func main() {
//		sum := 1
//		for sum < 10 {
//			sum += 10
//			fmt.Println(sum)
//		}
//	}
package main

import (
	"fmt"
	"sort"
)

// Function to get the optimal arrangement
func getOptimalArrangement(power []int) []int {
	// Sort the array in descending order
	sort.Slice(power, func(i, j int) bool {
		return power[i] > power[j]
	})
	return power
}

func main() {
	// Example usage
	var n int
	fmt.Scan(&n)
	power := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&power[i])
	}
	result := getOptimalArrangement(power)
	for _, val := range result {
		fmt.Println(val)
	}
}
