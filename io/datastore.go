package io

import (
	"cloud.google.com/go/datastore"
	"cloudfunction/model"
	"context"
	"fmt"
	"log"
	"os"
)

func GetOEMDataStoreWriter(c chan model.MakeModelResponseOEM) {
	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := os.Getenv("GCLOUD_PROJECT")

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the kind for the new entity.
	kind := "makes"
	// Sets the name/ID for the new entity.

	count := 0
	for o := range c {
		// Creates a Key instance.
		name := o.Title
		makeKey := datastore.NameKey(kind, name, nil)
		// Saves the new entity.
		if _, err := client.Put(ctx, makeKey, &o); err != nil {
			log.Fatalf("Failed to save task: %v", err)
		}
		count++
	}
	fmt.Println(fmt.Sprintf("Saved makes: %i", count))
}
