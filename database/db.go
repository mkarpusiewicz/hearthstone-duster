package database

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/boltdb/bolt"
	homedir "github.com/mitchellh/go-homedir"
)

var (
	db *bolt.DB
)

// GetDirectory - Get directory location with database and other files
func GetDirectory() string {
	dbPath, err := homedir.Expand(configurationDirectory)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		os.Mkdir(dbPath, 0600)
	}

	return dbPath
}

// Initialize database
func Initialize() {
	var err error

	dbFile := path.Join(GetDirectory(), databaseFileName)
	if db, err = bolt.Open(dbFile, 0600, nil); err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		for _, bucket := range [][]byte{bucketConfig, bucketCards} {
			if _, bErr := tx.CreateBucketIfNotExists(bucket); err != nil {
				return fmt.Errorf("create bucket of '%s' failed: %s", bucket, bErr)
			}
		}

		return nil
	})
}

// Close database
func Close() {
	db.Close()
}

// GetHearthPwnUser - Get hearthpwn.com user name for scraping data
func GetHearthPwnUser() string {
	var user string
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketConfig)
		user = string(b.Get(hearthPwnUserKey))
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return user
}

// SetHearthPwnUser - Set hearthpwn.com user name for scraping data
func SetHearthPwnUser(user string) {
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketConfig)
		pErr := b.Put(hearthPwnUserKey, []byte(user))
		return pErr
	}); err != nil {
		log.Fatal(err)
	}
}
