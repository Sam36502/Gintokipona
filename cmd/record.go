package cmd

import (
	"fmt"
	"time"

	"github.com/Sam36502/Gintokipona/recording"
	"github.com/spf13/cobra"
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Records an entry that a certain action was completed.",
	Long:  `Records an entry that a certain action was completed.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Load record file
		filename := cmd.Flag(FILE_SOURCE_FLAG).Value.String()
		rec, err := recording.LoadRecords(filename)
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

		// Increment/Append record for today
		i, r := rec.GetActionRecordByDate(act, time.Now())
		if r == nil {
			rec.RecordList = append(rec.RecordList, recording.Record{
				ActionName: act.Name,
				Time:       recording.JSONTime(time.Now()),
				Amount:     1,
			})
		} else {
			rec.RecordList[i].Amount++
		}

		// Save records
		err = recording.SaveRecords(filename, rec)
		if err != nil {
			return
		}
		fmt.Printf("Fulfilment of '%s' recorded in '%s'.\n", act.Name, filename)

	},
}

func init() {
	rootCmd.AddCommand(recordCmd)

	// Record Flags
}
