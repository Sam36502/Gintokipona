package cmd

import (
	"fmt"
	"time"

	"github.com/Sam36502/Gintokipona/recording"
	"github.com/hako/durafmt"
	"github.com/spf13/cobra"
)

const (
	ACTION_FLAG = "action"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Shows the stats of a given action.",
	Long:  `Shows the stats of a given action.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Load record file
		rec, err := recording.LoadRecords(cmd.Flag(FILE_SOURCE_FLAG).Value.String())
		if err != nil {
			return
		}

		// Get action arg
		if len(args) != 1 {
			fmt.Println("[Error] Exactly 1 argument is required with the name of the action.")
			fmt.Println("See `ginpona list` for a list of recorded actions.")
			return
		}
		act := rec.GetActionByName(args[0])
		if act == nil {
			fmt.Printf("[Error] No action '%s' is recorded in the data.\n", args[0])
			fmt.Println("See `ginpona list` for a list of recorded actions.")
			return
		}

		// Display Stats
		stats := rec.GetActionStats(act)

		title := fmt.Sprintf("  Action '%s' Stats:", act.Name)
		fmt.Println(title)
		for range title {
			fmt.Print("-")
		}
		fmt.Println("--")

		fmt.Printf("   Time since start: %s\n", durafmt.Parse(stats.TimeSinceStart.Round(time.Minute)).LimitToUnit("days"))
		fmt.Printf("             Period: %s\n", durafmt.Parse(act.Period.Duration()))
		fmt.Printf("    Elapsed Periods: %d\n", stats.ElapsedPeriods)
		fmt.Printf("      Total Records: %d\n", stats.TotalRecords)
		fmt.Printf(" Completion Percent: %.2f%%\n", stats.CompletionPercent)

	},
}

func init() {
	rootCmd.AddCommand(statsCmd)

	// Stats Flags
	statsCmd.Flags().StringP(ACTION_FLAG, "a", "", "Which action to fetch the stats of")
}
