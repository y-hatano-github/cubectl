package logger

import (
	"fmt"
	"os"
	"time"
)

type LogLevel int

const (
	Info LogLevel = iota
	Warn
	Error
)

type Message struct {
	File string
	Line int
	Text string
}

type LogMessage struct {
	File      string
	Line      int
	Text      string
	Level     LogLevel
	TimeStamp string
}

func (m LogMessage) String() string {
	l := []string{"I", "W", "E"}[m.Level]
	return fmt.Sprintf("%s%s %5d %s:%d] %s", l, m.TimeStamp, os.Getpid(), m.File, m.Line, m.Text)
}

type Logger struct {
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Sinfo(m Message) string {
	return fmt.Sprintf("I%s %5d %s:%d] %s\n", Timestamp(), os.Getpid(), m.File, m.Line, m.Text)
}

func (l *Logger) Info(m Message) {
	fmt.Fprint(os.Stderr, l.Sinfo(m))
}

func (l *Logger) Serror(m Message) string {
	return fmt.Sprintf("E%s %5d %s:%d] %s\n", Timestamp(), os.Getpid(), m.File, m.Line, m.Text)
}

func (l *Logger) Error(m Message) {
	fmt.Fprint(os.Stderr, l.Serror(m))
}

func (l *Logger) Swarn(m Message) string {
	return fmt.Sprintf("W%s %5d %s:%d] %s\n", Timestamp(), os.Getpid(), m.File, m.Line, m.Text)
}

func (l *Logger) Warn(m Message) {
	fmt.Fprint(os.Stderr, l.Swarn(m))
}

func Timestamp() string {
	now := time.Now()
	return fmt.Sprintf(
		"%s %s",
		now.Format("0102"),            // MMDD
		now.Format("15:04:05.000000"), // HH:MM:SS.microsec
	)
}
