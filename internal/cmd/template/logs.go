package template

var LogsPodHelpTemplate = `Print the logs for a cubectl.
This is not a real pod, just a joke.

Examples:

  # Begin streaming the logs of a cube 3d model.
  cubectl logs cube -f

  # Display only the most recent 20 lines of output
  cubectl logs pod --tail=20

Options:
  -f, --follow       Specify if the logs should be streamed.
  -h, --help         help for logs
      --tail int32   Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided. (default -1)

Usage:
  cubectl logs [cube|pod] [Options]
`

var LogsPodUsageTemplate = `Print the logs for a cubectl.
This is not a real pod, just a joke.

Examples:

  # Begin streaming the logs of a cube 3d model.
  cubectl logs cube -f

  # Display only the most recent 20 lines of output
  cubectl logs pod --tail=20

Options:
  -f, --follow       Specify if the logs should be streamed.
  -h, --help         help for logs
      --tail int32   Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided. (default -1)

Usage:
  cubectl logs [cube|pod] [Options]
`
