package describe

import (
	"context"
	"fmt"
)

type Options struct {
	Name *string
}

func DescribePods(ctx context.Context) {
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
}

func DescribePod(ctx context.Context, ots Options) {

	name := "default-cube"
	if !isEmpty(ots.Name) {
		name = *ots.Name
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
}

func isEmpty(s *string) bool {
	return s == nil || *s == ""
}
