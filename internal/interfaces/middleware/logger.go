package middleware

import (
	"time"

	"golang-domain-driven-design/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// LoggerMiddleware returns a gin middleware for logging HTTP requests with Logrus
func LoggerMiddleware(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate request ID
		requestID := uuid.New().String()
		c.Set("request_id", requestID)

		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Log the request
		log.LogHTTPRequest(
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			c.Request.UserAgent(),
			c.Writer.Status(),
			latency,
			requestID,
		)

		// Log errors if any
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				log.WithField("request_id", requestID).
					WithError(err.Err).
					Error("Request processing error")
			}
		}
	}
}

// Logger returns the legacy gin middleware for backward compatibility
// Deprecated: Use LoggerMiddleware instead
func Logger() gin.HandlerFunc {
	return gin.Logger()
}
