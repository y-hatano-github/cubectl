package delete

import (
	"cubectl/internal/app/cube"

	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete resources",
		Long: `Delete resources from cubectl.
This command mimics 'kubectl delete' but operates on cubes and other joke resources.`,
		GroupID: "basic",
		RunE: func(cmd *cobra.Command, args []string) error {
			cube.RenderD(cmd.Context(), cube.Options{})
			return nil
		},
	}

	return cmd

}
