package routes

import (
	"golang-domain-driven-design/internal/application/services"
	"golang-domain-driven-design/internal/interfaces/controllers"
	"golang-domain-driven-design/internal/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(
	healthController *controllers.HealthController,
	userController *controllers.UserController,
	jwtService services.JWTService,
) *gin.Engine {
	// Create gin router
	router := gin.New()

	// Add middleware
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.Use(gin.Recovery())

	// Health check route
	router.GET("/health", healthController.Check)

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
