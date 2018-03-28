package services

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func readFromRemote(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf.Bytes()
}

func readFromLocal(file string) []byte {
	respByte, err := ioutil.ReadFile("_tools/cards_collectible.json")
	if err != nil {
		log.Fatal(err)
	}
	return respByte
}
