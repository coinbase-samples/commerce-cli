package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "commerce",
	Short: "Commerce is a CLI for managing e-commerce operations.",
	Long:  `A fast and flexible CLI built with Cobra for managing e-commerce operations.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(eventsCmd)
	rootCmd.AddCommand(chargesCmd)

}
