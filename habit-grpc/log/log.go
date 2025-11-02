package log

import (
	"io"
	"log"
	"sync"
)

// a logger that can log messages
type Logger struct {
	sync.Mutex
	logger *log.Logger
}

// NewLogger returns a new logger
func NewLogger(output io.Writer) *Logger {
	return &Logger{
		logger: log.New(output, "", log.Ldate|log.Ltime),
	}
}

func (l *Logger) Logf(format string, args ...any) {
	l.Lock()
	defer l.Unlock()

	l.logger.Printf(format, args...)
}
