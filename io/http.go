package io

import (
	"bytes"
	"cloudfunction/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var requestBody = []byte(`{"searchQueryModel":{"vehicleCategory":1, "makeModels":[], "sortOrder":0, "pageNumber":1}, "searchText":""}`)

func GetAllOEMs() []model.MakeModelResponseOEM {
	start := time.Now()
	resp, err := http.Post(os.Getenv("MAKES_URL"), "Application/json", bytes.NewReader(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Took", time.Since(start), "to get makes")
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	r := model.MakeModelResponse{}
	if err = json.Unmarshal(b, &r); err != nil {
		log.Fatal(err)
	}

	return r.OEMs
}
