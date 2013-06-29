package main

import (
	"os"
)

type Logger struct {
	out *os.File
}

func (l logger) Log(s string {
	out := l.out
	if (out == nil) {
		out = os.Stderr
	}
	fmt.Fprintf(out, "%s [%d]: %s\n", os.Args[0], os.Getpid(), s)
}

func (l *Logger) SetOutput(out *os.File) {
	l.out
}

