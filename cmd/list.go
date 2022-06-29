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
	Short: "Shows a list of all actions and participants.",
	Long:  `Shows a list of all actions and participants.`,
	Run: func(cmd *cobra.Command, args []string) {
		rec, err := recording.LoadRecords(recording.DEFAULT_RECORD_FILENAME)
		if err != nil {
			os.Exit(1)
		}

		fmt.Println("Actions:")
		for _, a := range rec.GetActionNames() {
			fmt.Printf("    %s\n", a)
		}

		fmt.Println("Participants:")
		for _, p := range rec.GetParticipants() {
			fmt.Printf("    %s\n", p)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// List Flags
	// like so:
	// statsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
