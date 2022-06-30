package cmd

import (
	"os"

	"github.com/Sam36502/Gintokipona/recording"
	"github.com/spf13/cobra"
)

const (
	FILE_SOURCE_FLAG = "file-source"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ginpona",
	Short: "A tool to record and check gintokipona progress",
	Long: `This allows you to easily record your progress in the
gintokipona challenge and check the stats.

Examples:
ginpona record tp-vocab  Records that you studied toki-pona vocab today
ginpona stats tp-vocab   Shows stats for the toki-pona vocab`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// Global Flags
	rootCmd.PersistentFlags().StringP(FILE_SOURCE_FLAG, "f", recording.DEFAULT_RECORD_FILENAME, "The file to load the data from")
}
