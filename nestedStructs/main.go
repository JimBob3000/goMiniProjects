package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	pets      []Animal
}

type Animal struct {
	name    string
	species string
}

func main() {
	firstAnimal := Animal{"Sam", "Cat"}
	secondAnimal := Animal{"Willis", "Dog"}
	newPerson := Person{
		"James",
		"Elliott",
		28,
		[]Animal{
			firstAnimal,
			secondAnimal,
		},
	}

	fmt.Println("First name: ", newPerson.firstName)
	fmt.Println("Last name: ", newPerson.lastName)
	fmt.Println("Age: ", newPerson.age)

	for i, pet := range newPerson.pets {
		fmt.Println("Pet ", i+1)
		fmt.Println("Name: ", pet.name)
		fmt.Println("Species: ", pet.species)
	}
}
