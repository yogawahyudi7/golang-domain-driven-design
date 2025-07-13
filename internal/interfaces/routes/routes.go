package routes

import (
	"golang-domain-driven-design/internal/application/services"
	"golang-domain-driven-design/internal/interfaces/controllers"
	"golang-domain-driven-design/internal/interfaces/middleware"
	"golang-domain-driven-design/pkg/logger"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(
	healthController *controllers.HealthController,
	userController *controllers.UserController,
	testController *controllers.TestController,
	jwtService services.JWTService,
	appLogger *logger.Logger,
) *gin.Engine {
	// Create gin router
	router := gin.New()

	// Add middleware
	router.Use(middleware.LoggerMiddleware(appLogger))
	router.Use(middleware.CORS())
	router.Use(middleware.PanicRecoveryMiddleware(appLogger)) // Custom panic recovery with logging

	// Health check route
	router.GET("/health", healthController.Check)

	// Test endpoints (for development only)
	test := router.Group("/test")
	{
		test.GET("/panic", testController.TestPanic)
		test.GET("/error", testController.TestError)
		test.GET("/success", testController.TestSuccess)
	}

	// Create JWT middleware instance
	jwtAuth := middleware.NewJWTAuth(jwtService)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Public authentication routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", userController.RegisterUser)
			auth.POST("/login", userController.AuthenticateUser)
			auth.POST("/refresh", userController.RefreshToken)
		}

		// Protected user routes
		users := v1.Group("/users")
		users.Use(jwtAuth.Middleware())
		{
			users.GET("/profile", userController.GetProfile)
			users.POST("", userController.CreateUser)
			users.GET("", userController.ListUsers)
			users.GET("/:id", userController.GetUser)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}
	}

	return router
}
