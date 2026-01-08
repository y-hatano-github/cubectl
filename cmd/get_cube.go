package cmd

import (
	"github.com/spf13/cobra"
)

var (
	output string
)

var getCubeCmd = &cobra.Command{
	Use:   "cube",
	Short: "Display a cube",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunCube(cmd, args)
	},
}

func init() {
	getCmd.AddCommand(getCubeCmd)

	getCubeCmd.Flags().StringVarP(&output, "output", "o", "wireframe", "Output format: wireframe|solid")
	getCubeCmd.Flags().Bool("watch", false, "Auto-rotate cube")

}
