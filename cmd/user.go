package cmd

import (
	"fmt"

	"github.com/mkarpusiewicz/hearthstone-duster/database"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Set hearthpwn.com user name for data scraping",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		user := args[0]
		database.SetHearthPwnUser(user)
		fmt.Printf("hearthpwn.com user set to: %s", user)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
