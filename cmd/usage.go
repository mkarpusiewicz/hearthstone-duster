package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/mkarpusiewicz/hearthstone-duster/services"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// usageCmd represents the usage command
var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Sync cards usage from hsreplay.net",
	Run: func(cmd *cobra.Command, args []string) {
		cardUsage := services.GetUsageInformation()

		topPlayed := cardUsage.CardsPlayedDetails.Series.Data.All
		sort.Slice(topPlayed, func(i, j int) bool {
			return topPlayed[i].Total > topPlayed[j].Total
		})

		fmt.Println("\r\nTop 100 played cards:\r\n")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Card ID", "Times", "Popularity %"})
		table.SetAutoWrapText(false)
		table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER})
		for _, row := range topPlayed[:100] {
			table.Append([]string{strconv.FormatInt(row.Id, 10), strconv.FormatInt(row.Total, 10), row.Popularity})
		}

		table.Render()
	},
}

func init() {
	syncCmd.AddCommand(usageCmd)
}
