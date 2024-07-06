package main

import "fmt"

func main() {
	var myAge int
	fmt.Println("Enter your age:")
	fmt.Scan(&myAge)

	if myAge == 5 {
		fmt.Println("You're too young")
	}

	if myAge == 17 {
		fmt.Println("So sweet")
	}

	if myAge >= 17 && myAge <= 30 {
		fmt.Println("My age is between 17 and 30")
	}

	if dadAge := 9; dadAge < myAge {
		fmt.Println("Dad's age is", dadAge)
	}
}
