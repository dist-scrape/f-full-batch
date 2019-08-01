// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var requestBody = []byte(`{"searchQueryModel":{"vehicleCategory":1, "makeModels":[], "sortOrder":0, "pageNumber":1}, "searchText":""}`)

type Response struct {
	MakeModels []Make `json:"makeModels"`
}

type Make struct {
	Make     string  `json:"make"`
	Children []Model `json:"children"`
}

type Model struct {
	Model string `json:"model"`
}

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	log.Println(string(m.Data))
	GetAllMakes()
	return nil
}

func GetAllMakes() (makes []string) {
	makes = make([]string, 0)

	resp, err := http.Post(os.Getenv("MAKES_URL"), "Application/json", bytes.NewReader(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	restResponse := Response{}
	err = json.Unmarshal(b, &restResponse)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range restResponse.MakeModels {
		for _, model := range m.Children {
			log.Println(m.Make, model.Model)
		}
	}

	return
}

/*
	gcloud functions deploy f-full-batch --region europe-west2 --entry-point HelloPubSub --runtime go112 --trigger-topic full-batch

*/
