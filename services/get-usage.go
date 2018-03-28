package services

import (
	"encoding/json"
	"log"

	"github.com/mkarpusiewicz/hearthstone-duster/types"
)

const cardsPlayedURL string = "https://hsreplay.net/analytics/query/card_played_popularity_report/?GameType=RANKED_STANDARD&RankRange=ALL&TimeRange=LAST_14_DAYS"
const cardsIncludedURL string = "https://hsreplay.net/analytics/query/card_included_popularity_report/?GameType=RANKED_STANDARD&RankRange=ALL&TimeRange=LAST_14_DAYS"

// GetUsageInformation - Get card usages from hsreplay.net
func GetUsageInformation() types.CardUsageData {
	var result types.CardUsageData

	result.CardsPlayedDetails = getPlayedCards()
	result.CardsInDecksDetails = getIncludedCards()

	return result
}

func getPlayedCards() types.CardPlayedResponse {
	var result types.CardPlayedResponse

	respByte := readFromRemote(cardsPlayedURL)
	//respByte := readFromLocal("_tools/card_played.json")

	if err := json.Unmarshal(respByte, &result); err != nil {
		log.Fatal(err)
	}

	return result
}

func getIncludedCards() types.CardIncludedResponse {
	var result types.CardIncludedResponse

	respByte := readFromRemote(cardsIncludedURL)
	//respByte := readFromLocal("_tools/card_included.json")

	if err := json.Unmarshal(respByte, &result); err != nil {
		log.Fatal(err)
	}

	return result
}
