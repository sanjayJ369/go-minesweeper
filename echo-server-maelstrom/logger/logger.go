package logger

import (
	"fmt"
	"io"
)

type Logger interface {
	Logf(format string, args ...any)
}

type StreamLogger struct {
	Writer io.Writer
}

func (s *StreamLogger) Logf(format string, args ...any) {
	s.Writer.Write([]byte(fmt.Sprintf(format, args...)))
}

func NewStreamLogger(writer io.Writer) Logger {
	return &StreamLogger{
		Writer: writer,
	}
}
