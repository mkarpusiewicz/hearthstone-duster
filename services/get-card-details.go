package services

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/mkarpusiewicz/hearthstone-duster/types"
)

const apiURL string = "https://api.hearthstonejson.com/v1/latest/enUS/cards.collectible.json"

// GetAPICards - Get current api database of cards from hearthstonejson.com
func GetAPICards() []types.ApiCard {
	var result []types.ApiCard

	//respByte := readFromRemote(apiURL)

	respByte, err := ioutil.ReadFile("_tools/cards_collectible.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(respByte, &result); err != nil {
		log.Fatal(err)
	}

	return result
}
