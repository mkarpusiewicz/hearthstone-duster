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
			fmt.Printf("hearthpwn.com synced at: %s\r\n", mineSyncTime.Format(time.RFC1123))
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
