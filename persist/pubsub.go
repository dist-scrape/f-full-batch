package persist

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

func GetQueuePublisher(ctx context.Context, project, topicName string, c chan []byte) {
	client, err := pubsub.NewClient(ctx, project)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	t := client.Topic(topicName)
	go func() {
		for o := range c {
			t.Publish(ctx, &pubsub.Message{
				Data: o,
			})
		}
	}()

}
