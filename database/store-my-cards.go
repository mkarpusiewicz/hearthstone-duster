package database

import (
	"time"

	"github.com/mkarpusiewicz/hearthstone-duster/types"
)

// SaveMyCards - Save my cards info to database
func SaveMyCards(cards []types.MyCard) {
	saveCards(cards, myCardsKey, myCardsSyncTimeKey)
}

// GetMyCards - Get my cards info from database
func GetMyCards() []types.MyCard {
	var myCards []types.MyCard
	getCards(&myCards, myCardsKey)
	return myCards
}

// GetMyCardsSyncTime - Get time of hearthpwn.com synchronization
func GetMyCardsSyncTime() time.Time {
	return getCardsSyncTime(myCardsSyncTimeKey)
}
