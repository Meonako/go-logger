package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Logger struct {
	mutex  sync.Mutex
	output io.Writer
	prefix string
	buffer []byte

	// This function will be call every time Print function is call to update the prefix
	UpdatePrefix func() string
}

var std = New(os.Stderr, "")

func New(out io.Writer, prefix string) *Logger {
	return &Logger{output: out, prefix: prefix}
}

func GetDefault() *Logger { return std }

func (l *Logger) SetPrefix(prefix string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.prefix = prefix
}

func (l *Logger) SetOutput(w io.Writer) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.output = w
}

func (l *Logger) outpuT(s string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buffer = l.buffer[:0]
	l.buffer = append(l.buffer, l.prefix...)
	l.buffer = append(l.buffer, s...)
	_, err := l.output.Write(l.buffer)
	return err
}

func (l *Logger) Print(v ...any) {
	if l.UpdatePrefix != nil {
		l.prefix = l.UpdatePrefix()
	}

	l.outpuT(fmt.Sprint(v...))
}

func (l *Logger) Println(v ...any) {
	if l.UpdatePrefix != nil {
		l.prefix = l.UpdatePrefix()
	}

	l.outpuT(fmt.Sprintln(v...))
}

func (l *Logger) Printf(format string, v ...any) {
	if l.UpdatePrefix != nil {
		l.prefix = l.UpdatePrefix()
	}

	l.outpuT(fmt.Sprintf(format, v...) + "\n")
}
