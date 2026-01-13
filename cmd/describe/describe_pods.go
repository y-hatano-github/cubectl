// cmd/describe_cube.go
package describe

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDescribePodCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "pod [name]",
		Short: "Describe a pod",
		Long:  `Show detailed information about a pod in a funny kubectl-like style.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := "default-pod"
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
			fmt.Println("Containers:")
			fmt.Println("  - container1: Rotating")
			fmt.Println("  - container2: Terminated")
			fmt.Println("Messages: This is not kubectl, but cubectl!")
		},
	}
}

func NewDescribePodsCmd() *cobra.Command {
	return &cobra.Command{

		Use:   "pods",
		Short: "Describe all cube",
		Long:  `Show detailed information about all pods (joke) in cubectl style.`,
		Run: func(cmd *cobra.Command, args []string) {
			// ダミーで複数のpodを表示
			pods := []string{"cube-1", "cube-2", "cube-3"}
			for _, pod := range pods {
				fmt.Printf("Name: %s\n", pod)
				fmt.Println("Namespace: default")
				fmt.Println("Type: 3D Cube")
				fmt.Println("Status: Rotating")
				fmt.Println("Vertices: 8")
				fmt.Println("Edges: 12")
				fmt.Println("Faces: 6")
				fmt.Println("Rotation: 0,0,0")
				fmt.Println("Containers:")
				fmt.Println("  - container1: Rotating")
				fmt.Println("  - container2: Terminated")
				fmt.Println("Messages: This is not kubectl, but cubectl!")
				fmt.Println("")
			}
		},
	}
}
