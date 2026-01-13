package get

import (
	"github.com/spf13/cobra"

	cube "cubectl/internal/cube"
)

func NewGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Display one or many cubes",
		Long: `Display one or many cubes.

This command mimics 'kubectl get' but operates on cubes and other joke resources.`,
		Example: `
  # Output cube
  cubectl get [pod|cube]

  # Output cube as wireframe
  cubectl get [pod|cube] -o wireframe

  # Output cube as solid
  cubectl get [pod|cube] -o solid`,
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

	cmd.AddCommand(NewGetPodsCmd())
	cmd.AddCommand(NewGetPodCmd())

	cmd.AddCommand(NewGetCubesCmd())
	cmd.AddCommand(NewGetCubeCmd())

	return cmd
}
