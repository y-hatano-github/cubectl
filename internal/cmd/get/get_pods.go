package get

import (
	pod "cubectl/internal/app/pod"
	"cubectl/internal/cmd/template"

	"github.com/spf13/cobra"
)

func NewGetPodsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pods",
		Short: "Display a cube",
		RunE: func(cmd *cobra.Command, args []string) error {
			output, _ := cmd.Flags().GetString("output")
			watch, _ := cmd.Flags().GetBool("watch")

			opts := pod.Options{
				Output: output,
				Watch:  watch,
			}
			return pod.Render(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "wireframe", "Output format: wireframe|solid")
	cmd.Flags().BoolP("watch", "w", false, "Watch for changes to the cube (it will keep spinning)")

	return cmd
}

func NewGetPodCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pod",
		Short: "Display a cube",
		RunE: func(cmd *cobra.Command, args []string) error {
			output, _ := cmd.Flags().GetString("output")
			watch, _ := cmd.Flags().GetBool("watch")

			opts := pod.Options{
				Output: output,
				Watch:  watch,
			}
			return pod.Render(cmd.Context(), opts)
		},
	}
	cmd.SetHelpTemplate(template.CubectlHelpTemplate)
	cmd.SetUsageTemplate(template.CubectlUsageTemplate)
	cmd.Flags().StringVarP(&output, "output", "o", "wireframe", "Output format: wireframe|solid")
	cmd.Flags().BoolP("watch", "w", false, "Watch for changes to the cube (it will keep spinning)")

	return cmd
}
