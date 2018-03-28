package cmd

import (
	"github.com/spf13/cobra"
)

var syncAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Sync all: hearthpwn.com, hsreplay.net and hearthstonejson.com",
	Run: func(cmd *cobra.Command, args []string) {
		syncDatabaseCmd.Run(cmd, args)
		syncUsageCmd.Run(cmd, args)
		syncMineCmd.Run(cmd, args)
	},
}

func init() {
	syncCmd.AddCommand(syncAllCmd)
}
