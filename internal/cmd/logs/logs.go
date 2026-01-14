package logs

import (
	"github.com/spf13/cobra"

	"cubectl/internal/app/logs"
)

func NewLogsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs [cube]",
		Short: "Print the logs for a cubectl",
		Long:  `Print the logs for a cubectl. This is not a real pod, just a joke.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			opts := logs.Options{}
			opts.Follow, _ = cmd.Flags().GetBool("follow")
			if len(args) == 1 {
				opts.Name = &args[0]
			}

			logs.Log(cmd.Context(), opts)
		},
		GroupID: "troubleshooting",
	}
	cmd.Flags().BoolP("follow", "f", false, "Follow the cube logs")

	return cmd
}
