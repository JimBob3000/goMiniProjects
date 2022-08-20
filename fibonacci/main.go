package main

import "fmt"

func main() {
	fibonacci(55)
}

func fibonacci(max int) {
	a := 0
	b := 1
	c := 1

	for a <= max {
		fmt.Println(a)
		a, b = b, c
		c = a + b
	}
}
