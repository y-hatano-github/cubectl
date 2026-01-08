package cmd

import "github.com/spf13/cobra"

var getPodsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Display a cube",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunCube(cmd, args)
	},
}

func init() {
	getCmd.AddCommand(getPodsCmd)

}
