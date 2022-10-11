package main

import "fmt"

func main() {
	term := 5
	result := factorial(term)

	fmt.Printf("%v factorial is %v.\n", term, result)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
