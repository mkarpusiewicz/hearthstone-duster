package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/mkarpusiewicz/hearthstone-duster/database"
	"github.com/mkarpusiewicz/hearthstone-duster/services"
	"github.com/mkarpusiewicz/hearthstone-duster/types"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var dustGoldenCmd = &cobra.Command{
	Use:   "golden",
	Short: "Show golden cards which can be dusted",
	Run: func(cmd *cobra.Command, args []string) {
		myCards := database.GetMyCards()
		apiCards := database.GetCardsDatabase()

		golden := getGoldenCards(myCards, apiCards)

		totalDustValue := 0

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Card", "Rarity", "Golden count", "Normal count", "Golds to dust", "Dust value"})
		table.SetAutoWrapText(false)
		table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER})
		for _, card := range golden {
			dustString := ""
			dustValueString := ""
			if card.ToDust > 0 {
				dustString = color.GreenString(strconv.Itoa(card.ToDust))

				dustValue := card.ToDust * services.GetDustValue(card.Rarity, true)
				totalDustValue += dustValue
				dustValueString = color.GreenString(strconv.Itoa(dustValue))
			}

			table.Append([]string{card.Name, services.GetRarityColoredString(card.Rarity, card.Rarity), color.YellowString(strconv.Itoa(card.GoldenCount)),
				strconv.Itoa(card.NormalCount), dustString, dustValueString})
		}
		table.SetFooter([]string{" ", " ", " ", " ", "Total", strconv.Itoa(totalDustValue)})
		table.Render()

		fmt.Printf("\r\n\r\nTo DUST for total %d:\r\n", totalDustValue)
		for _, card := range golden {
			if card.ToDust == 0 {
				continue
			}
			fmt.Printf("%d X %s\r\n", card.ToDust, card.Name)
		}
	},
}

func init() {
	dustCmd.AddCommand(dustGoldenCmd)
}

func getGoldenCards(myCards []types.MyCard, apiCards []types.ApiCard) []goldenCard {
	var goldenCards []goldenCard

	for _, card := range myCards {
		if !card.IsGold {
			continue
		}

		goldenCard := goldenCard{Name: card.Name, GoldenCount: card.Count}

		var cardFound bool
		for _, apiCard := range apiCards {
			if apiCard.Name == goldenCard.Name {
				cardFound = true
				goldenCard.Rarity = apiCard.Rarity
			}
		}
		if !cardFound {
			log.Println("WARNING: Card '%s' not found in card database", card.Name)
			continue
		}

		if goldenCard.Rarity == "FREE" {
			continue
		}

		goldenCards = append(goldenCards, goldenCard)
	}

	for i, goldenCard := range goldenCards {
		for _, card := range myCards {
			if card.Name == goldenCard.Name && card.IsGold == false {
				goldenCards[i].NormalCount = card.Count

				if goldenCards[i].NormalCount > 1 {
					goldenCards[i].ToDust = goldenCards[i].GoldenCount
				} else if goldenCards[i].NormalCount == 1 {
					goldenCards[i].ToDust = goldenCards[i].GoldenCount - 1
				}
			}
		}
	}

	return goldenCards
}

type goldenCard struct {
	Name        string
	GoldenCount int
	NormalCount int
	ToDust      int
	Rarity      string
}
