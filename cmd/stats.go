package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Shows the stats of a given participant.",
	Long:  `Shows the stats of a given participant.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stats called")
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)

	// Stats Flags
	// like so:
	// statsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
