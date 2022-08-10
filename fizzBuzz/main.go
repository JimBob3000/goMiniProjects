package main

import "fmt"

func main() {
	fizzBuzz(5, -30)
}

func fizzBuzz(start, end int) {
	if start > end {
		start, end = end, start
	}

	for i := start; i <= end; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
