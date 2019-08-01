// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"cloudfunction/io"
	"cloudfunction/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// Router holds a list of functions that allow the same codebase to serve multiple cloud functions
var router = map[string]func(b []byte){
	"f-full-batch": fFullBatch,
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m model.PubSubMessage) error {
	checkConfig()

	log.Println(string(m.Data))
	router[os.Getenv("FUNCTION")](m.Data)
	return nil
}

func fFullBatch(b []byte) {
	//Write makes to DB
	c := make(chan model.MakeModelResponseOEM)
	go io.GetOEMDataStoreWriter(c)
	oems := io.GetAllOEMs()
	start := time.Now()
	for _, oem := range oems {
		c <- oem
	}
	log.Println("Took", time.Since(start), "to write makes to datastore")
	close(c)

	//Write items to the queue
	ch := make(chan []byte)
	go io.GetQueuePublisher(model.TopicOEM, ch)
	start = time.Now()
	for _, oem := range oems {
		arr := model.CreateMessageSubModelFromMakeModelResponseOEM(oem)
		for _, msgSubModel := range arr {
			b, err := json.Marshal(msgSubModel)
			if err != nil {
				log.Println(err)
			}
			ch <- b
		}
	}
	log.Println("Took", time.Since(start), "to write makes to queue")
	close(ch)

}

func checkConfig() {
	proj := os.Getenv("PROJECT_ID")
	if proj == "" {
		fmt.Fprintf(os.Stderr, "PROJECT_ID environment variable must be set.\n")
		os.Exit(1)
	}
}

/*
	gcloud functions deploy f-full-batch --region europe-west2 --entry-point HelloPubSub --runtime go112 --trigger-topic full-batch

*/
