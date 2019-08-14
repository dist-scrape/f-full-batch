package persist

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
)

type HasID interface {
	GetID() string
}

func GetDataStoreWriter(ctx context.Context, projectID string, kind string, in chan HasID) {

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	for row := range in {
		// Sets the name/ID for the new entity.
		name := row.GetID()
		k := datastore.NameKey(kind, name, nil)

		// Saves the new entity.
		if _, err := client.Put(ctx, k, &row); err != nil {
			log.Fatalf("Failed to save task: %v", err)
		}

	}
}
