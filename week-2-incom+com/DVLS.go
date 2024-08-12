package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pickVocals("alterra academy"))
	fmt.Println(pickVocals("i love programming"))
	fmt.Println(pickVocals("go is awesome programming language"))
}

func pickVocals(sentence string) string {
	vowels := "aeiou"
	result := ""

	words := strings.Split(sentence, " ")

	for _, word := range words {
		vocalPart := ""
		for _, char := range word {

			if strings.Contains(vowels, string(char)) {
				vocalPart += string(char)
			}
		}
		result += vocalPart + " "
	}

	return strings.TrimSpace(result)
}
