#!/bin/bash

# Development script for golang-domain-driven-design

set -e

echo "Starting development environment..."

# Check if .env file exists
if [ ! -f .env ]; then
    echo "Creating .env file from .env.example..."
    cp .env.example .env
    echo "Please update .env file with your database credentials"
fi

# Install dependencies
echo "Installing dependencies..."
go mod download
go mod tidy

# Run the application
echo "Starting the application..."
go run cmd/api/main.go
