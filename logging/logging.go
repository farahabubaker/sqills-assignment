package logging

import (
	"fmt"
)

type Logging interface {
	Info(msg string, loc string)
	Debug(msg string, loc string)
	Error(msg string, loc string)
}

type Logs struct {
	debugMode bool
}

func SetDebugMode(l *Logs) {
	l.debugMode = true
}

func (l *Logs) Info(msg string, loc string) {
	fmt.Printf("[INFO]: (%s): %s\n", loc, msg)
}

func (l *Logs) Error(msg string, loc string) {
	fmt.Printf("[ERROR]: (%s): %s\n", loc, msg)
}

func (l *Logs) Debug(msg string, loc string) {
	if l.debugMode {
		fmt.Printf("[DEBUG]: (%s): %s\n", loc, msg)
	}
}
