// cmd/describe_cube.go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var describePodsCmd = &cobra.Command{
	Use:   "pods [name]",
	Short: "Describe a cube",
	Long:  `Show detailed information about a pod in a funny kubectl-like style.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := "default-cube"
		if len(args) > 0 {
			name = args[0]
		}

		fmt.Printf("Name: %s\n", name)
		fmt.Println("Namespace: default")
		fmt.Println("Type: 3D Cube")
		fmt.Println("Status: Ready")
		fmt.Println("Vertices: 8")
		fmt.Println("Edges: 12")
		fmt.Println("Faces: 6")
		fmt.Println("Rotation: 0,0,0")
		fmt.Println("Message: This is not kubectl, but cubectl!")
	},
}

func init() {
	describeCmd.AddCommand(describePodsCmd)
}
