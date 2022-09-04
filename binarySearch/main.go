package main

import (
	"fmt"
	"math"
)

func main() {
	targetNumber := 45
	numbers := []int{}
	for i := 0; i < 100; i++ {
		numbers = append(numbers, i)
	}

	found, guessCount := binarySearch(numbers, targetNumber)

	if found {
		fmt.Printf("Target number %v found in %v guessses.\n", targetNumber, guessCount)
	} else {
		fmt.Printf("Target number %v not in slice, searched %v times.\n", targetNumber, guessCount)
	}
}

func binarySearch(numbers []int, targetNumber int) (bool, int) {
	min := 0
	max := len(numbers)
	guess := math.Floor(float64((min + max) / 2))
	previousGuess := guess
	guessCount := 1

	for int(guess) != targetNumber {
		if targetNumber > int(guess) {
			min = int(guess)
		} else {
			max = int(guess)
		}

		previousGuess = guess
		guess = math.Floor(float64((min + max) / 2))

		if guess == previousGuess {
			return false, guessCount
		}

		guessCount++
	}

	return true, guessCount
}
