package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cubectl",
	Short: "cubectl controls cube instead of Kubernetes clusters.",
	Long: `cubectl controls cube instead of Kubernetes clusters.

Find more information at:
  https://github.com/y-hatano-github/cubectl`,
	Run: func(cmd *cobra.Command, args []string) {
		// default action
		RunCube(cmd, args)
	},
	SilenceUsage: true,
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
