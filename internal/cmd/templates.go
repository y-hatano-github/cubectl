package cmd

const cubectlHelpTemplate = `{{with or .Long .Short }}{{. | trimTrailingWhitespaces}}
{{end}}{{if gt (len .Groups) 0}}{{range $g := .Groups}}
{{$g.Title}}
{{- range $.Commands}}
{{- if (and (not .Hidden) (eq .GroupID $g.ID))}}
  {{rpad .Name .NamePadding }} {{.Short}}
{{- end}}
{{- end}}
{{end}}
{{end}}
{{if .HasExample}}Examples:
{{.Example | trimTrailingWhitespaces}}
{{end}}
{{if .HasAvailableLocalFlags}}Flags:
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
