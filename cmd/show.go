package cmd

import (
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show synced information overview",
}

func init() {
	rootCmd.AddCommand(showCmd)
}
