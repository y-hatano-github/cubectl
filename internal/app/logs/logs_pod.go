package logs

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	g "cubectl/internal/graphics"
	"cubectl/internal/logger"
)

func LogPod(ctx context.Context, ots Options) error {
	podName := "pod"
	if !isEmpty(ots.Name) {
		podName = *ots.Name
	}

	clog := logger.New()
	logs := []string{
		clog.Swarn(logger.Message{File: "loader.go", Line: 0, Text: "Warning: This output is a joke."}),
		clog.Serror(logger.Message{File: "loader.go", Line: 223, Text: "Error loading kubeconfig:\nunable to read config file \"/home/user/.kube/config\": no such file or directory"}),
		clog.Serror(logger.Message{File: "round_trippers.go", Line: 45, Text: "Failed to create Kubernetes client:\nno configuration has been provided"}),
		clog.Serror(logger.Message{File: "command.go", Line: 112, Text: "error: unknown command \"kubectl\""}),
		clog.Swarn(logger.Message{File: "command.go", Line: 112, Text: "This is not \"kubectl\" but \"cubectl\"\nDid you mean this?\n    kubectl"}),
		clog.Sinfo(logger.Message{File: "cube.go", Line: 112, Text: fmt.Sprintf("Initializing cube rendering engine for %q", podName)}),
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	lcnt := 0
	step := 1

	// pea pod vertices
	v := g.VertexData{
		[3]int{6, 1, 0}, // 0: Right tip

		// --- Slice 1 (Right-mid) ---
		[3]int{4, 1, -1}, // 1: Back
		[3]int{4, 1, 1},  // 2: Front
		[3]int{4, 0, 0},  // 3: Bottom curve
		[3]int{4, 2, 0},  // 4: Top curve

		// --- Slice 2 (Right-center) ---
		[3]int{2, 0, -2}, // 5: Back
		[3]int{2, 0, 2},  // 6: Front
		[3]int{2, -1, 0}, // 7: Bottom curve
		[3]int{2, 1, 0},  // 8: Top curve

		// --- Slice 3 (Left-center) ---
		[3]int{-2, 0, -2}, // 9: Back
		[3]int{-2, 0, 2},  // 10: Front
		[3]int{-2, -1, 0}, // 11: Bottom curve
		[3]int{-2, 1, 0},  // 12: Top curve

		// --- Slice 4 (Left-mid) ---
		[3]int{-4, 1, -1}, // 13: Back
		[3]int{-4, 1, 1},  // 14: Front
		[3]int{-4, 0, 0},  // 15: Bottom curve
		[3]int{-4, 2, 0},  // 16: Top curve

		[3]int{-6, 1, 0}, // 17: Left tip
	}

	// pea pod faces
	f := g.FaceData{
		// --- Top Tip Cap ---
		[]int{0, 1, 4}, // Front-left
		[]int{0, 4, 2}, // Front-right
		[]int{0, 2, 3}, // Back-right
		[]int{0, 3, 1}, // Back-left

		// --- Upper Segment ---
		[]int{1, 5, 8, 4}, // Front-left panel
		[]int{4, 8, 6, 2}, // Front-right panel
		[]int{2, 6, 7, 3}, // Back-right panel
		[]int{3, 7, 5, 1}, // Back-left panel

		// --- Middle Segment ---
		[]int{5, 9, 12, 8},
		[]int{8, 12, 10, 6},
		[]int{6, 10, 11, 7},
		[]int{7, 11, 9, 5},

		// --- Lower Segment ---
		[]int{9, 13, 16, 12},
		[]int{12, 16, 14, 10},
		[]int{10, 14, 15, 11},
		[]int{11, 15, 13, 9},

		// --- Bottom Tip Cap ---
		[]int{17, 16, 13}, // Front-left
		[]int{17, 14, 16}, // Front-right
		[]int{17, 15, 14}, // Back-right
		[]int{17, 13, 15}, // Back-left
	}

	m := g.NewModel(v, f, 8)

	yaw := 0.48
	pitch := 0.24
	scale := 0.2
	twoPi := 2 * math.Pi
Loop:
	for {
		select {
		case <-ch:
			break Loop
		default:
			if lcnt < len(logs) {
				fmt.Println(logs[lcnt])
				lcnt++
			} else {
				if ots.Follow || ots.Tail > 0 {
					clog.Info(logger.Message{
						File: "cube.go",
						Line: 88,
						Text: fmt.Sprintf("Telemetry: yaw=%.2f pitch=%.2f scale=%.1f", yaw, pitch, scale),
					})
					if step >= 20 {
						step = 1
					} else {
						step++
					}
					yaw = math.Mod(yaw+0.08, twoPi)
					pitch = math.Mod(pitch+0.04, twoPi)
					drawObject(&m, yaw, pitch, scale)
					time.Sleep(300 * time.Millisecond)
				} else {
					drawObject(&m, yaw, pitch, scale)
					break Loop
				}
			}

			if ots.Tail > 0 {
				ots.Tail--
				if ots.Tail == 0 {
					break Loop
				}
			}

		}
	}

	return nil
}
