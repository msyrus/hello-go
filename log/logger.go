package log

import (
	"log"
	"os"
)

// Logger is an interface to print logs
type Logger interface {
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

// DefaultErrorLogger logs into stderr
var DefaultErrorLogger = log.New(os.Stderr, "", log.LstdFlags)

// DefaultOutputLogger logs into stdout
var DefaultOutputLogger = log.New(os.Stdout, "", log.LstdFlags)
