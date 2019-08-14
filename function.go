package p

import (
	"cloudfunction/bus"
	"cloudfunction/domain"
	"context"
	"encoding/json"
	"net/http"
)

func ProcessAllOEMs(w http.ResponseWriter, r *http.Request) {
	bus.ProcessAllOEMs(context.Background(), true, true)
	req := struct {
		OK bool `json:"ok"`
	}{OK: true}
	b, _ := json.Marshal(req)
	w.Write(b)
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m domain.PubSubMessage) error {
	bus.ProcessAllOEMPages(ctx, string(m.Data))

	return nil
}

/*

	gcloud functions deploy all-oems --region europe-west2 --set-env-vars GCLOUD_PROJECT=wesbank-ta --entry-point ProcessAllOEMs --runtime go112 --trigger-http
	gcloud functions deploy all-oem-page-urls --region europe-west2 --entry-point HelloPubSub --runtime go112 --trigger-topic QueueOEMs

*/
