package cube

import (
	"context"
	g "cubectl/internal/graphics"
	"fmt"
	"math"
	"strings"
	"time"

	"cubectl/internal/logger"
	"cubectl/internal/terminal"
)

type Options struct {
	Output string
	Watch  bool
}

func Render(ctx context.Context, opts Options) error {
	output := opts.Output
	w := opts.Watch

	if output == "" {
		output = "wireframe" // default
	}
	switch output {
	case "wireframe", "solid":
		// valid
	default:
		return fmt.Errorf("unknown output format %q", output)
	}

	clog := logger.New()

	logs := []string{
		clog.Swarn(logger.Message{File: "loader.go", Line: 0, Text: "Warning: This output is a joke."}),
		clog.Serror(logger.Message{File: "loader.go", Line: 223, Text: "Error loading kubeconfig:\nunable to read config file \"/home/user/.kube/config\": no such file or directory"}),
		clog.Serror(logger.Message{File: "round_trippers.go", Line: 45, Text: "Failed to create Kubernetes client:\nno configuration has been provided"}),
		clog.Serror(logger.Message{File: "command.go", Line: 112, Text: "error: unknown command \"kubectl\""}),
		clog.Swarn(logger.Message{File: "command.go", Line: 112, Text: "This is not \"kubectl\" but \"cubectl\"\nDid you mean this?\n    kubectl"}),
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

	s := terminal.New()
	if err := s.Init(); err != nil {
		return err
	}
	defer s.Close()

	s.SetOutputMode()
	s.Clear()

	ch := make(chan terminal.Event)
	go keyEvent(ch, s)

	yaw := 0.0
	pitch := 0.0
	scale := 0.6

	drawString := func(x, y int, str string) {
		for i, r := range str {
			s.SetCell(x+i, y, r, terminal.ColorDefault, terminal.ColorBlack)
		}
	}

loop:
	for {
		select {
		case ev := <-ch:
			switch ev.Type {
			case terminal.EventKey:
				if ev.Key == terminal.KeyCtrlC || ev.Key == terminal.KeyEsc {
					break loop
				}
				if ev.Key == terminal.KeyArrowLeft || string(ev.Rune) == "a" {
					yaw -= 0.1
				}
				if ev.Key == terminal.KeyArrowRight || string(ev.Rune) == "d" {
					yaw += 0.1
				}
				if ev.Key == terminal.KeyArrowUp || string(ev.Rune) == "w" {
					pitch -= 0.1
				}
				if ev.Key == terminal.KeyArrowDown || string(ev.Rune) == "s" {
					pitch += 0.1
				}
				if string(ev.Rune) == "z" {
					scale += 0.1
				}
				if string(ev.Rune) == "x" {
					scale -= 0.1
					scale = math.Max(0.1, scale-0.1)
				}
			}
		default:
			s.Clear()

			r := 0
			for l := range logIndex {
				lines := strings.Split(logs[l], "\n")
				for _, line := range lines {
					drawString(0, r, line)
					r = r + 1
				}
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
						s.SetCell(p.X, p.Y, ' ', terminal.ColorDefault, terminal.ColorBlack)
					}
				}
				for _, p := range fd.Outline {
					s.SetCell(p.X, p.Y, ' ', terminal.ColorDefault, terminal.ColorWhite)
				}
			}

			s.Flush()
			time.Sleep(10 * time.Millisecond)
		}
	}

	return nil
}

func keyEvent(ch chan terminal.Event, s terminal.Screen) {
	for {
		ch <- s.PollEvent()
	}
}

func CubeTimestamp() string {
	now := time.Now()
	return fmt.Sprintf(
		"%s %s",
		now.Format("0102"),            // MMDD
		now.Format("15:04:05.000000"), // HH:MM:SS.microsec
	)
}
