package io

import (
	"bytes"
	"cloudfunction/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var requestBody = []byte(`{"searchQueryModel":{"vehicleCategory":1, "makeModels":[], "sortOrder":0, "pageNumber":1}, "searchText":""}`)

func GetAllOEMs() []model.OEM {

	resp, err := http.Post(os.Getenv("MAKES_URL"), "Application/json", bytes.NewReader(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	r := model.Response{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		log.Fatal(err)
	}

	return r.OEMs
}
