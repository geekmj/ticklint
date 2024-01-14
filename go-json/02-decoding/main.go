package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Name string
	Body Signal
	Time time.Time
}

type Signal struct {
	Type string
	Info string
}

func main() {
	fmt.Printf("Json Decoding Stuff")

	b := []byte("[{\"Name\":\"Alice\",\"Body\":{\"Type\": \"GEN\", \"Info\": \"HURRY\"},\"Time\":\"2024-01-10T11:50:12Z\"}]\n")

	var m []Message
	err := json.Unmarshal(b, &m)

	if err != nil {
		fmt.Printf("Error in JSON Decoding %s", err.Error())
	}

	fmt.Printf("JSON String converted to Message Struct Name:%s Body:%s Time:%s", m[0].Name, m[0].Body, m[0].Time)

}
