package main

import "fmt"

type Animal struct {
	species string
	noise   string
	speed   int
}

func (a Animal) speak() {
	fmt.Println(a.noise)
}

func (a Animal) speedClassification() string {
	if a.speed > 2 {
		return "fast"
	}
	return "slow"
}

func (a Animal) distanceMoved(duration int) int {
	return duration * a.speed
}

func main() {
	exampleAnimal := Animal{
		"Dog",
		"Woof",
		5,
	}

	exampleAnimal.speak()

	speedClassification := exampleAnimal.speedClassification()
	fmt.Printf("Animal %v is %v\n", exampleAnimal.species, speedClassification)

	duration := 10
	distance := exampleAnimal.distanceMoved(duration)
	fmt.Printf("Animal %v moves %v meters over %v seconds\n", exampleAnimal.species, distance, duration)
}
