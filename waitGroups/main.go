package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("App started")

	wg.Add(1)
	go welcome("James")

	fmt.Println("App finished - goroutine may finish after main func!")

	wg.Wait()
}

func welcome(name string) {
	fmt.Println("Welcome,", name)
	wg.Done()
}
