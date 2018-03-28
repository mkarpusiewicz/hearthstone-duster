package cmd

import (
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize data from web services",
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
