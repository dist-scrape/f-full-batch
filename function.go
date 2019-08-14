// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"cloudfunction/domain"
	"context"
	"log"
	"os"
)

var router = map[string]func(b []byte){
	"f-full-batch": fFullBatch,
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m domain.PubSubMessage) error {
	log.Println(string(m.Data))
	router[os.Getenv("FUNCTION")](m.Data)
	return nil
}

func fFullBatch(b []byte) {
	//oems := persist.()
	//for _, oem := range oems {
	//	persist.WriteOEM(oem)
	//}

}

/*
	gcloud functions deploy f-full-batch --region europe-west2 --entry-point HelloPubSub --runtime go112 --trigger-topic full-batch

*/
