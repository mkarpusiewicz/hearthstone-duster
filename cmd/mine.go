package cmd

import (
	"fmt"
	"log"

	"github.com/mkarpusiewicz/hearthstone-duster/database"
	"github.com/mkarpusiewicz/hearthstone-duster/services"
	"github.com/spf13/cobra"
)

// mineCmd represents the mine command
var mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Synchronize mine cards from hearthpwn.com",
	Run: func(cmd *cobra.Command, args []string) {
		user := database.GetHearthPwnUser()

		if user == "" {
			log.Fatal("No user specified, please use 'user' command to set user name first")
		}

		myCards := services.GetUserCards(user)

		for _, card := range myCards {
			fmt.Printf("%s: %d - %t\n", card.Name, card.Count, card.IsGold)
		}
	},
}

func init() {
	syncCmd.AddCommand(mineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
