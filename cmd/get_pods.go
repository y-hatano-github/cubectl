package cmd

import (
	cube "cubectl/internal/cube"

	"github.com/spf13/cobra"
)

var getPodsCmd = &cobra.Command{
	Use:   "pods",
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

var getPodCmd = &cobra.Command{
	Use:   "pod",
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
	getCmd.AddCommand(getPodsCmd)
	getCmd.AddCommand(getPodCmd)

	getPodsCmd.Flags().StringVarP(&output, "output", "o", "wireframe", "Output format: wireframe|solid")
	getPodsCmd.Flags().Bool("watch", false, "Auto-rotate cube")

	getPodCmd.Flags().StringVarP(&output, "output", "o", "wireframe", "Output format: wireframe|solid")
	getPodCmd.Flags().Bool("watch", false, "Auto-rotate cube")

}
