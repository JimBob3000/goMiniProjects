package main

import "fmt"

func main() {
	numbers := []int{6, 4, 17, 3, 2, 9, 13}
	equal := firstAndLastEqual(numbers)

	if equal {
		fmt.Println("First and last are equal")
	} else {
		fmt.Println("First and last are NOT equal")
	}
}

func firstAndLastEqual(numbers []int) bool {
	firstNumber := numbers[0]
	lastNumber := numbers[len(numbers)-1]

	if firstNumber == lastNumber {
		return true
	}

	return false
}
