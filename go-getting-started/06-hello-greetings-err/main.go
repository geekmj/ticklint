package main

import (
	"fmt"
	"log"
	"ticklint/greetingserr"
)

func main() {
	log.SetPrefix("greetingsErr: ")
	log.SetFlags(0)
	// Get a greeting message and print it.
	names := []string{"Ram", "Shyam", "John"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
