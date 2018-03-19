package cmd

import (
	"fmt"

	"github.com/mkarpusiewicz/hearthstone-duster/database"
	"github.com/mkarpusiewicz/hearthstone-duster/services"
	"github.com/spf13/cobra"
)

var syncUsageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Sync cards usage from hsreplay.net",
	Run: func(cmd *cobra.Command, args []string) {
		cardsUsage := services.GetUsageInformation()

		database.SaveCardsUsage(cardsUsage)

		fmt.Printf("Synced cards usage\r\n")
	},
}

func init() {
	syncCmd.AddCommand(syncUsageCmd)
}
