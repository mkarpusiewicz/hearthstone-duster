package database

import (
	"time"

	"github.com/mkarpusiewicz/hearthstone-duster/types"
)

// SaveCardsUsage - Save cards usage to database
func SaveCardsUsage(cardsUsage types.CardUsageData) {
	saveCards(cardsUsage, cardsUsageKey, cardsUsageSyncTimeKey)
}

// GetCardsUsage - Get cards usage from database
func GetCardsUsage() types.CardUsageData {
	var cardsUsage types.CardUsageData
	getCards(&cardsUsage, cardsUsageKey)
	return cardsUsage
}

// GetCardsUsageSyncTime - Get time of hsreplay.net synchronization
func GetCardsUsageSyncTime() time.Time {
	return getCardsSyncTime(cardsUsageSyncTimeKey)
}
