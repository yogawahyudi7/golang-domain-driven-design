package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger represents the application logger with Logrus
type Logger struct {
	*logrus.Logger
}

// PrettyJSONFormatter formats logs as indented JSON
type PrettyJSONFormatter struct {
	*logrus.JSONFormatter
}

// Format implements the logrus.Formatter interface with pretty printing
func (f *PrettyJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Get the standard JSON output first
	data, err := f.JSONFormatter.Format(entry)
	if err != nil {
		return nil, err
	}

	// Parse and re-format with indentation
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		// If indentation fails, return original
		return data, nil
	}

	// Add newline for better readability
	prettyJSON.WriteString("\n")
	return prettyJSON.Bytes(), nil
}

// Config holds logger configuration
type Config struct {
	Level       string `json:"level" mapstructure:"level"`
	Format      string `json:"format" mapstructure:"format"`             // json, text
	Output      string `json:"output" mapstructure:"output"`             // stdout, file, both
	PrettyPrint bool   `json:"pretty_print" mapstructure:"pretty_print"` // pretty print JSON logs
	FilePath    string `json:"file_path" mapstructure:"file_path"`
	MaxSize     int    `json:"max_size" mapstructure:"max_size"`       // MB
	MaxBackups  int    `json:"max_backups" mapstructure:"max_backups"` // number of backup files
	MaxAge      int    `json:"max_age" mapstructure:"max_age"`         // days
	Compress    bool   `json:"compress" mapstructure:"compress"`
}

// DefaultConfig returns default logger configuration
func DefaultConfig() *Config {
	return &Config{
		Level:       "info",
		Format:      "json",
		Output:      "both",
		PrettyPrint: false, // set to true for development
		FilePath:    "logs/app.log",
		MaxSize:     100, // 100MB
		MaxBackups:  5,
		MaxAge:      30, // 30 days
		Compress:    true,
	}
}

// New creates a new logger instance with best practices
func New(config *Config) *Logger {
	if config == nil {
		config = DefaultConfig()
	}

	log := logrus.New()

	// Set log level
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	log.SetLevel(level)

	// Set formatter
	switch config.Format {
	case "json":
		if config.PrettyPrint {
			log.SetFormatter(&PrettyJSONFormatter{
				&logrus.JSONFormatter{
					TimestampFormat: time.RFC3339,
					FieldMap: logrus.FieldMap{
						logrus.FieldKeyTime:  "timestamp",
						logrus.FieldKeyLevel: "level",
						logrus.FieldKeyMsg:   "message",
						logrus.FieldKeyFunc:  "caller",
					},
				},
			})
		} else {
			log.SetFormatter(&logrus.JSONFormatter{
				TimestampFormat: time.RFC3339,
				FieldMap: logrus.FieldMap{
					logrus.FieldKeyTime:  "timestamp",
					logrus.FieldKeyLevel: "level",
					logrus.FieldKeyMsg:   "message",
					logrus.FieldKeyFunc:  "caller",
				},
			})
		}
	default:
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
			DisableColors:   false,
			ForceColors:     true,
		})
	}

	// Set output
	switch config.Output {
	case "file":
		log.SetOutput(getFileWriter(config))
	case "stdout":
		log.SetOutput(os.Stdout)
	case "both":
		log.SetOutput(io.MultiWriter(os.Stdout, getFileWriter(config)))
	default:
		log.SetOutput(os.Stdout)
	}

	// Add caller info for development
	log.SetReportCaller(true)

	return &Logger{Logger: log}
}

// getFileWriter creates a file writer with rotation
func getFileWriter(config *Config) io.Writer {
	// Ensure log directory exists
	logDir := filepath.Dir(config.FilePath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		logrus.WithError(err).Error("Failed to create log directory")
		return os.Stdout
	}

	return &lumberjack.Logger{
		Filename:   config.FilePath,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
		LocalTime:  true,
	}
}

// WithField adds a single field to the logger
func (l *Logger) WithField(key string, value interface{}) *logrus.Entry {
	return l.Logger.WithField(key, value)
}

// WithFields adds multiple fields to the logger
func (l *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.Logger.WithFields(fields)
}

// WithError adds an error field to the logger
func (l *Logger) WithError(err error) *logrus.Entry {
	return l.Logger.WithError(err)
}

// WithContext adds request context fields
func (l *Logger) WithContext(requestID, userID, method, path string) *logrus.Entry {
	return l.Logger.WithFields(logrus.Fields{
		"request_id": requestID,
		"user_id":    userID,
		"method":     method,
		"path":       path,
	})
}

// HTTP Request logging helpers
func (l *Logger) LogHTTPRequest(method, path, clientIP, userAgent string, statusCode int, latency time.Duration, requestID string) {
	entry := l.WithFields(logrus.Fields{
		"request_id":  requestID,
		"method":      method,
		"path":        path,
		"client_ip":   clientIP,
		"user_agent":  userAgent,
		"status_code": statusCode,
		"latency_ms":  latency.Milliseconds(),
		"type":        "http_request",
	})

	if statusCode >= 500 {
		entry.Error("HTTP request completed with server error")
	} else if statusCode >= 400 {
		entry.Warn("HTTP request completed with client error")
	} else {
		entry.Info("HTTP request completed successfully")
	}
}

// Database operation logging
func (l *Logger) LogDBOperation(operation, table string, duration time.Duration, err error) {
	entry := l.WithFields(logrus.Fields{
		"operation":   operation,
		"table":       table,
		"duration_ms": duration.Milliseconds(),
		"type":        "database",
	})

	if err != nil {
		entry.WithError(err).Error("Database operation failed")
	} else {
		entry.Info("Database operation completed")
	}
}

// Business logic logging
func (l *Logger) LogBusinessEvent(event string, entityID string, userID string, metadata map[string]interface{}) {
	fields := logrus.Fields{
		"event":     event,
		"entity_id": entityID,
		"user_id":   userID,
		"type":      "business_event",
	}

	// Add metadata fields
	for k, v := range metadata {
		fields[k] = v
	}

	l.WithFields(fields).Info("Business event occurred")
}

// Security logging
func (l *Logger) LogSecurityEvent(event, userID, clientIP, details string) {
	l.WithFields(logrus.Fields{
		"event":     event,
		"user_id":   userID,
		"client_ip": clientIP,
		"details":   details,
		"type":      "security",
	}).Warn("Security event detected")
}

// Performance monitoring
func (l *Logger) LogPerformance(operation string, duration time.Duration, metadata map[string]interface{}) {
	fields := logrus.Fields{
		"operation":   operation,
		"duration_ms": duration.Milliseconds(),
		"type":        "performance",
	}

	for k, v := range metadata {
		fields[k] = v
	}

	entry := l.WithFields(fields)

	if duration > 5*time.Second {
		entry.Warn("Slow operation detected")
	} else if duration > 1*time.Second {
		entry.Info("Operation completed with moderate duration")
	} else {
		entry.Debug("Operation completed quickly")
	}
}
