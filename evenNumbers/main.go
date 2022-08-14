package main

import "fmt"

func main() {
	printEvenNumbersUpTo(25)

}

func printEvenNumbersUpTo(max int) {
	for i := 2; i <= max; i += 2 {
		fmt.Println(i)
	}
}
