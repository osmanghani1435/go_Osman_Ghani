package main

import "fmt"

// type cousrse struct {
// 	Name        string
// 	Description string
// 	Code        string
// }

func main() {

	length := 5
	width := 6
	area := calculateRectangleArea(length, width)
	fmt.Println(area)

	// 	name := "osman"
	// 	description := "lets learn "
	// 	couserseCode := "c001"

	// 	course := course{

	// 		Name:        name,
	// 		Description: description,
	// 		Code:        code,
	// 	}
	// 	fmt.Println(course)

}

func calculateRectangleArea(length, width int) int {

	return length * width

}
