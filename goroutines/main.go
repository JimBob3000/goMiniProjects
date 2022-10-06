package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("App started")

	go welcome("James")

	// Without a sleep, the main func finshes execution before the goroutine.
	// This means the app closes and we would not see the goroutine output.
	// We need to use wait groups to keep concurrency, not sleeps!
	time.Sleep(100 * time.Millisecond)

	fmt.Println("App finished - goroutine may not have finished!")
}

func welcome(name string) {
	fmt.Println("Welcome,", name)
}
