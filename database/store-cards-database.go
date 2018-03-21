package database

import (
	"time"

	"github.com/mkarpusiewicz/hearthstone-duster/types"
)

// SaveCardsDatabase - Save cards database to database
func SaveCardsDatabase(cardsDatabase []types.ApiCard) {
	saveCards(cardsDatabase, cardsDatabaseKey, cardsDatabaseyncTimeKey)
}

// GetCardsDatabase - Get cards database from database
func GetCardsDatabase() []types.ApiCard {
	var cardsDatabase []types.ApiCard
	getCards(&cardsDatabase, cardsDatabaseKey)
	return cardsDatabase
}

// GetCardsDatabaseSyncTime - Get time of hearthstonejson.com synchronization
func GetCardsDatabaseSyncTime() time.Time {
	return getCardsSyncTime(cardsDatabaseyncTimeKey)
}
