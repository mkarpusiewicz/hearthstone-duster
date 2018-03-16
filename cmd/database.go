package cmd

import (
	"os"
	"strconv"

	"github.com/mkarpusiewicz/hearthstone-duster/services"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// databaseCmd represents the database command
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Sync cards database from hearthstonejson.com",
	Run: func(cmd *cobra.Command, args []string) {
		apiCards := services.GetAPICards()

		cardsCount := 0
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Card", "Id", "Rarity", "Set", "Type"})
		table.SetAutoWrapText(false)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		for _, card := range apiCards {
			table.Append([]string{card.Name, strconv.FormatInt(card.DbfID, 10), card.Rarity, card.Set, card.Type})
			cardsCount++
		}

		table.SetFooter([]string{" ", " ", " ", "Total", strconv.Itoa(cardsCount)})

		table.Render()
	},
}

func init() {
	syncCmd.AddCommand(databaseCmd)
}
