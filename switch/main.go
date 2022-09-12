package main

import "fmt"

func main() {
	food := "salad"

	switch food {
	case "salad":
		fmt.Println("boo!")
	case "pizza", "burger":
		fmt.Println("amazing!")
	default:
		fmt.Println("yum!")
	}
}
