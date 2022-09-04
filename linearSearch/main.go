package main

import "fmt"

func main() {
	numbers := []int{6, 14, 11, 39, 5, 7, 28, 29, 40}
	targetNumber := 28

	found, guessCount := linearSearch(numbers, targetNumber)

	if found {
		fmt.Printf("Target number %v found in %v guesses.\n", targetNumber, guessCount)
	} else {
		fmt.Printf("Target number %v not in slice.\n", targetNumber)
	}
}

func linearSearch(numbers []int, targetNumber int) (bool, int) {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == targetNumber {
			return true, i
		}
	}

	return false, len(numbers)
}
