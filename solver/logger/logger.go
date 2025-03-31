package logger 

import (
	"io"
	"os"
	"strings"
)

type Logger struct {
	logs []string
}

func (l *Logger) Log(s string) {
	l.logs = append(l.logs, s)
}

func (l *Logger) Write(w io.Writer) (int, error) {
	return w.Write([]byte(strings.Join(l.logs, "\n")))
}

func (l *Logger) WriteStdOut() error {
	_, err := l.Write(os.Stdout)
	return err
}

