package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mkarpusiewicz/hearthstone-duster/database"
	"github.com/mkarpusiewicz/hearthstone-duster/services"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var showDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Show cards database from hearthstonejson.com",
	Run: func(cmd *cobra.Command, args []string) {
		syncTime := database.GetCardsDatabaseSyncTime()
		if syncTime.IsZero() {
			fmt.Print("hearthstonejson.com was not synced yet\r\n")
			return
		}
		fmt.Printf("hearthstonejson.com synced at: %s\r\n\r\n", syncTime.Format(time.RFC1123))

		apiCards := database.GetCardsDatabase()

		cardsCount := 0
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Card", "Id", "Rarity", "Set", "Type"})
		table.SetAutoWrapText(false)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		for _, card := range apiCards {
			rarityString := services.GetRarityColoredString(card.Rarity, card.Rarity)
			table.Append([]string{card.Name, strconv.FormatInt(card.DbfID, 10), rarityString, card.Set, card.Type})
			cardsCount++
		}

		table.SetFooter([]string{" ", " ", " ", "Total", strconv.Itoa(cardsCount)})

		table.Render()
	},
}

func init() {
	showCmd.AddCommand(showDatabaseCmd)
}
