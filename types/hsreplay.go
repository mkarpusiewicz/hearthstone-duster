package types

type CardPlayedInfoAll struct {
	Id         int64   `json:"dbf_id"`
	Total      int64   `json:"total"`
	Popularity string  `json:"popularity"`
	WinRate    float64 `json:"winrate"`
}

type CardPlayedInfo struct {
	Id         int64   `json:"dbf_id"`
	Total      int64   `json:"total"`
	Popularity float64 `json:"popularity"`
	WinRate    float64 `json:"winrate"`
}

type CardIncludedInfoAll struct {
	Id         int64   `json:"dbf_id"`
	Decks      int64   `json:"decks"`
	Count      float64 `json:"count"`
	Popularity float64 `json:"popularity"`
	WinRate    float64 `json:"winrate"`
}

type CardIncludedInfo struct {
	Id         int64   `json:"dbf_id"`
	Decks      int64   `json:"decks"`
	Count      float64 `json:"count"`
	Popularity float64 `json:"popularity"`
	WinRate    float64 `json:"winrate"`
}

type CardUsageData struct {
	CardsPlayedDetails  CardPlayedResponse
	CardsInDecksDetails CardIncludedResponse
}

type CardPlayedResponse struct {
	AsOf     string `json:"as_of"`
	RenderAs string `json:"render_as"`
	Series   struct {
		Data struct {
			All     []CardPlayedInfoAll `json:"ALL"`
			Druid   []CardPlayedInfo    `json:"DRUID"`
			Hunter  []CardPlayedInfo    `json:"HUNTER"`
			Mage    []CardPlayedInfo    `json:"MAGE"`
			Paladin []CardPlayedInfo    `json:"PALADIN"`
			Priest  []CardPlayedInfo    `json:"PRIEST"`
			Rogue   []CardPlayedInfo    `json:"ROGUE"`
			Shaman  []CardPlayedInfo    `json:"SHAMAN"`
			Warlock []CardPlayedInfo    `json:"WARLOCK"`
			Warrior []CardPlayedInfo    `json:"WARRIOR"`
		} `json:"data"`
		Metadata struct {
			TotalPlayedCardsCount int64 `json:"total_played_cards_count"`
		} `json:"metadata"`
	} `json:"series"`
}

type CardIncludedResponse struct {
	AsOf     string `json:"as_of"`
	RenderAs string `json:"render_as"`
	Series   struct {
		Data struct {
			All     []CardIncludedInfoAll `json:"ALL"`
			Druid   []CardIncludedInfo    `json:"DRUID"`
			Hunter  []CardIncludedInfo    `json:"HUNTER"`
			Mage    []CardIncludedInfo    `json:"MAGE"`
			Paladin []CardIncludedInfo    `json:"PALADIN"`
			Priest  []CardIncludedInfo    `json:"PRIEST"`
			Rogue   []CardIncludedInfo    `json:"ROGUE"`
			Shaman  []CardIncludedInfo    `json:"SHAMAN"`
			Warlock []CardIncludedInfo    `json:"WARLOCK"`
			Warrior []CardIncludedInfo    `json:"WARRIOR"`
		} `json:"data"`
		Metadata struct {
			TotalPlayedDecksCount int64 `json:"total_played_decks_count"`
		} `json:"metadata"`
	} `json:"series"`
}
