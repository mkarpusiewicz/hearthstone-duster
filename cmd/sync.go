package cmd

import (
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync all: hearthpwn.com, hsreplay.net and hearthstonejson.com",
	Run: func(cmd *cobra.Command, args []string) {
		syncDatabaseCmd.Run(cmd, args)
		syncMineCmd.Run(cmd, args)
		syncUsageCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
