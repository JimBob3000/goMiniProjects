package main

import "fmt"

func main() {
	numberOfRows := 5

	for i := 0; i < numberOfRows; i++ {
		for j := 0; j <= numberOfRows-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k <= i; k++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}
