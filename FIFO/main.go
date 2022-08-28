package main

import "fmt"

func main() {
	fifoQueue := []int{}

	// Add elements
	addToQueue(&fifoQueue, 1)
	addToQueue(&fifoQueue, 2)
	addToQueue(&fifoQueue, 3)

	fmt.Printf("Elements added to queue, queue is now: %v\n", fifoQueue)

	// Pop element
	poppedNumber := getNextInQueue(&fifoQueue)
	fmt.Printf("Number %v popped from FIFO queue, new queue: %v\n", poppedNumber, fifoQueue)

	// Pop another element
	poppedNumber = getNextInQueue(&fifoQueue)
	fmt.Printf("Number %v popped from FIFO stack, new stack: %v\n", poppedNumber, fifoQueue)
}

func addToQueue(queue *[]int, number int) {
	*queue = append(*queue, number)
}

func getNextInQueue(queue *[]int) int {
	nextInQueue := (*queue)[0]
	*queue = (*queue)[1:]

	return nextInQueue
}
