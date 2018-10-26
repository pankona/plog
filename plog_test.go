package plog

import (
	"bytes"
	"testing"
	"time"
)

func TestLogFuncs(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(buf)
	l.SetDebug(true)

	dp := &defaultPrinter{}
	dp.setOutput(buf)
	now := time.Now().UTC()
	dp.now = func() time.Time {
		return now
	}
	l.SetPrinter(dp)

	msg := "test"

	tcs := []struct {
		in   string
		f    func(f string, args ...interface{})
		want string
	}{
		{in: msg, f: l.Infof, want: dateFormat(now) + " [INFO] " + msg + "\n"},
		{in: msg, f: l.Debugf, want: dateFormat(now) + " [DEBUG] " + msg + "\n"},
		{in: msg, f: l.Errorf, want: dateFormat(now) + " [ERROR] " + msg + "\n"},
	}

	for _, tc := range tcs {
		tc.f(tc.in)
		if buf.String() != tc.want {
			t.Errorf("[got] %s [want] %s", buf.String(), tc.want)
		}
		buf.Reset()
	}
}
