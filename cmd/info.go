package cmd

import (
	"fmt"

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
		fmt.Printf("hearthpwn.com user: %s\r\n", user)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
