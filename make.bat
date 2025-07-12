@echo off
REM Batch script untuk Windows Command Prompt
REM Alternatif untuk Makefile

setlocal enabledelayedexpansion

set APP_NAME=golang-domain-driven-design
set BUILD_DIR=build
set MAIN_PATH=cmd/api/main.go

if "%1"=="" goto help
if "%1"=="build" goto build
if "%1"=="run" goto run
if "%1"=="test" goto test
if "%1"=="test-coverage" goto test-coverage
if "%1"=="clean" goto clean
if "%1"=="deps" goto deps
if "%1"=="fmt" goto fmt
if "%1"=="lint" goto lint
if "%1"=="help" goto help
goto unknown

:build
echo Building %APP_NAME%...
if not exist %BUILD_DIR% mkdir %BUILD_DIR%
go build -o %BUILD_DIR%/%APP_NAME%.exe %MAIN_PATH%
echo Build completed successfully!
goto end

:run
echo Running %APP_NAME%...
go run %MAIN_PATH%
goto end

:test
echo Running tests...
go test -v ./...
goto end

:test-coverage
echo Running tests with coverage...
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
echo Coverage report generated: coverage.html
goto end

:clean
echo Cleaning...
if exist %BUILD_DIR% rmdir /s /q %BUILD_DIR%
if exist coverage.out del coverage.out
if exist coverage.html del coverage.html
echo Clean completed!
goto end

:deps
echo Installing dependencies...
go mod download
go mod tidy
echo Dependencies installed!
goto end

:fmt
echo Formatting code...
go fmt ./...
echo Code formatted!
goto end

:lint
echo Linting code...
golangci-lint run 2>nul || echo golangci-lint not installed. Install from: https://golangci-lint.run/usage/install/
goto end

:help
echo Available commands:
echo   build         Build the application
echo   run           Run the application
echo   test          Run tests
echo   test-coverage Run tests with coverage
echo   clean         Clean build artifacts
echo   deps          Install dependencies
echo   fmt           Format the code
echo   lint          Lint the code
echo   help          Show this help message
echo.
echo Usage examples:
echo   make.bat run
echo   make.bat build
echo   make.bat test
goto end

:unknown
echo Unknown command: %1
goto help

:end
