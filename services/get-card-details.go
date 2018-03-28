package services

import (
	"encoding/json"
	"log"

	"github.com/mkarpusiewicz/hearthstone-duster/types"
)

const apiURL string = "https://api.hearthstonejson.com/v1/latest/enUS/cards.collectible.json"

// GetAPICards - Get current api database of cards from hearthstonejson.com
func GetAPICards() []types.ApiCard {
	var result []types.ApiCard

	respByte := readFromRemote(apiURL)
	//respByte := readFromLocal("_tools/cards_collectible.json")

	if err := json.Unmarshal(respByte, &result); err != nil {
		log.Fatal(err)
	}

	return result
}
