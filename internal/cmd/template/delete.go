package template

var DeletePodHelpTemplate = `Delete models from cubectl.

This command mimics 'kubectl delete' but operates on models.
Note: The structural integrity of the model will be compromised upon deletion.

Examples:
  # Delete a cube model. (default)
  cubectl delete cube

  # Delete a pod (pea pod) 3D model.
  cubectl delete pod

Options:
  -h, --help   help for delete

Usage:
  cubectl delete [cube|pod] [options]
`

var DeletePodUsageTemplate = `Delete models from cubectl.

This command mimics 'kubectl delete' but operates on models.
Note: The structural integrity of the model will be compromised upon deletion.

Examples:
  # Delete a cube model. (default)
  cubectl delete cube

  # Delete a pod (pea pod) 3D model.
  cubectl delete pod

Options:
  -h, --help   help for delete

Usage:
  cubectl delete [cube|pod] [options]
`
