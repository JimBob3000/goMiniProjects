package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	text := "super secret"

	hash := sha256.New()
	hash.Write([]byte(text))
	hashedText := hash.Sum(nil)

	fmt.Println("Plain text: ", text)
	fmt.Printf("Hashed text: %x\n", hashedText)
}
