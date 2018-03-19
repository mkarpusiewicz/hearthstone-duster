package database

import (
	"encoding/json"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func saveCards(cardData interface{}, cardsKey []byte, cardsSyncTimeKey []byte) {
	if err := db.Update(func(tx *bolt.Tx) error {
		data, jErr := json.Marshal(cardData)
		if jErr != nil {
			return jErr
		}

		b := tx.Bucket(bucketCards)
		pErr := b.Put(cardsKey, data)
		if pErr != nil {
			return pErr
		}

		c := tx.Bucket(bucketConfig)
		timeBin, _ := time.Now().MarshalBinary()
		p2Err := c.Put(cardsSyncTimeKey, timeBin)
		if p2Err != nil {
			return p2Err
		}

		return nil

	}); err != nil {
		log.Fatal(err)
	}
}

func getCards(cardData interface{}, cardsKey []byte) {
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketCards)
		err := json.Unmarshal(b.Get(cardsKey), cardData)
		return err
	}); err != nil {
		log.Fatal(err)
	}
}

func getCardsSyncTime(cardsSyncTimeKey []byte) time.Time {
	var syncTime time.Time
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketConfig)
		syncTime.UnmarshalBinary(b.Get(cardsSyncTimeKey))
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return syncTime
}
