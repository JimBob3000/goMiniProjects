package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	newPerson := Person{"James", "Elliott", 28}

	fmt.Println("First name: ", newPerson.firstName)
	fmt.Println("Last name: ", newPerson.lastName)
	fmt.Println("Age: ", newPerson.age)
}
