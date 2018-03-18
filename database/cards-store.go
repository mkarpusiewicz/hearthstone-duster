package database

import (
	"encoding/json"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/mkarpusiewicz/hearthstone-duster/types"
)

// SaveMyCards - Save my cards info to database
func SaveMyCards(cards []types.MyCard) {
	if err := db.Update(func(tx *bolt.Tx) error {
		data, jErr := json.Marshal(cards)
		if jErr != nil {
			return jErr
		}

		b := tx.Bucket(bucketCards)
		pErr := b.Put(myCardsKey, data)
		if pErr != nil {
			return pErr
		}

		c := tx.Bucket(bucketConfig)
		timeBin, _ := time.Now().MarshalBinary()
		p2Err := c.Put(myCardsSyncTimeKey, timeBin)
		if p2Err != nil {
			return p2Err
		}

		return nil

	}); err != nil {
		log.Fatal(err)
	}
}

// GetMyCardsSyncTime - Get time of hearthpwn.com synchronization
func GetMyCardsSyncTime() time.Time {
	var syncTime time.Time
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketConfig)
		syncTime.UnmarshalBinary(b.Get(myCardsSyncTimeKey))
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return syncTime
}
