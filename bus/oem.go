package bus

import (
	"cloudfunction/domain"
	"cloudfunction/persist"
	"cloudfunction/scrape"
	"context"
	"log"
	"os"
)

func ProcessAllOEMs(ctx context.Context, writeToDB, writeToQueue bool) {

	log.Println("Starting full batch")

	//TODO: use fan-out pattern to ensure

	w := make(chan persist.HasID, 100)
	persist.GetDataStoreWriter(ctx, os.Getenv("GCLOUD_PROJECT"), "makes", w)

	q := make(chan []byte, 100)
	persist.GetQueuePublisher(ctx, os.Getenv("GCLOUD_PROJECT"), "makes", q)

	c := scrape.GetAllOEMs(domain.GetOEMURL())
	for row := range c {
		if writeToDB {
			w <- item{Title: string(row)}
		}
		if writeToQueue {
			q <- []byte(row)
		}
	}
	close(w)
	close(q)

}

type item struct {
	Title string
}

func (i item) GetID() string {
	return i.Title
}
