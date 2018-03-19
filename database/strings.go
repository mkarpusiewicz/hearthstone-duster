package database

var (
	bucketConfig     = []byte("config")
	hearthPwnUserKey = []byte("hearthpwn-user")

	bucketCards = []byte("cards")

	myCardsKey         = []byte("my-cards")
	myCardsSyncTimeKey = []byte("my-cards-sync-time")

	cardsUsageKey         = []byte("cards-usage")
	cardsUsageSyncTimeKey = []byte("cards-usage-sync-time")
)

const (
	configurationDirectory string = "~/.hearthstone-duster"
	databaseFileName       string = "hs.db"
)
