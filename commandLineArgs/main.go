package main

import (
	"fmt"
	"os"
)

func main() {
	programName := os.Args[0]
	args := os.Args[1:]

	fmt.Printf("Program name: %v\n", programName)
	for index, arg := range args {
		fmt.Printf("Argument %v: %v\n", index, arg)
	}
}
