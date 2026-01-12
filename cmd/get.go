package cmd

import (
	"github.com/spf13/cobra"

	cube "cubectl/internal/cube"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get cube",
	Long: `Get resources from cubectl.
This command mimics 'kubectl get' but operates on cubes and other joke resources.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		output, _ := cmd.Flags().GetString("output")
		watch, _ := cmd.Flags().GetBool("watch")

		opts := cube.Options{
			Output: output,
			Watch:  watch,
		}
		return cube.Run(cmd.Context(), opts)
	},
	GroupID: "basic",
}

func init() {
	rootCmd.AddCommand(getCmd)
}
