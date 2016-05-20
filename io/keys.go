package io

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Keys struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
}

func readKeys() Keys {
	var keys Keys
	text, err := ioutil.ReadFile("../keys.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(text, &keys)
	return keys
}
