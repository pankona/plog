package plog

import (
	"bytes"
	"testing"
	"time"
)

func TestLogFuncs(t *testing.T) {
	buf := &bytes.Buffer{}
	p := New(buf)
	p.SetDebug(true)

	dp := &defaultPrinter{}
	dp.setOutput(buf)
	now := time.Now().UTC()
	dp.now = func() time.Time {
		return now
	}
	p.SetPrinter(dp)

	msg := "test"

	tcs := []struct {
		in   string
		f    func(f string, args ...interface{})
		want string
	}{
		{in: msg, f: p.Infof, want: dateFormat(now) + " [INFO] " + msg + "\n"},
		{in: msg, f: p.Debugf, want: dateFormat(now) + " [DEBUG] " + msg + "\n"},
		{in: msg, f: p.Errorf, want: dateFormat(now) + " [ERROR] " + msg + "\n"},
	}

	for _, tc := range tcs {
		tc.f(tc.in)
		if buf.String() != tc.want {
			t.Errorf("[got] %s [want] %s", buf.String(), tc.want)
		}
		buf.Reset()
	}
}
