package cmd

import (
	"fmt"
	"os"

	"github.com/Sam36502/Gintokipona/recording"
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows a list of all actions.",
	Long:  `Shows a list of all actions.`,
	Run: func(cmd *cobra.Command, args []string) {
		rec, err := recording.LoadRecords(recording.DEFAULT_RECORD_FILENAME)
		if err != nil {
			os.Exit(1)
		}

		fmt.Println("Actions:")
		for _, a := range rec.ActionList {
			fmt.Printf("    %s\n", a.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// List Flags
	// like so:
	// statsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
