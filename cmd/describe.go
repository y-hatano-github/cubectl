package cmd

import (
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe resources",
	Long: `Describe resources from cubectl.
This command mimics 'kubectl describe' but operates on cubes and other joke resources.`,
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
