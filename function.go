// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"cloudfunction/io"
	"cloudfunction/model"
	"context"
	"log"
	"os"
)

// Router holds a list of functions that allow the same codebase to serve multiple cloud functions
var router = map[string]func(b []byte){
	"f-full-batch": fFullBatch,
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m model.PubSubMessage) error {
	log.Println(string(m.Data))
	router[os.Getenv("FUNCTION")](m.Data)
	return nil
}

func fFullBatch(b []byte) {
	c := make(chan model.MakeModelResponseOEM)
	go io.WriteOEM(c)
	oems := io.GetAllOEMs()
	for _, oem := range oems {
		c <- oem
	}
	close(c)

}

/*
	gcloud functions deploy f-full-batch --region europe-west2 --entry-point HelloPubSub --runtime go112 --trigger-topic full-batch

*/
