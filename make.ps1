# PowerShell Build Script untuk Windows
# Alternatif untuk Makefile

param(
    [Parameter(Position=0)]
    [string]$Command = "help"
)

# Variables
$APP_NAME = "golang-domain-driven-design"
$BUILD_DIR = "build"
$MAIN_PATH = "cmd/api/main.go"

function Write-Info {
    param([string]$Message)
    Write-Host $Message -ForegroundColor Green
}

function Build {
    Write-Info "Building $APP_NAME..."
    if (!(Test-Path $BUILD_DIR)) {
        New-Item -ItemType Directory -Path $BUILD_DIR | Out-Null
    }
    go build -o "$BUILD_DIR/$APP_NAME.exe" $MAIN_PATH
    Write-Info "Build completed successfully!"
}

function Run {
    Write-Info "Running $APP_NAME..."
    go run $MAIN_PATH
}

function Test {
    Write-Info "Running tests..."
    go test -v ./...
}

function Test-Coverage {
    Write-Info "Running tests with coverage..."
    go test -v -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html
    Write-Info "Coverage report generated: coverage.html"
}

function Benchmark {
    Write-Info "Running Go benchmarks..."
    go test -bench=. -benchmem ./benchmarks
}

function BenchmarkFull {
    Write-Info "Running comprehensive benchmark suite..."
    & ".\scripts\benchmark.ps1"
}

function BenchmarkMem {
    Write-Info "Running memory profile benchmark..."
    go test -bench=. -memprofile=mem.prof ./benchmarks
    go tool pprof -http=:8081 mem.prof
}

function BenchmarkCPU {
    Write-Info "Running CPU profile benchmark..."
    go test -bench=. -cpuprofile=cpu.prof ./benchmarks
    go tool pprof -http=:8081 cpu.prof
}

function LoadTest {
    Write-Info "Running load test..."
    if (Get-Command hey -ErrorAction SilentlyContinue) {
        Write-Info "Starting application in background..."
        $AppProcess = Start-Process -FilePath "go" -ArgumentList "run", $MAIN_PATH -PassThru -WindowStyle Hidden
        Start-Sleep -Seconds 3
        
        Write-Info "Running load test with hey..."
        hey -n 10000 -c 100 http://localhost:8080/health
        
        $AppProcess.Kill()
        $AppProcess.WaitForExit(5000) | Out-Null
    } else {
        Write-Host "hey not found. Install with: go install github.com/rakyll/hey@latest" -ForegroundColor Yellow
    }
}

function PerfAnalysis {
    Write-Info "Running complete performance analysis..."
    Benchmark
    LoadTest
    Write-Info "Performance analysis complete"
}

function Clean {
    Write-Info "Cleaning..."
    if (Test-Path $BUILD_DIR) {
        Remove-Item -Recurse -Force $BUILD_DIR
    }
    if (Test-Path "coverage.out") {
        Remove-Item "coverage.out"
    }
    if (Test-Path "coverage.html") {
        Remove-Item "coverage.html"
    }
    Write-Info "Clean completed!"
}

function Install-Dependencies {
    Write-Info "Installing dependencies..."
    go mod download
    go mod tidy
    Write-Info "Dependencies installed!"
}

function Format {
    Write-Info "Formatting code..."
    go fmt ./...
    Write-Info "Code formatted!"
}

function Lint {
    Write-Info "Linting code..."
    if (Get-Command golangci-lint -ErrorAction SilentlyContinue) {
        golangci-lint run
    } else {
        Write-Warning "golangci-lint not installed. Install from: https://golangci-lint.run/usage/install/"
    }
}

function Show-Help {
    Write-Host "Available commands:" -ForegroundColor Cyan
    Write-Host "  build         Build the application" -ForegroundColor White
    Write-Host "  run           Run the application" -ForegroundColor White
    Write-Host "  test          Run tests" -ForegroundColor White
    Write-Host "  test-coverage Run tests with coverage" -ForegroundColor White
    Write-Host "  benchmark     Run Go benchmarks" -ForegroundColor White
    Write-Host "  benchmark-full Run comprehensive benchmark suite" -ForegroundColor White
    Write-Host "  benchmark-mem Run memory profile benchmark" -ForegroundColor White
    Write-Host "  benchmark-cpu Run CPU profile benchmark" -ForegroundColor White
    Write-Host "  load-test     Run load test with hey" -ForegroundColor White
    Write-Host "  perf-analysis Run complete performance analysis" -ForegroundColor White
    Write-Host "  clean         Clean build artifacts" -ForegroundColor White
    Write-Host "  deps          Install dependencies" -ForegroundColor White
    Write-Host "  fmt           Format the code" -ForegroundColor White
    Write-Host "  lint          Lint the code" -ForegroundColor White
    Write-Host "  help          Show this help message" -ForegroundColor White
    Write-Host ""
    Write-Host "Usage examples:" -ForegroundColor Yellow
    Write-Host "  .\make.ps1 run" -ForegroundColor Gray
    Write-Host "  .\make.ps1 build" -ForegroundColor Gray
    Write-Host "  .\make.ps1 test" -ForegroundColor Gray
}

# Execute command
switch ($Command.ToLower()) {
    "build" { Build }
    "run" { Run }
    "test" { Test }
    "test-coverage" { Test-Coverage }
    "benchmark" { Benchmark }
    "benchmark-full" { BenchmarkFull }
    "benchmark-mem" { BenchmarkMem }
    "benchmark-cpu" { BenchmarkCPU }
    "load-test" { LoadTest }
    "perf-analysis" { PerfAnalysis }
    "clean" { Clean }
    "deps" { Install-Dependencies }
    "fmt" { Format }
    "lint" { Lint }
    "help" { Show-Help }
    default { 
        Write-Host "Unknown command: $Command" -ForegroundColor Red
        Show-Help 
    }
}
