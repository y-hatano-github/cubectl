package cmd

import (
	g "cubectl/graphics"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
)

func RunCube(cmd *cobra.Command, args []string) error {
	if output == "" {
		output = "wireframe" // default
	}
	switch output {
	case "wireframe", "solid":
		// ok
	default:
		return fmt.Errorf("unknown output format %q", output)
	}
	w, _ := cmd.Flags().GetBool("watch") // watch フラグ

	pid := os.Getpid()
	logs := []string{
		fmt.Sprintf("%s %5d loader.go:223] Error loading kubeconfig:\n", cubeTimestamp(), pid),
		fmt.Sprintf("unable to read config file %q: no such file or directory\n", "/home/user/.kube/config"),
		fmt.Sprintf("%s %5d round_trippers.go:45] Failed to create Kubernetes client:\n", cubeTimestamp(), pid),
		"no configuration has been provided\n",
		fmt.Sprintf("%s %5d command.go:112] error: unknown command %q\n\n", cubeTimestamp(), pid, "kubectl"),
		fmt.Sprintf("%s %5d command.go:112] This is not \"kubectl\" but \"cubectl\"", cubeTimestamp(), pid),
		"Did you mean this?\n",
		"    kubectl\n\n",
	}
	logIndex := 0

	// Cube vertices
	v := g.VertexData{
		[3]int{-2, -2, -2},
		[3]int{2, -2, -2},
		[3]int{-2, 2, -2},
		[3]int{2, 2, -2},
		[3]int{-2, -2, 2},
		[3]int{2, -2, 2},
		[3]int{-2, 2, 2},
		[3]int{2, 2, 2},
	}

	f := g.FaceData{
		[]int{0, 1, 3, 2},
		[]int{5, 4, 6, 7},
		[]int{0, 1, 5, 4},
		[]int{3, 2, 6, 7},
		[]int{0, 2, 6, 4},
		[]int{3, 1, 5, 7},
	}

	m := g.NewModel(40, 20)
	m.Set(v, f)

	if err := termbox.Init(); err != nil {
		return err
	}
	defer termbox.Close()

	termbox.SetOutputMode(termbox.Output256)
	termbox.Clear(termbox.ColorDefault, termbox.ColorBlack)

	ch := make(chan termbox.Event)
	go keyEvent(ch)

	yaw := 0.0
	pitch := 0.0
	scale := 0.6

	drawString := func(x, y int, str string) {
		for i, r := range str {
			termbox.SetCell(x+i, y, r, termbox.ColorDefault, termbox.ColorBlack)
		}
	}

loop:
	for {
		select {
		case ev := <-ch:
			switch ev.Type {
			case termbox.EventKey:
				if ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyEsc {
					break loop
				}
				if ev.Key == termbox.KeyArrowLeft || string(ev.Ch) == "a" {
					yaw -= 0.1
				}
				if ev.Key == termbox.KeyArrowRight || string(ev.Ch) == "d" {
					yaw += 0.1
				}
				if ev.Key == termbox.KeyArrowUp || string(ev.Ch) == "w" {
					pitch -= 0.1
				}
				if ev.Key == termbox.KeyArrowDown || string(ev.Ch) == "s" {
					pitch += 0.1
				}
				if string(ev.Ch) == "z" {
					scale += 0.1
				}
				if string(ev.Ch) == "x" {
					scale -= 0.1
					scale = math.Max(0.1, scale-0.1)
				}
			}
		default:
			termbox.Clear(termbox.ColorDefault, termbox.ColorBlack)

			for l := 0; l < logIndex; l++ {
				drawString(0, l, logs[l])
			}
			if logIndex < len(logs) {
				logIndex++
			}

			if w {
				yaw += 0.02
				pitch += 0.01
			}

			faceData := m.GetShape(yaw, pitch, scale, 20, 10)
			for _, fd := range faceData {
				if output == "solid" {
					for _, p := range fd.Fill {
						termbox.SetCell(p.X, p.Y, ' ', termbox.ColorDefault, termbox.ColorBlack)
					}
				}
				for _, p := range fd.Outline {
					termbox.SetCell(p.X, p.Y, ' ', termbox.ColorDefault, termbox.ColorWhite)
				}
			}

			termbox.Flush()
			time.Sleep(10 * time.Millisecond)
		}
	}

	return nil
}

func keyEvent(ch chan termbox.Event) {
	for {
		ch <- termbox.PollEvent()
	}
}

func cubeTimestamp() string {
	now := time.Now()
	return fmt.Sprintf(
		"E%s %s",
		now.Format("0102"),            // MMDD
		now.Format("15:04:05.000000"), // HH:MM:SS.microsec
	)
}
