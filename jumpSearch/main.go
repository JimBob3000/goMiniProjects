package main

import "fmt"

func main() {
	numbers := []int{1, 1, 1, 2, 5, 7, 8, 11, 11, 15, 19, 27, 29, 30, 31, 32, 33, 38, 40, 45, 46, 50}
	targetNumber := 33

	found, guessCount := jumpSearch(numbers, targetNumber)

	if found {
		fmt.Printf("Target number %v found in %v guessses.\n", targetNumber, guessCount)
	} else {
		fmt.Printf("Target number %v not in slice, searched %v times.\n", targetNumber, guessCount)
	}
}

func jumpSearch(numbers []int, targetNumber int) (bool, int) {
	guessCount := 1
	currentIndex := 0
	numberLength := len(numbers)
	jump := 4

	for {
		if currentIndex > numberLength {
			currentIndex = currentIndex - jump
			break
		}

		if numbers[currentIndex] > targetNumber {
			currentIndex = currentIndex - jump
			break
		}

		currentIndex = currentIndex + jump
		guessCount++
	}

	for currentIndex < numberLength {
		if numbers[currentIndex] == targetNumber {
			return true, guessCount
		}

		currentIndex++
		guessCount++
	}

	return false, guessCount
}
