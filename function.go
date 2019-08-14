package p

import (
	"cloudfunction/bus"
	"cloudfunction/domain"
	"context"
	"log"
	"net/http"
	"os"
)

var router = map[string]func(b []byte){
	"f-full-batch": fFullBatch,
}

func ProcessAllOEMs(w http.ResponseWriter, r *http.Request) {
	bus.ProcessAllOEMs(context.Background(), false, false)
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

	gcloud functions deploy all-oems --region europe-west2 --set-env-vars GCLOUD_PROJECT=wesbank-ta --entry-point ProcessAllOEMs --runtime go112 --trigger-http
	gcloud functions deploy f-full-batch --region europe-west2 --entry-point HelloPubSub --runtime go112 --trigger-topic full-batch

*/
