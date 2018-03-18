package cmd

import (
	"fmt"
	"log"

	"github.com/mkarpusiewicz/hearthstone-duster/database"
	"github.com/mkarpusiewicz/hearthstone-duster/services"
	"github.com/spf13/cobra"
)

var syncMineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Synchronize mine cards from hearthpwn.com",
	Run: func(cmd *cobra.Command, args []string) {
		user := database.GetHearthPwnUser()

		if user == "" {
			log.Fatal("No user specified, please use 'user' command to set user name first")
		}

		myCards := services.GetUserCards(user)
		database.SaveMyCards(myCards)

		cardsCount := 0
		goldenCount := 0

		for _, card := range myCards {
			if card.IsGold {
				goldenCount += card.Count
			} else {
				cardsCount += card.Count
			}
		}

		fmt.Printf("Synced %d (including %d golden) cards for user %s\r\n", cardsCount+goldenCount, goldenCount, user)
	},
}

func init() {
	syncCmd.AddCommand(syncMineCmd)
}
