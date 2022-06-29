package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Records an entry that a certain action was completed.",
	Long:  `Records an entry that a certain action was completed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("record called")
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)

	// Record Flags
	// like so:
	// recordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
