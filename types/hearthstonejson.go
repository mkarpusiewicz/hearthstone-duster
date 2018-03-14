package types

type ApiCard struct {
	Armor              int64    `json:"armor"`
	Artist             string   `json:"artist"`
	Attack             int64    `json:"attack"`
	CardClass          string   `json:"cardClass"`
	Classes            []string `json:"classes"`
	Collectible        bool     `json:"collectible"`
	CollectionText     string   `json:"collectionText"`
	Cost               int64    `json:"cost"`
	DbfID              int64    `json:"dbfId"`
	Durability         int64    `json:"durability"`
	Elite              bool     `json:"elite"`
	Entourage          []string `json:"entourage"`
	Faction            string   `json:"faction"`
	Flavor             string   `json:"flavor"`
	Health             int64    `json:"health"`
	HideStats          bool     `json:"hideStats"`
	HowToEarn          string   `json:"howToEarn"`
	HowToEarnGolden    string   `json:"howToEarnGolden"`
	ID                 string   `json:"id"`
	Mechanics          []string `json:"mechanics"`
	MultiClassGroup    string   `json:"multiClassGroup"`
	Name               string   `json:"name"`
	Overload           int64    `json:"overload"`
	QuestReward        string   `json:"questReward"`
	Race               string   `json:"race"`
	Rarity             string   `json:"rarity"`
	ReferencedTags     []string `json:"referencedTags"`
	Set                string   `json:"set"`
	SpellDamage        int64    `json:"spellDamage"`
	TargetingArrowText string   `json:"targetingArrowText"`
	Text               string   `json:"text"`
	Type               string   `json:"type"`
}
