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
	cubeLog(fmt.Sprintf("%s %5d command.go:112] unknown command %q\n\n", cubeTimestamp(), pid, "kubectl"), 1000)
	cubeLog("Hint: Did you mean \"kubectl\"?\n\n", 1000)
	cubeLog("Initializing cube rendering engine...\n", 2500)

	// 立方体の頂点
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

	// 面（頂点インデックス）
	f := g.FaceData{
		[]int{0, 1, 3, 2},
		[]int{5, 4, 6, 7},
		[]int{0, 1, 5, 4},
		[]int{3, 2, 6, 7},
		[]int{0, 2, 6, 4},
		[]int{3, 1, 5, 7},
	}

	// モデル生成（画面上の論理サイズ）
	m := g.NewModel(40, 20) // 横幅は少し大きめに
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

	// 回転角
	yaw := 0.0   // 左右（Y軸）
	pitch := 0.0 // 上下（X軸）
	scale := 0.8 // 拡大率

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
			drawString(0, 0, fmt.Sprintf("%s %5d This is not \"kubectl\" but \"cubectl\"", ts, pid))

			// 立方体の形状（線分群）を取得
			s := m.GetShape(yaw, pitch, scale, 20, 10)

			for _, ps := range s {
				for _, p := range ps {
					// モデル側でX方向の2倍補正済みなので、そのまま描画
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
