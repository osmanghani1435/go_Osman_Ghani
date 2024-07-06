package main

import "fmt"

func main() {
	hour := 15
	if hour < 12 {
		fmt.Println("morning")

	} else if hour < 18 {
		fmt.Println("overtime")

	}
}
