package main

import "fmt"

func main() {
	word := "racecar"
	result := isPalindrome(word)

	fmt.Printf("%v is a palindrome: %v\n", word, result)
}

func isPalindrome(word string) bool {
	wordLength := len(word)

	for i := 0; i < wordLength/2; i++ {
		if word[i] != word[wordLength-i-1] {
			return false
		}
	}

	return true
}
