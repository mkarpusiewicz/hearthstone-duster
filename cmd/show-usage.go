package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/mkarpusiewicz/hearthstone-duster/database"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var showUsageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Sync cards usage from hsreplay.net",
	Run: func(cmd *cobra.Command, args []string) {
		syncTime := database.GetCardsUsageSyncTime()
		if syncTime.IsZero() {
			fmt.Print("hsreplay.net was not synced yet\r\n")
			return
		}
		fmt.Printf("hsreplay.net synced at: %s\r\n\r\n", syncTime.Format(time.RFC1123))

		cardUsage := database.GetCardsUsage()

		topPlayed := cardUsage.CardsPlayedDetails.Series.Data.All
		sort.Slice(topPlayed, func(i, j int) bool {
			return topPlayed[i].Total > topPlayed[j].Total
		})

		fmt.Println("Top 100 played cards:")
		fmt.Println()

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
	showCmd.AddCommand(showUsageCmd)
}
