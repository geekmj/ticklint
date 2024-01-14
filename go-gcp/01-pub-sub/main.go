package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"os"
)

func main() {
	err := os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "single-mix-317905")
	if err != nil {
		fmt.Errorf("pubsub: NewClient: %w", err)
	}
	defer client.Close()
	// Run when Emulator restarted
	// t, err := client.CreateTopic(ctx, "test-topic")
	t := client.Topic("test-topic")
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello World"),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Errorf("pubsub: result.Get: %w", err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)

}
