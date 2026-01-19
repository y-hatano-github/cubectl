package logs

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cubectl/internal/logger"
)

type Options struct {
	Name   *string
	Follow bool
}

func Log(ctx context.Context, ots Options) error {
	cubeName := "cube"
	if !isEmpty(ots.Name) {
		cubeName = *ots.Name
	}

	clog := logger.New()
	logs := []string{
		clog.Swarn(logger.Message{File: "loader.go", Line: 0, Text: "Warning: This output is a joke."}),
		clog.Serror(logger.Message{File: "loader.go", Line: 223, Text: "Error loading kubeconfig:\nunable to read config file \"/home/user/.kube/config\": no such file or directory"}),
		clog.Serror(logger.Message{File: "round_trippers.go", Line: 45, Text: "Failed to create Kubernetes client:\nno configuration has been provided"}),
		clog.Serror(logger.Message{File: "command.go", Line: 112, Text: "error: unknown command \"kubectl\""}),
		clog.Swarn(logger.Message{File: "command.go", Line: 112, Text: "This is not \"kubectl\" but \"cubectl\"\nDid you mean this?\n    kubectl"}),
		clog.Sinfo(logger.Message{File: "cube.go", Line: 112, Text: fmt.Sprintf("Initializing cube rendering engine for %q", cubeName)}),
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	lcnt := 0
	step := 1
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
				if ots.Follow {
					clog.Info(logger.Message{File: "cube.go", Line: 100, Text: fmt.Sprintf("rotating cube... step %d", step)})
					if step >= 20 {
						step = 1
					} else {
						step++
					}
				} else {
					break Loop
				}
			}

			time.Sleep(200 * time.Millisecond)
		}
	}

	return nil
}

func isEmpty(s *string) bool {
	return s == nil || *s == ""
}
