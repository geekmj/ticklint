package main

import (
	"fmt"
	"ticklint/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Mrityunjay")
	fmt.Println(message)
}
