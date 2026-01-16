// cmd/describe_cube.go
package describe

import (
	"cubectl/internal/app/describe"

	"github.com/spf13/cobra"
)

func NewDescribePodCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "pod [name]",
		Short: "Describe a pod",
		Long:  `Show detailed information about a pod in a funny kubectl-like style.`,
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
}

func NewDescribePodsCmd() *cobra.Command {
	return &cobra.Command{

		Use:   "pods",
		Short: "Describe all cube",
		Long:  `Show detailed information about all pods (joke) in cubectl style.`,
		Run: func(cmd *cobra.Command, args []string) {
			describe.DescribePods(cmd.Context())
		},
	}
}
