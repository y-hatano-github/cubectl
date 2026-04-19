package template

var GetPodHelpTemplate = `Display a 3D model.

This command mimics 'kubectl get' but operates on cubes and pea pods.

Examples:
  # Output a cube 3D model as wireframe (default)
  cubectl get

  # Output a cube 3D model as solid
  cubectl get cube -o solid

  # Output a pod (pea pod) 3D model as wireframe
  cubectl get pod -o wireframe

Options:
  -h, --help            help for get
  -o, --output string   Output format: wireframe|solid (default "wireframe")
  -w, --watch           Watch for changes to the resource (it will keep spinning)

Usage:
  cubectl get [cube|pod] [options]
`

var GetPodUsageTemplate = `Display a 3D model.

This command mimics 'kubectl get' but operates on cubes and pea pods.

Examples:
  # Output a cube 3D model as wireframe (default)
  cubectl get

  # Output a cube 3D model as solid
  cubectl get cube -o solid

  # Output a pod (pea pod) 3D model as wireframe
  cubectl get pod -o wireframe

Options:
  -h, --help            help for get
  -o, --output string   Output format: wireframe|solid (default "wireframe")
  -w, --watch           Watch for changes to the resource (it will keep spinning)

Usage:
  cubectl get [cube|pod] [options]
`
