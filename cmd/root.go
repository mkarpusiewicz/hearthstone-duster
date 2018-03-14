package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hs",
	Short: "Use card popularity from hsreplay and your collection from hearthpwn and suggest cards to dust",
	Long: `Use card popularity from hsreplay and your collection from hearthpwn and suggest cards to dust
	
Example:

# set your hearthpwn.com user name
hs user HearthPwnUserName 

# sync hsreplay.net usage data, hearthstonejson.com card database and hearthpwn.com profile data
hs sync 

# show dust propositions below 1.0% card popularity
hs dust 1.0`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
