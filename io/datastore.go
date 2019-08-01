package io

import (
	"cloud.google.com/go/datastore"
	"cloudfunction/model"
	"context"
	"fmt"
	"log"
	"os"
)

func WriteOEM(o model.MakeModelResponseOEM) {
	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := os.Getenv("PROJECT_ID")

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the kind for the new entity.
	kind := "makes"
	// Sets the name/ID for the new entity.
	name := o.Title
	// Creates a Key instance.
	makeKey := datastore.NameKey(kind, name, nil)

	// Saves the new entity.
	if _, err := client.Put(ctx, makeKey, &o); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

	fmt.Println(fmt.Sprintf("Saved make %s", o.Title))
}
