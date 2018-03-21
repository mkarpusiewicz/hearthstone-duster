package cmd

import (
	"fmt"
	"time"

	"github.com/mkarpusiewicz/hearthstone-duster/database"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show info about config and sync data",
	Run: func(cmd *cobra.Command, args []string) {
		dir := database.GetDirectory()
		user := database.GetHearthPwnUser()

		fmt.Printf("Configuration location: %s\r\n", dir)
		if user == "" {
			fmt.Print("hearthpwn.com user is not set\r\n")
		} else {
			fmt.Printf("hearthpwn.com user: %s\r\n", user)
		}
		fmt.Println()

		mineSyncTime := database.GetMyCardsSyncTime()
		if mineSyncTime.IsZero() {
			fmt.Print("hearthpwn.com was not synced yet\r\n")
		} else {
			fmt.Printf("hearthpwn.com synced at: \t%s\r\n", mineSyncTime.Format(time.RFC1123))
		}

		cardsUsageSyncTime := database.GetCardsUsageSyncTime()
		if cardsUsageSyncTime.IsZero() {
			fmt.Print("hsreplay.net was not synced yet\r\n")
		} else {
			fmt.Printf("hsreplay.net synced at: \t%s\r\n", cardsUsageSyncTime.Format(time.RFC1123))
		}

		cardsDatabaseSyncTime := database.GetCardsDatabaseSyncTime()
		if cardsDatabaseSyncTime.IsZero() {
			fmt.Print("hearthstonejson.com was not synced yet\r\n")
		} else {
			fmt.Printf("hearthstonejson.com synced at: \t%s\r\n", cardsDatabaseSyncTime.Format(time.RFC1123))
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
