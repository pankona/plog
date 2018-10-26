package plog

import (
	"io"
	"os"
	"time"
)

var l *Logger

func init() {
	l = New(os.Stdout)
}

// SetOutput configures destination of global Logger
func SetOutput(w io.Writer) {
	l.SetOutput(w)
}

// SetDebug enables debug log
func SetDebug(isDebug bool) {
	l.isDebug = isDebug
}

// Infof outputs specified arguments as info log using global Logger
func Infof(f string, args ...interface{}) {
	l.Infof(f, args...)
}

// Debugf outputs specified arguments as debug log using global Logger
func Debugf(f string, args ...interface{}) {
	l.Debugf(f, args...)
}

// Errorf outputs specified arguments as error log using global Logger
func Errorf(f string, args ...interface{}) {
	l.Errorf(f, args...)
}

// SetPrinter sets a Printer implementation
func SetPrinter(p Printer) {
	l.printer = p
}

// LogLevel represents log level
//go:generate stringer -type LogLevel -linecomment
type LogLevel int

const (
	// LogLevelInfo is log level for info
	LogLevelInfo LogLevel = iota // INFO

	// LogLevelDebug is log level for debug
	// Only when debug is enabled via SetDebug function,
	// this level's log will be output
	LogLevelDebug // DEBUG

	// LogLevelError is log level for error
	LogLevelError // ERROR
)

// Printer is an interface to customize logging format
type Printer interface {
	Printf(w io.Writer, level LogLevel, f string, args ...interface{})
}

// Logger is used to control logging with log level functionality
type Logger struct {
	isDebug bool
	writer  io.Writer
	printer Printer
}

// New returns an instance of Logger
func New(w io.Writer) *Logger {
	l := &Logger{
		writer: w,
	}

	dp := &defaultPrinter{now: time.Now}
	dp.setOutput(w)
	l.printer = dp

	return l
}

// SetOutput configures destination of logging
func (l *Logger) SetOutput(w io.Writer) {
	l.writer = w
}

// SetDebug enables debug log
func (l *Logger) SetDebug(isDebug bool) {
	l.isDebug = isDebug
}

// Infof outputs specified arguments as info
func (l *Logger) Infof(f string, args ...interface{}) {
	l.printer.Printf(l.writer, LogLevelInfo, f, args...)
}

// Debugf outputs specified arguments as debug
// This function effects only if debug is enabled via SetDebug
func (l *Logger) Debugf(f string, args ...interface{}) {
	if l.isDebug {
		l.printer.Printf(l.writer, LogLevelDebug, f, args...)
	}
}

// Errorf outputs specified arguments as error
func (l *Logger) Errorf(f string, args ...interface{}) {
	l.printer.Printf(l.writer, LogLevelError, f, args...)
}

// SetPrinter sets a Printer implementation
func (l *Logger) SetPrinter(p Printer) {
	l.printer = p
}
