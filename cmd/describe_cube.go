// cmd/describe_cube.go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var describeCubeCmd = &cobra.Command{
	Use:   "cube [name]",
	Short: "Describe a cube",
	Long:  `Show detailed information about a cube in a funny kubectl-like style.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := "default-cube"
		if len(args) > 0 {
			name = args[0]
		}

		fmt.Printf("Name: %s\n", name)
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

var describeCubesCmd = &cobra.Command{

	Use:   "cubes",
	Short: "Describe all cube",
	Long:  `Show detailed information about all cubes (joke) in cubectl style.`,
	Run: func(cmd *cobra.Command, args []string) {
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

func init() {
	describeCmd.AddCommand(describeCubeCmd)
	describeCmd.AddCommand(describeCubesCmd)
}
