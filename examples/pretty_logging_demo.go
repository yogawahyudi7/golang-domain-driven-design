package main

import (
	"fmt"
	"golang-domain-driven-design/pkg/logger"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== DEMONSTRASI PRETTY JSON LOGGING ===")

	// 1. Pretty Print Enabled (Development)
	fmt.Println("1. Logger dengan Pretty Print ENABLED:")
	prettyLogger := logger.New(&logger.Config{
		Level:       "info",
		Format:      "json",
		Output:      "console",
		PrettyPrint: true, // Pretty formatting
	})

	prettyLogger.WithFields(map[string]interface{}{
		"user_id":    12345,
		"action":     "login",
		"ip_address": "192.168.1.100",
		"success":    true,
	}).Info("User login attempt")

	prettyLogger.LogHTTPRequest("POST", "/api/v1/auth/login", "192.168.1.100", "Mozilla/5.0", 200,
		125*time.Millisecond, "req-123-456")

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// 2. Pretty Print Disabled (Production)
	fmt.Println("2. Logger dengan Pretty Print DISABLED:")
	compactLogger := logger.New(&logger.Config{
		Level:       "info",
		Format:      "json",
		Output:      "console",
		PrettyPrint: false, // Compact formatting
	})

	compactLogger.WithFields(map[string]interface{}{
		"user_id":    12345,
		"action":     "login",
		"ip_address": "192.168.1.100",
		"success":    true,
	}).Info("User login attempt")

	compactLogger.LogHTTPRequest("POST", "/api/v1/auth/login", "192.168.1.100", "Mozilla/5.0", 200,
		125*time.Millisecond, "req-123-456")

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// 3. Business Event Logging
	fmt.Println("3. Business Event dengan Pretty Print:")
	prettyLogger.LogBusinessEvent("payment_processed", "txn_789012", "cust_123", map[string]interface{}{
		"amount":         299.99,
		"currency":       "USD",
		"payment_method": "credit_card",
		"merchant_id":    "merch_456",
	})

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// 4. Security Event Logging
	fmt.Println("4. Security Event dengan Pretty Print:")
	prettyLogger.LogSecurityEvent("failed_login_attempt", "user_123", "192.168.1.200",
		"invalid_password - attempt 3/5")

	fmt.Println("\n=== SELESAI ===")
}
