package describe

import (
	"context"
	"fmt"
	"net"
	"time"
)

type Options struct {
	Name *string
}

func DescribePods(ctx context.Context) {
	pods := []string{"cube-1", "cube-2", "cube-3"}

	for _, pod := range pods {
		printPod(pod)
		fmt.Println("")
	}
}

func DescribePod(ctx context.Context, ots Options) {
	name := "default-cube"
	if !isEmpty(ots.Name) {
		name = *ots.Name
	}

	printPod(name)
}

func isEmpty(s *string) bool {
	return s == nil || *s == ""
}

func printPod(name string) {
	ip := "127.0.0.1"
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err == nil {
		ip = conn.LocalAddr().(*net.UDPAddr).IP.String()
	}

	defer conn.Close()

	startTime := time.Now().Add(-10 * time.Minute).Format(time.RFC1123Z)

	fmt.Printf("Name:         %s\n", name)
	fmt.Printf("Namespace:    default\n")
	fmt.Printf("Priority:     0\n")
	fmt.Printf("Node:         geometry-node-01/192.168.1.1\n")
	fmt.Printf("Start Time:   %s\n", startTime)
	fmt.Printf("Labels:       app=cubectl\n")
	fmt.Printf("              shape=cube\n")
	fmt.Printf("Annotations:  <none>\n")
	fmt.Printf("Status:       Rotating\n")
	fmt.Printf("IP:           %s\n", ip)
	fmt.Printf("IPs:\n")
	fmt.Printf("  IP:           %s\n", ip)

	fmt.Println("\nContainers:")
	fmt.Println("  cube-container:")
	fmt.Println("    Container ID:   containerd://abcdef123456")
	fmt.Println("    Image:          geometry/cube:latest")
	fmt.Println("    Port:           <none>")
	fmt.Println("    Host Port:      <none>")
	fmt.Println("    State:          Running")
	fmt.Printf("      Started:      %s\n", startTime)
	fmt.Println("    Ready:          True")
	fmt.Println("    Restart Count:  0")
	fmt.Println("    Environment:    <none>")

	fmt.Println("\nConditions:")
	fmt.Println("  Type              Status")
	fmt.Println("  Initialized       True ")
	fmt.Println("  Ready             True ")
	fmt.Println("  ContainersReady   True ")
	fmt.Println("  PodScheduled      True ")

	fmt.Println("\nGeometry-Specs:")
	fmt.Println("  Vertices:  8")
	fmt.Println("  Edges:     12")
	fmt.Println("  Faces:     6")

	fmt.Println("\nEvents:")
	fmt.Println("  Type    Reason     Age    From               Message")
	fmt.Println("  ----    ------     ----   ----               -------")
	fmt.Println("  Normal  Scheduled  11m    default-scheduler  Successfully assigned default/cube to geometry-node-01")
	fmt.Println("  Normal  Pulling    11m    kubelet            Pulling image \"geometry/cube:latest\"")
	fmt.Println("  Normal  Pulled     10m    kubelet            Successfully pulled image \"geometry/cube:latest\"")
	fmt.Println("  Normal  Created    10m    kubelet            Created container cube-container")
	fmt.Println("  Normal  Started    10m    kubelet            Started container cube-container")
}
