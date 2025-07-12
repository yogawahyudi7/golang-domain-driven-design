#!/bin/bash

# Build script for golang-domain-driven-design

set -e

echo "Building golang-domain-driven-design..."

# Set build variables
APP_NAME="golang-domain-driven-design"
BUILD_DIR="build"
MAIN_PATH="cmd/api/main.go"

# Create build directory
mkdir -p $BUILD_DIR

# Build for different platforms
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}-linux-amd64 $MAIN_PATH

echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}-windows-amd64.exe $MAIN_PATH

echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}-darwin-amd64 $MAIN_PATH

echo "Build completed successfully!"
echo "Binaries are available in the $BUILD_DIR directory."
