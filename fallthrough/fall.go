package main

import "fmt"

func main() {

	value := 42
	switch value {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
	default:
		fmt.Println("default")
	}
}
