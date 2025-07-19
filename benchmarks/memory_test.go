package benchmarks

import (
	"net/http/httptest"
	"runtime"
	"testing"

	"golang-domain-driven-design/internal/interfaces/controllers"
	"golang-domain-driven-design/internal/interfaces/middleware"
	"golang-domain-driven-design/pkg/logger"

	"github.com/gin-gonic/gin"
)

func BenchmarkMemoryUsage(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)

	config := logger.DefaultConfig()
	appLogger := logger.New(config)

	router := gin.New()
	router.Use(middleware.LoggerMiddleware(appLogger))
	router.Use(middleware.CORS())
	router.Use(middleware.PanicRecoveryMiddleware(appLogger))

	healthController := &controllers.HealthController{}
	router.GET("/health", healthController.Check)

	req := httptest.NewRequest("GET", "/health", nil)

	// Force garbage collection before benchmark
	runtime.GC()

	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}

	runtime.ReadMemStats(&m2)

	b.ReportMetric(float64(m2.Alloc-m1.Alloc)/float64(b.N), "B/op")
	b.ReportMetric(float64(m2.Mallocs-m1.Mallocs)/float64(b.N), "allocs/op")
}

func BenchmarkStartupTime(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		// Simulate app startup
		config := logger.DefaultConfig()
		appLogger := logger.New(config)

		router := gin.New()
		router.Use(middleware.LoggerMiddleware(appLogger))
		router.Use(middleware.CORS())
		router.Use(middleware.PanicRecoveryMiddleware(appLogger))

		healthController := &controllers.HealthController{}
		router.GET("/health", healthController.Check)

		b.StartTimer()

		// Simulate a simple request to "warm up" the server
		req := httptest.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}

func BenchmarkConcurrentRequests(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)

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
