package cmd

import (
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/mkarpusiewicz/hearthstone-duster/database"
	"github.com/mkarpusiewicz/hearthstone-duster/types"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var dustGoldenCmd = &cobra.Command{
	Use:   "golden",
	Short: "Show golden cards which can be dusted",
	Run: func(cmd *cobra.Command, args []string) {
		myCards := database.GetMyCards()

		golden := getGoldenCards(myCards)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Card", "Golden count", "Normal count", "Golds to dust"})
		table.SetAutoWrapText(false)
		table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER})
		for _, card := range golden {
			dustString := ""
			if card.ToDust > 0 {
				dustString = color.GreenString(strconv.Itoa(card.ToDust))
			}

			table.Append([]string{color.YellowString(card.Name), color.YellowString(strconv.Itoa(card.GoldenCount)),
				strconv.Itoa(card.NormalCount), dustString})
		}

		table.Render()
	},
}

func init() {
	dustCmd.AddCommand(dustGoldenCmd)
}

func getGoldenCards(myCards []types.MyCard) []goldenCard {
	var goldenCards []goldenCard

	for _, card := range myCards {
		if card.IsGold {
			goldenCards = append(goldenCards, goldenCard{Name: card.Name, GoldenCount: card.Count})
		}
	}

	// todo: check in database which cards are not basic ones and add dust value from rarity

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
}
