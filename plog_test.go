package plog

import (
	"bytes"
	"strings"
	"testing"
)

func TestLogFuncs(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(buf)
	l.SetDebug(true)

	msg := "test"

	tcs := []struct {
		in   string
		f    func(f string, args ...interface{})
		want string
	}{
		{in: msg, f: l.Infof, want: "[INFO] " + msg},
		{in: msg, f: l.Debugf, want: "[DEBUG] " + msg},
		{in: msg, f: l.Errorf, want: "[ERROR] " + msg},
	}

	for _, tc := range tcs {
		tc.f(tc.in)
		if !strings.Contains(buf.String(), tc.want) {
			t.Errorf("[got] %s [want] %s", buf.String(), tc.want)
		}
	}
}
