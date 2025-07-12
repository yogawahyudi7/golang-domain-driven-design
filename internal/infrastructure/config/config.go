package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config represents the application configuration
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
	App      AppConfig
}

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Host string
	Port int
}

// JWTConfig represents JWT configuration
type JWTConfig struct {
	Secret    string
	ExpiresIn time.Duration
}

// AppConfig represents application configuration
type AppConfig struct {
	Name string
	Env  string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	serverPort, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid SERVER_PORT: %w", err)
	}

	jwtExpiresIn, err := time.ParseDuration(getEnv("JWT_EXPIRES_IN", "24h"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRES_IN: %w", err)
	}

	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "golang_domain_driven_design"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: serverPort,
		},
		JWT: JWTConfig{
			Secret:    getEnv("JWT_SECRET", "your-secret-key"),
			ExpiresIn: jwtExpiresIn,
		},
		App: AppConfig{
			Name: getEnv("APP_NAME", "golang-domain-driven-design"),
			Env:  getEnv("APP_ENV", "development"),
		},
	}, nil
}

// GetDSN returns the database connection string
func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}

// GetServerAddress returns the server address
func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
