package logs

import (
	"context"
	"cubectl/internal/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	logs := []logger.LogMessage{
		{File: "loader.go", Line: 0, Text: "Warning: This output is a joke.", Level: logger.Warn, TimeStamp: logger.Timestamp()},
		{File: "loader.go", Line: 223, Text: "Error loading kubeconfig:\nunable to read config file \"/home/user/.kube/config\": no such file or directory", Level: logger.Error, TimeStamp: logger.Timestamp()},
		{File: "round_trippers.go", Line: 45, Text: "Failed to create Kubernetes client:\nno configuration has been provided", Level: logger.Error, TimeStamp: logger.Timestamp()},
		{File: "command.go", Line: 112, Text: "error: unknown command \"kubectl\"", Level: logger.Error, TimeStamp: logger.Timestamp()},
		{File: "cube.go", Line: 112, Text: fmt.Sprintf("Initializing cube rendering engine for %q", cubeName), Level: logger.Info, TimeStamp: logger.Timestamp()},
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	log := logger.New()
	lcnt := 0
	step := 1
Loop:
	for {
		select {
		case <-ch:
			break Loop
		default:
			if lcnt < len(logs) {
				fmt.Println(logs[lcnt].String())
				lcnt++
			} else {
				if ots.Follow {
					log.Info(logger.Message{File: "cube.go", Line: 100, Text: fmt.Sprintf("rotating cube... step %d", step)})
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
