package cmd

import (
	cube "cubectl/internal/cube"

	"github.com/spf13/cobra"
)

var (
	output string
)

var getCubeCmd = &cobra.Command{
	Use:   "cube",
	Short: "Display a cube",
	RunE: func(cmd *cobra.Command, args []string) error {
		output, _ := cmd.Flags().GetString("output")
		watch, _ := cmd.Flags().GetBool("watch")

		opts := cube.Options{
			Output: output,
			Watch:  watch,
		}
		return cube.Run(cmd.Context(), opts)
	},
}

var getCubesCmd = &cobra.Command{
	Use:   "cubes",
	Short: "Display a cube",
	RunE: func(cmd *cobra.Command, args []string) error {
		output, _ := cmd.Flags().GetString("output")
		watch, _ := cmd.Flags().GetBool("watch")

		opts := cube.Options{
			Output: output,
			Watch:  watch,
		}
		return cube.Run(cmd.Context(), opts)
	},
}

func init() {
	getCmd.AddCommand(getCubeCmd)
	getCmd.AddCommand(getCubesCmd)

	getCubeCmd.Flags().StringVarP(&output, "output", "o", "wireframe", "Output format: wireframe|solid")
	getCubeCmd.Flags().BoolP("watch", "w", false, "Watch for changes to the cube (it will keep spinning)")

	getCubesCmd.Flags().StringVarP(&output, "output", "o", "wireframe", "Output format: wireframe|solid")
	getCubesCmd.Flags().BoolP("watch", "w", false, "Watch for changes to the cube (it will keep spinning)")

}
