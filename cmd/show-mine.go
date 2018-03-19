package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/mkarpusiewicz/hearthstone-duster/database"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var showMineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Show mine cards from hearthpwn.com last synchronization",
	Run: func(cmd *cobra.Command, args []string) {
		user := database.GetHearthPwnUser()
		if user == "" {
			log.Fatal("No user specified, please use 'user' command to set user name first")
		}
		fmt.Printf("\r\nhearthpwn.com user: %s\r\n", user)

		syncTime := database.GetMyCardsSyncTime()
		if syncTime.IsZero() {
			fmt.Print("hearthpwn.com was not synced yet\r\n")
			return
		}
		fmt.Printf("hearthpwn.com synced at: %s\r\n\r\n", syncTime.Format(time.RFC1123))

		myCards := database.GetMyCards()

		cardsCount := 0
		goldenCount := 0

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Card", "Count", "Type"})
		table.SetAutoWrapText(false)
		table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_CENTER})
		for _, card := range myCards {
			if card.IsGold {
				table.Append([]string{color.YellowString(card.Name), color.YellowString(strconv.Itoa(card.Count)), color.YellowString("GOLDEN")})
				goldenCount += card.Count
			} else {
				table.Append([]string{card.Name, strconv.Itoa(card.Count), ""})
				cardsCount += card.Count
			}
		}

		table.SetFooter([]string{"Total", strconv.Itoa(cardsCount + goldenCount), " "})

		table.Render()
	},
}

func init() {
	showCmd.AddCommand(showMineCmd)
}
