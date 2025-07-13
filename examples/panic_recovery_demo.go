//go:build example
// +build example

package main

import (
	"fmt"
	"golang-domain-driven-design/pkg/logger"
	"net/http"
	"time"
)

func main() {
	fmt.Println("=== DEMONSTRASI PANIC RECOVERY TESTING ===")

	// 1. Konfigurasi logger
	config := &logger.Config{
		Level:         "info",
		Format:        "json",
		Output:        "file",
		PrettyPrint:   true,
		LogDir:        "logs",
		AppLogFile:    "panic-test.log",
		ErrorLogFile:  "panic-error.log",
		AccessLogFile: "panic-access.log",
		MaxSize:       100,
		MaxBackups:    5,
		MaxAge:        30,
		Compress:      true,
		DateFormat:    "2006-01-02",
	}

	log := logger.New(config)

	fmt.Println("📋 Testing panic recovery endpoints...")
	fmt.Println("🚀 Make sure the API server is running first:")
	fmt.Println("   go run cmd/api/main.go")
	fmt.Println("")

	// Test endpoints yang bisa dicoba
	endpoints := []struct {
		name        string
		url         string
		description string
	}{
		{
			name:        "Health Check",
			url:         "http://localhost:8080/health",
			description: "Normal endpoint - should work",
		},
		{
			name:        "Test Success",
			url:         "http://localhost:8080/test/success",
			description: "Test endpoint - should return success",
		},
		{
			name:        "Test Error",
			url:         "http://localhost:8080/test/error",
			description: "Test error handling (not panic)",
		},
		{
			name:        "Test Panic (String)",
			url:         "http://localhost:8080/test/panic",
			description: "Test panic recovery with string panic",
		},
		{
			name:        "Test Panic (Nil Pointer)",
			url:         "http://localhost:8080/test/panic?type=nil",
			description: "Test panic recovery with nil pointer dereference",
		},
		{
			name:        "Test Panic (Slice Bounds)",
			url:         "http://localhost:8080/test/panic?type=slice",
			description: "Test panic recovery with slice index out of bounds",
		},
		{
			name:        "Test Panic (Nil Map)",
			url:         "http://localhost:8080/test/panic?type=map",
			description: "Test panic recovery with nil map access",
		},
		{
			name:        "Test Panic (Custom)",
			url:         "http://localhost:8080/test/panic?type=custom&message=Custom panic for testing!",
			description: "Test panic recovery with custom message",
		},
	}

	fmt.Println("🔗 Available test endpoints:")
	for i, endpoint := range endpoints {
		fmt.Printf("%d. %s\n", i+1, endpoint.name)
		fmt.Printf("   URL: %s\n", endpoint.url)
		fmt.Printf("   Description: %s\n\n", endpoint.description)
	}

	fmt.Println("🧪 Testing panic endpoints automatically...")

	// Test beberapa endpoints secara otomatis
	testEndpoints := []string{
		"http://localhost:8080/health",
		"http://localhost:8080/test/success",
		"http://localhost:8080/test/error",
		"http://localhost:8080/test/panic",
		"http://localhost:8080/test/panic?type=nil",
		"http://localhost:8080/test/panic?type=custom&message=Automated test panic",
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for i, url := range testEndpoints {
		fmt.Printf("Testing %d: %s\n", i+1, url)
		
		resp, err := client.Get(url)
		if err != nil {
			log.WithError(err).Error("Failed to make request")
			fmt.Printf("❌ Error: %v\n\n", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Printf("✅ Response Status: %d %s\n", resp.StatusCode, resp.Status)
		
		// Log the test
		log.WithFields(map[string]interface{}{
			"test_url":       url,
			"response_code":  resp.StatusCode,
			"response_status": resp.Status,
		}).Info("Panic recovery test completed")
		
		fmt.Println("")
		time.Sleep(500 * time.Millisecond) // Small delay between requests
	}

	fmt.Println("=== HASIL TESTING ===")
	fmt.Println("✅ Jika semua endpoint merespons dengan status code (bahkan 500), berarti panic recovery bekerja")
	fmt.Println("❌ Jika ada endpoint yang tidak merespons atau server crash, panic recovery tidak bekerja")
	fmt.Println("")
	fmt.Println("📁 Check log files:")
	fmt.Println("   logs/panic-test.log      - Main application logs")
	fmt.Println("   logs/panic-error.log     - Error logs including panic details")
	fmt.Println("   logs/panic-access.log    - HTTP access logs")
	fmt.Println("")
	fmt.Println("🔍 What to look for in logs:")
	fmt.Println("   • 'Panic recovered' messages in error logs")
	fmt.Println("   • Stack traces for debugging")
	fmt.Println("   • Security events for panic monitoring")
	fmt.Println("   • Business events for system health tracking")
}
