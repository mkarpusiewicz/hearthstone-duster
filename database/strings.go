package database

var (
	bucketConfig       = []byte("config")
	hearthPwnUserKey   = []byte("hearthpwn-user")
	myCardsSyncTimeKey = []byte("my-cards-sync-time")

	bucketCards = []byte("cards")
	myCardsKey  = []byte("my-cards")
)

const (
	configurationDirectory string = "~/.hearthstone-duster"
	databaseFileName       string = "hs.db"
)
