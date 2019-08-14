package persist

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
)

type HasID interface {
	GetID() string
}

func GetDataStoreWriter(ctx context.Context, projectID string, kind string, in chan string) {

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	go func() {
		defer client.Close()
		for row := range in {
			// Sets the name/ID for the new entity.
			k := datastore.NameKey(kind, row, nil)

			// Saves the new entity.
			if _, err := client.Put(ctx, k, &row); err != nil {
				log.Println("~~~ Persistence issue ~~~")
				log.Println("kind:", kind, "name:", row)
				log.Println("Failed to save task: %v", err)
			}
		}
	}()

}
