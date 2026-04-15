package logs

import (
	"github.com/spf13/cobra"

	"cubectl/internal/app/logs"
	"cubectl/internal/cmd/template"
)

func NewLogsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs [CUBE]",
		Short: "Print the logs for a cubectl",
		Long:  `Print the logs for a cubectl. This is not a real pod, just a joke.`,
		Example: `
  # Begin streaming the logs of a cube
  cubectl logs [CUBE] -f

  # Display only the most recent 20 lines of output
  cubectl logs [CUBE] --tail=20`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			opts := logs.Options{}
			opts.Follow, _ = cmd.Flags().GetBool("follow")
			opts.Tail, _ = cmd.Flags().GetInt32("tail")

			if len(args) == 1 {
				opts.Name = &args[0]
			}

			logs.Log(cmd.Context(), opts)
		},
		GroupID: "troubleshooting",
	}

	cmd.SetHelpTemplate(template.CubectlHelpTemplate)
	cmd.SetUsageTemplate(template.CubectlUsageTemplate)
	cmd.Flags().BoolP("follow", "f", false, "Specify if the logs should be streamed.")
	cmd.Flags().Int32("tail", -1, "Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided.")
	return cmd
}
