package cmd

import (
	"fmt"

	"github.com/mkarpusiewicz/hearthstone-duster/services"
	"github.com/spf13/cobra"
)

// usageCmd represents the usage command
var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Sync cards usage from hsreplay.net",
	Run: func(cmd *cobra.Command, args []string) {
		cardUsage := services.GetUsageInformation()

		for _, row := range cardUsage.CardsPlayedDetails.Series.Data.All {
			fmt.Println(row)
		}
	},
}

func init() {
	syncCmd.AddCommand(usageCmd)
}
