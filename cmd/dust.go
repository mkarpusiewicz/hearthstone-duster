package cmd

import (
	"github.com/spf13/cobra"
)

var dustCmd = &cobra.Command{
	Use:   "dust",
	Short: "Dust sugestions based on acquired data",
}

func init() {
	rootCmd.AddCommand(dustCmd)
}
