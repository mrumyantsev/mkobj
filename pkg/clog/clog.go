package clog

import "os"

const (
	eol           = "\n"
	fatalExitCode = 1
)

func Fatal(msg string) {
	os.Stderr.WriteString(msg + eol)

	os.Exit(fatalExitCode)
}
