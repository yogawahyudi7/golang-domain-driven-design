package benchmarks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-domain-driven-design/internal/interfaces/controllers"
	"golang-domain-driven-design/internal/interfaces/middleware"
	"golang-domain-driven-design/pkg/logger"

	"github.com/gin-gonic/gin"
)

func init() {
	// Set Gin to release mode for benchmarks
	gin.SetMode(gin.ReleaseMode)
}

func BenchmarkHealthEndpoint(b *testing.B) {
	router := gin.New()
	healthController := &controllers.HealthController{}
	router.GET("/health", healthController.Check)

	req := httptest.NewRequest("GET", "/health", nil)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}

func BenchmarkHealthEndpointWithMiddleware(b *testing.B) {
	config := logger.DefaultConfig()
	appLogger := logger.New(config)

	router := gin.New()
	router.Use(middleware.LoggerMiddleware(appLogger))
	router.Use(middleware.CORS())
	router.Use(middleware.PanicRecoveryMiddleware(appLogger))

	healthController := &controllers.HealthController{}
	router.GET("/health", healthController.Check)

	req := httptest.NewRequest("GET", "/health", nil)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}

func BenchmarkJSONResponse(b *testing.B) {
	router := gin.New()
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "JSON response benchmark",
			"data": map[string]interface{}{
				"id":    123,
				"name":  "Test User",
				"email": "test@example.com",
			},
		})
	})

	req := httptest.NewRequest("GET", "/json", nil)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}

func BenchmarkPOSTWithJSONBody(b *testing.B) {
	router := gin.New()
	router.POST("/test", func(c *gin.Context) {
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"received": data})
	})

	testData := map[string]interface{}{
		"name":  "Test User",
		"email": "test@example.com",
		"age":   25,
	}
	jsonData, _ := json.Marshal(testData)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := httptest.NewRequest("POST", "/test", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}

func BenchmarkPanicRecoveryMiddleware(b *testing.B) {
	config := logger.DefaultConfig()
	appLogger := logger.New(config)

	router := gin.New()

	// Add only panic recovery middleware for this test
	router.Use(middleware.PanicRecoveryMiddleware(appLogger))

	router.GET("/panic", func(c *gin.Context) {
		// Simulate potential panic scenario but handle it gracefully
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "recovered from panic"})
			}
		}()
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	req := httptest.NewRequest("GET", "/panic", nil)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}

func BenchmarkCompleteStack(b *testing.B) {
	config := logger.DefaultConfig()
	appLogger := logger.New(config)

	router := gin.New()
	router.Use(middleware.LoggerMiddleware(appLogger))
	router.Use(middleware.CORS())
	router.Use(middleware.PanicRecoveryMiddleware(appLogger))

	// Add controllers
	healthController := &controllers.HealthController{}
	testController := &controllers.TestController{}

	// Add routes
	router.GET("/health", healthController.Check)
	router.GET("/test/success", testController.TestSuccess)

	// Test various endpoints
	endpoints := []string{
		"/health",
		"/test/success",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, endpoint := range endpoints {
				req := httptest.NewRequest("GET", endpoint, nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
			}
		}
	})
}
