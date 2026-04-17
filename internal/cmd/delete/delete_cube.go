package delete

import (
	"cubectl/internal/app/cube"
	"cubectl/internal/cmd/template"

	"github.com/spf13/cobra"
)

func NewDeleteCubeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cube",
		Short: "Delete a cube",
		Long: `Delete a cube from cubectl.

This command mimics 'kubectl delete' but operates on cubes.
Note: The structural integrity of the cube will be compromised upon deletion.`,
		Example: `  # Delete a cube.
  cubectl delete cube`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cube.RenderD(cmd.Context(), cube.Options{})
			return nil
		},
	}

	cmd.SetHelpTemplate(template.CubectlHelpTemplate)
	cmd.SetUsageTemplate(template.CubectlUsageTemplate)
	return cmd

}
