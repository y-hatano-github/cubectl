package main

import (
	g "cubectl/graphics"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

func cubeLog(msg string, w time.Duration) {
	fmt.Printf(msg)
	time.Sleep(w * time.Millisecond)
}

func cubeTimestamp() string {
	now := time.Now()
	return fmt.Sprintf(
		"E%s %s",
		now.Format("0102"),            // MMDD
		now.Format("15:04:05.000000"), // HH:MM:SS.microsec
	)
}

func printHelp() {
	fmt.Println(`Usage: cubectl [Flags]

Control cube in your terminal instead of controlling Kubernetes.

Controls:
  Arrow keys or wasd: Rotate the cube
  z: Zoom in
  x: Zoom out
  Ctrl+C or Esc: Exit

Flags:
  -h, --help    help for cubectl`)
}

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			printHelp()
			return
		}
	}

	pid := os.Getpid()
	cubeLog(fmt.Sprintf("%s %5d loader.go:223] Error loading kubeconfig:\n", cubeTimestamp(), pid), 1000)
	cubeLog(fmt.Sprintf("unable to read config file %q: no such file or directory\n", "/home/user/.kube/config"), 1000)
	cubeLog(fmt.Sprintf("%s %5d round_trippers.go:45] Failed to create Kubernetes client:\n", cubeTimestamp(), pid), 200)
	cubeLog("no configuration has been provided\n", 800)
	cubeLog(fmt.Sprintf("%s %5d command.go:112] error: unknown command %q\n\n", cubeTimestamp(), pid, "kubectl"), 1000)
	cubeLog("Did you mean this?\n", 1)
	cubeLog("    kubectl\n\n", 1000)
	cubeLog("Initializing cube rendering engine...\n", 2500)

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

	// Faces (vertex indices)
	f := g.FaceData{
		[]int{0, 1, 3, 2},
		[]int{5, 4, 6, 7},
		[]int{0, 1, 5, 4},
		[]int{3, 2, 6, 7},
		[]int{0, 2, 6, 4},
		[]int{3, 1, 5, 7},
	}

	// Create model (logical screen size)
	m := g.NewModel(40, 20) // Slightly wider for better aspect ratio
	m.Set(v, f)

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetOutputMode(termbox.Output256)
	termbox.Clear(termbox.ColorDefault, termbox.ColorBlack)

	ch := make(chan termbox.Event)
	go keyEvent(ch)

	// Rotation angles
	yaw := 0.0   // Left-right (Y-axis)
	pitch := 0.0 // Up-down (X-axis)
	scale := 0.8 // Zoom factor

	drawString := func(x, y int, str string) {
		runes := []rune(str)
		for i, v := range runes {
			termbox.SetCell(x+i, y, v, termbox.ColorDefault, termbox.ColorBlack)
		}
	}

	ts := cubeTimestamp()

loop:
	for {
		select {
		case ev := <-ch:
			switch ev.Type {
			case termbox.EventKey:
				if ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyEsc {
					break loop
				}
				if ev.Key == termbox.KeyArrowLeft {
					yaw -= 0.1
				}
				if ev.Key == termbox.KeyArrowRight {
					yaw += 0.1
				}
				if ev.Key == termbox.KeyArrowUp {
					pitch -= 0.1
				}
				if ev.Key == termbox.KeyArrowDown {
					pitch += 0.1
				}
				if string(ev.Ch) == "a" {
					yaw -= 0.1
				}
				if string(ev.Ch) == "d" {
					yaw += 0.1
				}
				if string(ev.Ch) == "w" {
					pitch -= 0.1
				}
				if string(ev.Ch) == "s" {
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
			drawString(0, 0, fmt.Sprintf("%s %5d command.go:112] This is not \"kubectl\" but \"cubectl\"", ts, pid))

			// Get cube shape (list of line segments)
			s := m.GetShape(yaw, pitch, scale, 20, 10)

			for _, ps := range s {
				for _, p := range ps {
					// X-direction scaling is already applied in the model, so draw as-is
					termbox.SetCell(p.X, p.Y, ' ', termbox.ColorDefault, termbox.ColorGreen)
				}
			}
			termbox.Flush()
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func keyEvent(ch chan termbox.Event) {
	for {
		ch <- termbox.PollEvent()
	}
}
