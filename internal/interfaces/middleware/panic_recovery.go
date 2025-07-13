package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"golang-domain-driven-design/pkg/logger"

	"github.com/gin-gonic/gin"
)

// PanicRecoveryMiddleware returns a gin middleware for handling panics with proper logging
func PanicRecoveryMiddleware(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Get stack trace
				stack := debug.Stack()

				// Get request ID if available
				requestID, exists := c.Get("request_id")
				if !exists {
					requestID = "unknown"
				}

				// Create error details
				errorDetails := map[string]interface{}{
					"request_id":  requestID,
					"method":      c.Request.Method,
					"path":        c.Request.URL.Path,
					"client_ip":   c.ClientIP(),
					"user_agent":  c.Request.UserAgent(),
					"panic_value": fmt.Sprintf("%v", err),
					"stack_trace": string(stack),
				}

				// Log panic to main log
				log.WithFields(errorDetails).Error("Panic recovered")

				// Log panic to error file
				log.LogToErrorFile(fmt.Errorf("panic recovered: %v", err), errorDetails)

				// Log security event (panics might indicate security issues)
				log.LogSecurityEvent(
					"panic_recovered",
					fmt.Sprintf("user_%s", c.ClientIP()), // Use IP as user identifier
					c.ClientIP(),
					fmt.Sprintf("panic_value: %v", err),
				)

				// Log business event for monitoring
				log.LogBusinessEvent(
					"system_panic",
					"system",
					fmt.Sprintf("user_%s", c.ClientIP()), // Use IP as user identifier
					map[string]interface{}{
						"request_id":    requestID,
						"panic_message": fmt.Sprintf("%v", err),
						"endpoint":      c.Request.URL.Path,
						"method":        c.Request.Method,
						"recovery":      "successful",
					},
				)

				// Return JSON error response instead of HTML
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":      "Internal Server Error",
					"message":    "An unexpected error occurred. Please try again later.",
					"request_id": requestID,
				})

				// Abort the request chain
				c.Abort()
			}
		}()

		// Continue processing the request
		c.Next()
	}
}

// PanicRecoveryWithCustomResponse returns a panic recovery middleware with custom response handler
func PanicRecoveryWithCustomResponse(log *logger.Logger, customHandler func(*gin.Context, interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Get stack trace
				stack := debug.Stack()

				// Get request ID if available
				requestID, exists := c.Get("request_id")
				if !exists {
					requestID = "unknown"
				}

				// Create error details
				errorDetails := map[string]interface{}{
					"request_id":  requestID,
					"method":      c.Request.Method,
					"path":        c.Request.URL.Path,
					"client_ip":   c.ClientIP(),
					"user_agent":  c.Request.UserAgent(),
					"panic_value": fmt.Sprintf("%v", err),
					"stack_trace": string(stack),
				}

				// Log panic to main log
				log.WithFields(errorDetails).Error("Panic recovered with custom handler")

				// Log panic to error file
				log.LogToErrorFile(fmt.Errorf("panic recovered: %v", err), errorDetails)

				// Call custom handler if provided
				if customHandler != nil {
					customHandler(c, err)
				} else {
					// Default response
					c.JSON(http.StatusInternalServerError, gin.H{
						"error":      "Internal Server Error",
						"message":    "An unexpected error occurred. Please try again later.",
						"request_id": requestID,
					})
				}

				// Abort the request chain
				c.Abort()
			}
		}()

		// Continue processing the request
		c.Next()
	}
}
