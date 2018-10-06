package plog

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Logger is used to control logging with log level functionality
type Logger struct {
	log     log.Logger
	isDebug bool
}

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

// New returns an instance of Logger
func New(w io.Writer) *Logger {
	l := &Logger{}
	l.SetOutput(w)
	return l
}

// SetOutput configures destination of logging
func (l *Logger) SetOutput(w io.Writer) {
	l.log.SetOutput(w)
}

// SetDebug enables debug log
func (l *Logger) SetDebug(isDebug bool) {
	l.isDebug = isDebug
}

// Infof outputs specified arguments as info
func (l *Logger) Infof(f string, args ...interface{}) {
	l.printf("[INFO] "+f, args...)
}

// Debugf outputs specified arguments as debug
// This function effects only if debug is enabled via SetDebug
func (l *Logger) Debugf(f string, args ...interface{}) {
	if l.isDebug {
		l.printf("[DEBUG] "+f, args...)
	}
}

// Errorf outputs specified arguments as error
func (l *Logger) Errorf(f string, args ...interface{}) {
	l.printf("[ERROR] "+f, args...)
}

func (l *Logger) printf(f string, args ...interface{}) {
	t := time.Now().UTC()
	y, m, d := t.Date()
	h, min, sec := t.Clock()
	loc := t.Location().String()
	dateFormat := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d %s",
		y, m, d, h, min, sec, loc)
	l.log.Printf(dateFormat+" "+f, args...)
}
