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
	l.Printer = p
}

//go:generate stringer -type LogLevel -linecomment
type LogLevel int

const (
	LogLevelInfo  LogLevel = iota // INFO
	LogLevelDebug                 // DEBUG
	LogLevelError                 // ERROR
)

type Printer interface {
	Printf(w io.Writer, level LogLevel, f string, args ...interface{})
}

// Logger is used to control logging with log level functionality
type Logger struct {
	isDebug bool
	writer  io.Writer
	Printer
}

// New returns an instance of Logger
func New(w io.Writer) *Logger {
	l := &Logger{
		writer: w,
	}

	dp := &defaultPrinter{now: time.Now}
	dp.setOutput(w)
	l.Printer = dp

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
	l.Printf(l.writer, LogLevelInfo, f, args...)
}

// Debugf outputs specified arguments as debug
// This function effects only if debug is enabled via SetDebug
func (l *Logger) Debugf(f string, args ...interface{}) {
	if l.isDebug {
		l.Printf(l.writer, LogLevelDebug, f, args...)
	}
}

// Errorf outputs specified arguments as error
func (l *Logger) Errorf(f string, args ...interface{}) {
	l.Printf(l.writer, LogLevelError, f, args...)
}

// SetPrinter sets a Printer implementation
func (l *Logger) SetPrinter(p Printer) {
	l.Printer = p
}
