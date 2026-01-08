package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get cube",
	Long: `Get resources from cubectl.
This command mimics 'kubectl get' but operates on cubes and other joke resources.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunCube(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
