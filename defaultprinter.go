package plog

import (
	"fmt"
	"io"
	"log"
	"time"
)

type defaultPrinter struct {
	log log.Logger
	now func() time.Time
}

func (p *defaultPrinter) setOutput(w io.Writer) {
	p.log.SetOutput(w)
}

// Printf is default print function.
func (p *defaultPrinter) Printf(_ io.Writer, lvl LogLevel, format string, args ...interface{}) {
	t := p.now().UTC()
	p.log.Printf(dateFormat(t)+" ["+lvl.String()+"] "+format, args...)
}

func dateFormat(t time.Time) string {
	y, m, d := t.Date()
	h, min, sec := t.Clock()
	loc := t.Location().String()
	dateFormat := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d %s",
		y, m, d, h, min, sec, loc)

	return dateFormat
}
