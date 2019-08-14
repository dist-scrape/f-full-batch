package bus

import (
	"cloudfunction/domain"
	"cloudfunction/persist"
	"cloudfunction/scrape"
	"context"
	"log"
	"os"
)

var OEMPagesURL = domain.GetOEMPagesURL()

func ProcessAllOEMs(ctx context.Context, writeToDB, writeToQueue bool) {

	log.Println("All OEMs -> starting")

	//TODO: use fan-out pattern to ensure

	log.Println("All OEMs -> connecting to db")
	w := make(chan string, 100)
	persist.GetDataStoreWriter(ctx, os.Getenv("GCLOUD_PROJECT"), "makes", w)
	defer close(w)

	log.Println("All OEMs -> connecting to queue")
	q := make(chan []byte, 100)
	persist.GetQueuePublisher(ctx, os.Getenv("GCLOUD_PROJECT"), domain.QueueOEMs, q)
	defer close(q)

	log.Println("All OEMs -> getting all oems")
	c := scrape.GetAllOEMs(domain.GetOEMURL())
	for row := range c {
		if writeToDB {
			w <- string(row)
		}
		if writeToQueue {
			q <- []byte(row)
		}
		log.Println("All OEMs -> read...", row)
	}

	log.Println("All OEMs -> done")

}

func ProcessAllOEMPages(ctx context.Context, url string) {
	log.Println("All OEM pages -> starting")

	log.Println("All OEM pages -> connecting to queue")
	q := make(chan []byte, 100)
	persist.GetQueuePublisher(ctx, os.Getenv("GCLOUD_PROJECT"), domain.QueueOEMPageUrls, q)
	defer close(q)

	c := scrape.GetAllOEMPages(OEMPagesURL, domain.OEM(url))
	for row := range c {
		q <- []byte(string(row))
		log.Println("All OEM pages -> read...", row)
	}

	log.Println("All OEM pages -> done")

}
