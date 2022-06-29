package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ginpona",
	Short: "A tool to record and check gintokipona progress",
	Long: `This allows you to easily record your progress in the
gintokipona challenge and check the stats.

Examples:
ginpona record tp-vocab  Records that you studied toki-pona vocab today
ginpona stats weeb       Shows the weebs progress`,
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
	// like so:
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
