package cmd

import (
	"fmt"
	"os"

	cube "cubectl/internal/cube"

	"github.com/spf13/cobra"
)

const cubectlHelpTemplate = `{{with or .Long .Short }}{{. | trimTrailingWhitespaces}}
{{end}}{{if gt (len .Groups) 0}}{{range $g := .Groups}}
{{$g.Title}}
{{- range $.Commands}}
{{- if (and (not .Hidden) (eq .GroupID $g.ID))}}
  {{rpad .Name .NamePadding }} {{.Short}}
{{- end}}
{{- end}}
{{end}}
{{end}}{{if .HasAvailableLocalFlags}}Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

{{end}}{{if .HasAvailableInheritedFlags}}Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}

{{end}}{{if .Runnable}}Usage:
  {{.UseLine}}

{{end}}{{if .HasAvailableSubCommands}}Usage:
  {{.CommandPath}} [command]
{{end}}{{if .HasAvailableSubCommands}}
Use "{{.CommandPath}} <command> --help" for more information about a command.
{{end}}`

const cubectlUsageTemplate = `Usage:
  {{.UseLine}}

{{if .HasAvailableSubCommands}}Available Commands:
{{range .Commands}}
{{- if (and .IsAvailableCommand (not .Hidden))}}
  {{rpad .Name .NamePadding }} {{.Short}}
{{- end}}
{{end}}
{{end}}

{{if .HasAvailableLocalFlags}}Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
{{end}}

{{if .HasAvailableInheritedFlags}}
Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}
{{end}}

Use "{{.CommandPath}} [command] --help" for more information about a command.
`

var rootCmd = &cobra.Command{
	Use:   "cubectl",
	Short: "cubectl controls cube instead of Kubernetes clusters.",
	Long: `cubectl controls cube instead of Kubernetes clusters.

Find more information at:
  https://github.com/y-hatano-github/cubectl
  
Controls:
  Arrow keys or wasd: Rotate the cube
  z: Zoom in
  x: Zoom out
  Ctrl+C or Esc: Exit`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		watch, _ := cmd.Flags().GetBool("watch")

		opts := cube.Options{
			Output: output,
			Watch:  watch,
		}
		// default action
		cube.Run(cmd.Context(), opts)
	},
	SilenceUsage: true,
}

func init() {
	rootCmd.SetHelpTemplate(cubectlHelpTemplate)
	rootCmd.SetUsageTemplate(cubectlUsageTemplate)

	rootCmd.AddGroup(
		&cobra.Group{
			ID:    "basic",
			Title: "Basic Cube Commands (Beginner):",
		},
		&cobra.Group{
			ID:    "troubleshooting",
			Title: "Troubleshooting and Debugging Commands:",
		},
	)
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
