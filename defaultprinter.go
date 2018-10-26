package plog

import (
	"fmt"
	"io"
	"log"
	"time"
)

type defaultPrinter struct {
	log log.Logger
}

func (p *defaultPrinter) setOutput(w io.Writer) {
	p.log.SetOutput(w)
}

var logLevelToString = map[LogLevel]string{
	LogLevelInfo:  "[INFO]",
	LogLevelDebug: "[DEBUG]",
	LogLevelError: "[ERROR]",
}

// Printf is default print function.
func (p *defaultPrinter) Printf(_ io.Writer, lvl LogLevel, format string, args ...interface{}) {
	t := time.Now().UTC()
	y, m, d := t.Date()
	h, min, sec := t.Clock()
	loc := t.Location().String()
	dateFormat := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d %s",
		y, m, d, h, min, sec, loc)

	p.log.Printf(dateFormat+" "+logLevelToString[lvl]+" "+format, args...)
}
