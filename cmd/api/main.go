package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang-domain-driven-design/internal/application/services"
	"golang-domain-driven-design/internal/application/usecases"
	"golang-domain-driven-design/internal/infrastructure/config"
	"golang-domain-driven-design/internal/infrastructure/database"
	"golang-domain-driven-design/internal/infrastructure/repositories"
	"golang-domain-driven-design/internal/interfaces/controllers"
	"golang-domain-driven-design/internal/interfaces/routes"
	"golang-domain-driven-design/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	// Initialize logger
	appLogger := logger.New()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		appLogger.Warn("No .env file found, using environment variables")
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		appLogger.Fatal("Failed to load configuration:", err)
	}

	// Initialize database
	db, err := database.NewDatabase(cfg)
	if err != nil {
		appLogger.Fatal("Failed to initialize database:", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			appLogger.Error("Failed to close database:", err)
		}
	}()

	// Run database migrations
	if err := db.Migrate(); err != nil {
		appLogger.Fatal("Failed to run database migrations:", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db.DB)

	// Initialize use cases
	userUseCase := usecases.NewUserUseCase(userRepo)

	// Initialize services
	jwtService := services.NewJWTService(cfg)

	// Initialize controllers
	healthController := controllers.NewHealthController()
	userController := controllers.NewUserController(userUseCase, jwtService)

	// Setup routes
	router := routes.SetupRoutes(healthController, userController, jwtService)

	// Create server
	server := &http.Server{
		Addr:         cfg.GetServerAddress(),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		appLogger.Info("Server starting on", cfg.GetServerAddress())
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appLogger.Fatal("Failed to start server:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	appLogger.Info("Server shutting down...")

	// Give the server 30 seconds to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		appLogger.Fatal("Server forced to shutdown:", err)
	}

	appLogger.Info("Server exited")
}
