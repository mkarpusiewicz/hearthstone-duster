package cmd

import (
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show synced information overview",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
