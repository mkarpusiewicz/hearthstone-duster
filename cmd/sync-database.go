package cmd

import (
	"fmt"

	"github.com/mkarpusiewicz/hearthstone-duster/database"
	"github.com/mkarpusiewicz/hearthstone-duster/services"
	"github.com/spf13/cobra"
)

var syncDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Sync cards database from hearthstonejson.com",
	Run: func(cmd *cobra.Command, args []string) {
		apiCards := services.GetAPICards()

		database.SaveCardsDatabase(apiCards)

		cardsCount := len(apiCards)

		fmt.Printf("Synced %d cards\r\n", cardsCount)
	},
}

func init() {
	syncCmd.AddCommand(syncDatabaseCmd)
}
