package describe

import (
	"github.com/spf13/cobra"
)

func NewDescribeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe resources",
		Long: `Describe resources from cubectl.
This command mimics 'kubectl describe' but operates on cubes and other joke resources.`,
		GroupID: "troubleshooting",
	}

	cmd.AddCommand(NewDescribeCubeCmd())
	cmd.AddCommand(NewDescribeCubesCmd())

	cmd.AddCommand(NewDescribePodCmd())
	cmd.AddCommand(NewDescribePodsCmd())

	return cmd

}
