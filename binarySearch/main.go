package main

import (
	"fmt"
	"math"
)

func main() {
	targetNumber := 77
	numbers := []int{}
	for i := 0; i < 100; i++ {
		numbers = append(numbers, i)
	}

	binarySearch(numbers, targetNumber)
}

func binarySearch(numbers []int, targetNumber int) {
	min := 0
	max := len(numbers)
	guessCount := 0
	guess := math.Floor(float64((min + max) / 2))

	for int(guess) != targetNumber {
		fmt.Printf("Guessing %v...\n", int(guess))

		if targetNumber > int(guess) {
			min = int(guess)
		} else {
			max = int(guess)
		}

		guess = math.Floor(float64((min + max) / 2))
		guessCount++
	}

	fmt.Printf("Target number %v found in %v guesses\n", guess, guessCount)
}
