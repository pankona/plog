package plog

import (
	"io"
	"os"
	"sync"
	"time"
)

var plog *PLog

func init() {
	plog = New(os.Stdout)
}

// SetOutput configures destination of global PLog
func SetOutput(w io.Writer) {
	plog.SetOutput(w)
}

// SetDebug enables debug log
func SetDebug(isDebug bool) {
	plog.isDebug = isDebug
}

// Infof outputs specified arguments as info log using global PLog
func Infof(f string, args ...interface{}) {
	plog.Infof(f, args...)
}

// Debugf outputs specified arguments as debug log using global PLog
func Debugf(f string, args ...interface{}) {
	plog.Debugf(f, args...)
}

// Errorf outputs specified arguments as error log using global PLog
func Errorf(f string, args ...interface{}) {
	plog.Errorf(f, args...)
}

// SetPrinter sets a Printer implementation
func SetPrinter(p Printer) {
	plog.printer = p
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

// PLog is used to control logging with log level functionality
type PLog struct {
	isDebug bool
	writer  io.Writer
	printer Printer
	mu      sync.Mutex
}

// New returns an instance of PLog
func New(w io.Writer) *PLog {
	p := &PLog{
		writer: w,
	}

	dp := &defaultPrinter{now: time.Now}
	dp.setOutput(w)
	p.printer = dp

	return p
}

// SetOutput configures destination of logging
func (p *PLog) SetOutput(w io.Writer) {
	p.writer = w
}

// SetDebug enables debug log
func (p *PLog) SetDebug(isDebug bool) {
	p.isDebug = isDebug
}

// Infof outputs specified arguments as info
func (p *PLog) Infof(f string, args ...interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.printer.Printf(p.writer, LogLevelInfo, f, args...)
}

// Debugf outputs specified arguments as debug
// This function effects only if debug is enabled via SetDebug
func (p *PLog) Debugf(f string, args ...interface{}) {
	if p.isDebug {
		p.mu.Lock()
		defer p.mu.Unlock()
		p.printer.Printf(p.writer, LogLevelDebug, f, args...)
	}
}

// Errorf outputs specified arguments as error
func (p *PLog) Errorf(f string, args ...interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.printer.Printf(p.writer, LogLevelError, f, args...)
}

// SetPrinter sets a Printer implementation
func (p *PLog) SetPrinter(pr Printer) {
	p.printer = pr
}
