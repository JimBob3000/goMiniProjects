package main

import "fmt"

func main() {
	targetNumber := 32
	numbers := []int{1, 1, 1, 2, 5, 7, 8, 11, 11, 15, 19, 27, 29, 30, 31, 32, 33, 38, 40, 45, 46, 50}
	jumpSearch(numbers, targetNumber)
}

func jumpSearch(numbers []int, targetNumber int) {
	guessCount := 0
	currentIndex := 0

	for i := 0; i < len(numbers)-1; i += 4 {
		fmt.Printf("Jumping to index %v and guessing %v...\n", i, numbers[i])

		if numbers[i] >= targetNumber {
			fmt.Printf("Jump too large, reverting to index %v and moving to incremental guesses\n", currentIndex)
			break
		}

		currentIndex = i
		guessCount++
	}

	for currentIndex < len(numbers)-1 {
		fmt.Printf("Incrementing to index %v and guessing %v...\n", currentIndex, numbers[currentIndex])

		if numbers[currentIndex] == targetNumber {
			fmt.Printf("Target number %v found in %v guesses\n", numbers[currentIndex], guessCount)
			break
		}

		currentIndex++
		guessCount++
	}
}
