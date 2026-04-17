package delete

import (
	"cubectl/internal/app/pod"
	"cubectl/internal/cmd/template"

	"github.com/spf13/cobra"
)

func NewDeletePodCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pod",
		Short: "Delete a pea pod",
		Long: `Delete a pea pod from cubectl.

This command mimics 'kubectl delete' but operates on pea pod.
Note: The structural integrity of the pea pod will be compromised upon deletion.`,
		Example: `  # Delete a pea pod.
  cubectl delete pod`,
		RunE: func(cmd *cobra.Command, args []string) error {
			pod.RenderD(cmd.Context(), pod.Options{})
			return nil
		},
	}

	cmd.SetHelpTemplate(template.CubectlHelpTemplate)
	cmd.SetUsageTemplate(template.CubectlUsageTemplate)
	return cmd

}
