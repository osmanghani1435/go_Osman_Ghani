package main

import "fmt"

func main() {
	fmt.Println(groupPalindrome([]string{"katak", "civic", "kasur", "rusak"}))
	fmt.Println(groupPalindrome([]string{"racecar", "seru", "kasur", "civic", "bilik", "kak"}))
	fmt.Println(groupPalindrome([]string{"masuk", "civic", "hahah", "garam"}))
}

func groupPalindrome(words []string) [][]string {
	var palindromes []string
	var nonPalindromes []string

	for _, word := range words {
		if isPalindrome(word) {
			palindromes = append(palindromes, word)
		} else {
			nonPalindromes = append(nonPalindromes, word)
		}
	}

	result := [][]string{}
	if len(palindromes) > 0 {
		result = append(result, palindromes)
	}
	for _, word := range nonPalindromes {
		result = append(result, []string{word})
	}
	return result
}

func isPalindrome(word string) bool {

	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}
