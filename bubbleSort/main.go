package main

import "fmt"

func main() {
	numbers := []int{6, 4, 17, 3, 2, 9, 13}
	bubbleSort(numbers)
	fmt.Println(numbers)
}

func bubbleSort(numbers []int) {
	numbersLength := len(numbers) - 1

	for i := 0; i < numbersLength; i++ {
		for j := 0; j < numbersLength; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}
