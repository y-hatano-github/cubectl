package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	cube "cubectl/internal/cube"
)

var logsCmd = &cobra.Command{
	Use:   "logs [cube]",
	Short: "Print the logs for a cubectl",
	Long:  `Print the logs for a cubectl. This is not a real pod, just a joke.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cubeName := "cube"
		if len(args) == 1 {
			cubeName = args[0]
		}

		pid := os.Getpid()
		logs := []string{
			fmt.Sprintf("%s %5d loader.go:223] Error loading kubeconfig:\n", cube.CubeTimestamp(), pid),
			fmt.Sprintf("unable to read config file %q: no such file or directory\n", "/home/user/.kube/config"),
			fmt.Sprintf("%s %5d round_trippers.go:45] Failed to create Kubernetes client:\n", cube.CubeTimestamp(), pid),
			"no configuration has been provided\n",
			fmt.Sprintf("%s %5d command.go:112] error: unknown command %q\n\n", cube.CubeTimestamp(), pid, "kubectl"),
			fmt.Sprintf("%s %5d command.go:112] This is not \"kubectl\" but \"cubectl\"\n", cube.CubeTimestamp(), pid),
			fmt.Sprintf("%s %5d cube.go:99] Initializing cube rendering engine for %q\n", cube.CubeTimestamp(), pid, cubeName),
		}

		for _, l := range logs {
			fmt.Print(l)
		}

		follow, _ := cmd.Flags().GetBool("follow")
		if follow {
			for i := 0; i < 10; i++ {
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("%s %5d cube.go:100] rotating cube... step %d\n", cube.CubeTimestamp(), pid, i+1)
			}
		}
	},
	GroupID: "troubleshooting",
}

func init() {
	rootCmd.AddCommand(logsCmd)
	logsCmd.Flags().BoolP("follow", "f", false, "Follow the cube logs")
}
