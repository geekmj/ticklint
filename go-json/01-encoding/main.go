package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Name string
	Body string
	Time time.Time
}

func main() {

	fmt.Printf("Hi JSON Encoding Stuffs")
	m := Message{"Alice", "Hello", time.Now()}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error in encoding")
	}

	fmt.Printf("\nSuccessfully converted %s", b)
}
