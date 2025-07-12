package logger

import (
	"log"
	"os"
)

// Logger represents the application logger
type Logger struct {
	*log.Logger
}

// New creates a new logger instance
func New() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "[golang-domain-driven-design] ", log.LstdFlags|log.Lshortfile),
	}
}

// Info logs an info message
func (l *Logger) Info(v ...interface{}) {
	l.Println("[INFO]", v)
}

// Error logs an error message
func (l *Logger) Error(v ...interface{}) {
	l.Println("[ERROR]", v)
}

// Debug logs a debug message
func (l *Logger) Debug(v ...interface{}) {
	l.Println("[DEBUG]", v)
}

// Warn logs a warning message
func (l *Logger) Warn(v ...interface{}) {
	l.Println("[WARN]", v)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(v ...interface{}) {
	l.Println("[FATAL]", v)
	os.Exit(1)
}
