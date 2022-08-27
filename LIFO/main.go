package main

import "fmt"

func main() {
	lifoStack := []int{}

	// Add elements
	addToStack(&lifoStack, 1)
	addToStack(&lifoStack, 2)
	addToStack(&lifoStack, 3)

	fmt.Printf("Elements added to stack, stack is now: %v\n", lifoStack)

	// Pop element
	poppedNumber := getNextInStack(&lifoStack)
	fmt.Printf("Number %v popped from LIFO stack, new stack: %v\n", poppedNumber, lifoStack)

	// Pop another element
	poppedNumber = getNextInStack(&lifoStack)
	fmt.Printf("Number %v popped from LIFO stack, new stack: %v\n", poppedNumber, lifoStack)
}

func addToStack(stack *[]int, number int) {
	*stack = append(*stack, number)
}

func getNextInStack(stack *[]int) int {
	finalIndex := len(*stack) - 1
	nextInStack := (*stack)[finalIndex]
	*stack = (*stack)[:finalIndex]

	return nextInStack
}
