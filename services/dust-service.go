package services

import "github.com/fatih/color"

var (
	dustMap        = map[string]int{"COMMON": 5, "RARE": 20, "EPIC": 100, "LEGENDARY": 400}
	goldenDustMap  = map[string]int{"COMMON": 50, "RARE": 100, "EPIC": 400, "LEGENDARY": 1600}
	rarityColorMap = map[string]func(string, ...interface{}) string{
		"FREE": color.WhiteString, "COMMON": color.HiWhiteString, "RARE": color.BlueString, "EPIC": color.MagentaString, "LEGENDARY": color.YellowString}
)

// GetDustValue - get dust value based on card rarity and if it's golden
func GetDustValue(rarity string, isGolden bool) int {
	if isGolden {
		return goldenDustMap[rarity]
	}

	return dustMap[rarity]
}

// GetRarityColoredString - get colored string based on card rarity
func GetRarityColoredString(data string, rarity string) string {
	return rarityColorMap[rarity](data)
}
