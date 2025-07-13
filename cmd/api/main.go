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
	// Load environment variables first
	if err := godotenv.Load(); err != nil {
		// We can't log this yet as logger needs config
		// But .env file is optional, so we continue
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	// Initialize logger with configuration
	appLogger := logger.New(cfg.Logger)

	appLogger.WithFields(map[string]interface{}{
		"app_name": cfg.App.Name,
		"env":      cfg.App.Env,
		"version":  "1.0.0",
	}).Info("Starting application")

	// Log environment info
	if err := godotenv.Load(); err != nil {
		appLogger.Warn("No .env file found, using environment variables")
	}

	// Initialize database
	appLogger.Info("Initializing database connection")
	db, err := database.NewDatabase(cfg)
	if err != nil {
		appLogger.WithError(err).Fatal("Failed to initialize database")
	}
	defer func() {
		appLogger.Info("Closing database connection")
		if err := db.Close(); err != nil {
			appLogger.WithError(err).Error("Failed to close database")
		}
	}()

	// Run database migrations
	appLogger.Info("Running database migrations")
	if err := db.Migrate(); err != nil {
		appLogger.WithError(err).Fatal("Failed to run database migrations")
	}

	// Initialize repositories
	appLogger.Info("Initializing repositories")
	userRepo := repositories.NewUserRepository(db.DB)

	// Initialize use cases
	appLogger.Info("Initializing use cases")
	userUseCase := usecases.NewUserUseCase(userRepo)

	// Initialize services
	appLogger.Info("Initializing services")
	jwtService := services.NewJWTService(cfg)

	// Initialize controllers
	appLogger.Info("Initializing controllers")
	healthController := controllers.NewHealthController()
	userController := controllers.NewUserController(userUseCase, jwtService)
	testController := controllers.NewTestController()

	// Setup routes
	appLogger.Info("Setting up routes")
	router := routes.SetupRoutes(healthController, userController, testController, jwtService, appLogger)

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
		appLogger.WithField("address", cfg.GetServerAddress()).Info("Server starting")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appLogger.WithError(err).Fatal("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	appLogger.Info("Shutdown signal received, starting graceful shutdown...")

	// Give the server 30 seconds to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		appLogger.WithError(err).Fatal("Server forced to shutdown")
	}

	appLogger.Info("Server exited gracefully")
}
