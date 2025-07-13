//go:build example
// +build example

package main

import (
	"fmt"
	"golang-domain-driven-design/pkg/logger"
	"time"
)

func main() {
	fmt.Println("=== DEMONSTRASI LOG FILE NAMING BEST PRACTICES ===\n")

	// 1. Konfigurasi dengan multiple log files
	config := &logger.Config{
		Level:         "info",
		Format:        "json",
		Output:        "file",
		PrettyPrint:   true,
		LogDir:        "logs",
		AppLogFile:    "app.log",
		ErrorLogFile:  "error.log",
		AccessLogFile: "access.log",
		MaxSize:       100,
		MaxBackups:    5,
		MaxAge:        30,
		Compress:      true,
		DateFormat:    "2006-01-02",
	}

	log := logger.New(config)

	fmt.Println("1. Logging ke app.log (main application):")
	log.Info("Application started")
	log.WithFields(map[string]interface{}{
		"version": "1.0.0",
		"env":     "development",
	}).Info("Application configuration loaded")

	fmt.Println("\n2. Logging ke access.log (HTTP access):")
	log.LogToAccessFile("GET", "/api/users", "192.168.1.100", "curl/7.68.0", 200, 50*time.Millisecond, "req-123")
	log.LogToAccessFile("POST", "/api/auth/login", "192.168.1.101", "Mozilla/5.0", 401, 25*time.Millisecond, "req-124")

	fmt.Println("\n3. Logging ke error.log (errors only):")
	log.LogToErrorFile(fmt.Errorf("database connection failed"), map[string]interface{}{
		"host":     "localhost",
		"port":     5432,
		"database": "myapp",
	})

	log.LogToErrorFile(fmt.Errorf("user not found"), map[string]interface{}{
		"user_id": 12345,
		"action":  "get_profile",
	})

	fmt.Println("\n=== Log Files Created ===")

	// Show file structure
	fmt.Println("📁 Expected file structure:")
	fmt.Println("logs/")
	fmt.Println("├── app.log          # Main application logs")
	fmt.Println("├── error.log        # Error-only logs")
	fmt.Println("└── access.log       # HTTP access logs")

	fmt.Println("\n📋 With rotation, you'll see:")
	fmt.Println("logs/")
	fmt.Println("├── app.log                 # Current active log")
	fmt.Println("├── app.log.2025-07-13      # Yesterday's app log")
	fmt.Println("├── error.log               # Current error log")
	fmt.Println("├── error.log.2025-07-13    # Yesterday's errors")
	fmt.Println("├── access.log              # Current access log")
	fmt.Println("└── access.log.2025-07-13   # Yesterday's access")

	fmt.Println("\n✅ Best practices implemented:")
	fmt.Println("• Descriptive file names (app.log, error.log, access.log)")
	fmt.Println("• Date-based rotation by Lumberjack")
	fmt.Println("• Separate concerns (app, errors, access)")
	fmt.Println("• Configurable log directory")
	fmt.Println("• Compression for old files")
	fmt.Println("• Consistent naming convention")
}
