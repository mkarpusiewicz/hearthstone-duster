package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/mkarpusiewicz/hearthstone-duster/database"
	"github.com/mkarpusiewicz/hearthstone-duster/services"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// mineCmd represents the mine command
var mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Synchronize mine cards from hearthpwn.com",
	Run: func(cmd *cobra.Command, args []string) {
		user := database.GetHearthPwnUser()

		if user == "" {
			log.Fatal("No user specified, please use 'user' command to set user name first")
		}

		myCards := services.GetUserCards(user)

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
	syncCmd.AddCommand(mineCmd)
}
