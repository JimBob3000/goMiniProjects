package main

import "fmt"

func main() {
	numbers := []int{6, 4, 17, 3, 2, 9, 13}
	cocktailSort(numbers)
	fmt.Println(numbers)
}

func cocktailSort(numbers []int) {
	numbersLength := len(numbers) - 1

	for i := 0; i < numbersLength; i++ {
		altered := false

		for j := 0; j < numbersLength; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				altered = true
			}
		}

		if !altered {
			break
		}

		for j := numbersLength; j > 0; j-- {
			if numbers[j] < numbers[j-1] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
				altered = true
			}
		}

		if !altered {
			break
		}
	}
}
