package main

import "fmt"

func main() {
	sentence := "Hello World"
	mostCommonLetter(sentence)
}

func mostCommonLetter(sentence string) {
	letters := make(map[rune]int)

	for _, letter := range sentence {
		if letters[letter] > 0 {
			letters[letter]++
		} else {
			letters[letter] = 1
		}
	}

	biggestValue := 0
	biggestLetter := ""

	for letter, value := range letters {
		if value > biggestValue {
			biggestValue = value
			biggestLetter = string(letter)
		}
	}

	fmt.Printf("The most common letter is %v\n", biggestLetter)
}
