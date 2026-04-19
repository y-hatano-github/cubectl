// cmd/describe_cube.go
package describe

import (
	"cubectl/internal/app/describe"
	"cubectl/internal/cmd/template"

	"github.com/spf13/cobra"
)

func NewDescribeCubeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cube [name]",
		Short: "Describe a cube",
		Long:  `Show detailed information about a cube in a funny kubectl-like style.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			opts := describe.Options{}

			name := "default-cube"
			opts.Name = &name
			if len(args) > 0 {
				opts.Name = &args[0]
			}

			describe.DescribePod(cmd.Context(), opts)
		},
	}
	cmd.SetHelpTemplate(template.DescribePodHelpTemplate)
	cmd.SetUsageTemplate(template.DescribePodUsageTemplate)

	return cmd
}

func NewDescribeCubesCmd() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "cubes",
		Short: "Describe all cube",
		Long:  `Show detailed information about all cubes (joke) in cubectl style.`,
		Run: func(cmd *cobra.Command, args []string) {
			describe.DescribePods(cmd.Context())
		},
	}
	cmd.SetHelpTemplate(template.DescribePodHelpTemplate)
	cmd.SetUsageTemplate(template.DescribePodUsageTemplate)

	return cmd
}
