package ingress

import (
	"cloudfunction/persist"
	"context"
	"log"
	"net/http"
	"os"
)

func FunctionOEM(w http.ResponseWriter, r *http.Request) {

}

func SubscribeOEM(ctx context.Context, b []byte) {
	log.Println("Starting full batch")
	w := make(chan persist.HasID, 100)
	persist.GetDataStoreWriter(ctx, os.Getenv("GCLOUD_PROJECT"), "makes", w)
	q := make(chan []byte, 100)
	persist.GetQueuePublisher(ctx, os.Getenv("GCLOUD_PROJECT"), "makes", q)

	//c := scrape.ScrapeOEMS()
	//for row := range c {
	//	w <- item{Title: row}
	//	q <- []byte(row)
	//}
	close(w)

}

type item struct {
	Title string
}

func (i item) GetID() string {
	return i.Title
}
