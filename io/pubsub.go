package io

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
	"os"
)

func GetQueuePublisher(topicName string, c chan []byte) {
	ctx := context.Background()
	proj := os.Getenv("PROJECT_ID")
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}

	t := client.Topic(topicName)
	for o := range c {
		t.Publish(ctx, &pubsub.Message{
			Data: o,
		})
	}

}
