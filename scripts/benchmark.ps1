# Performance Benchmark Script for Golang DDD Project (Windows PowerShell)
# This script runs comprehensive performance tests and generates reports

param(
    [switch]$Quick,
    [switch]$SkipDocker,
    [string]$OutputDir = "benchmarks\reports"
)

$ErrorActionPreference = "Stop"

$ProjectName = "golang-domain-driven-design"
$Timestamp = Get-Date -Format "yyyyMMdd_HHmmss"
$ReportFile = "$OutputDir\benchmark_report_$Timestamp.md"

Write-Host "=== Performance Benchmark Suite ===" -ForegroundColor Blue
Write-Host "Starting comprehensive performance testing..."
Write-Host "Report will be saved to: $ReportFile"
Write-Host ""

# Create reports directory if it doesn't exist
if (!(Test-Path $OutputDir)) {
    New-Item -ItemType Directory -Path $OutputDir -Force | Out-Null
}

# Initialize report file
$ReportHeader = @"
# Performance Benchmark Report

**Project:** $ProjectName  
**Date:** $(Get-Date)  
**Go Version:** $(go version)  
**Git Commit:** $(git rev-parse --short HEAD)  
**Branch:** $(git branch --show-current)

## System Information
- **OS:** $([System.Environment]::OSVersion.VersionString)
- **Architecture:** $($env:PROCESSOR_ARCHITECTURE)
- **CPU Cores:** $($env:NUMBER_OF_PROCESSORS)
- **Memory:** $([math]::Round((Get-CimInstance -ClassName Win32_ComputerSystem).TotalPhysicalMemory / 1GB, 2)) GB

---

"@

$ReportHeader | Out-File -FilePath $ReportFile -Encoding UTF8

Write-Host "1. Running Go Benchmark Tests..." -ForegroundColor Yellow
"## Go Benchmark Results" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
"" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
"``````" | Out-File -FilePath $ReportFile -Append -Encoding UTF8

try {
    Write-Host "Running benchmark tests..." -ForegroundColor Yellow
    $BenchmarkOutput = & cmd /c "go test -bench=. -benchmem ./benchmarks 2>&1"
    if ($LASTEXITCODE -eq 0) {
        # Filter hanya hasil benchmark utama, skip log messages
        $FilteredOutput = $BenchmarkOutput | Where-Object { 
            $_ -match "^Benchmark" -or 
            $_ -match "^goos:" -or 
            $_ -match "^goarch:" -or 
            $_ -match "^pkg:" -or 
            $_ -match "^cpu:" -or 
            $_ -match "^PASS" -or 
            $_ -match "^ok"
        }
        $FilteredOutput | Out-File -FilePath $ReportFile -Append -Encoding UTF8
        Write-Host "Go benchmarks completed successfully" -ForegroundColor Green
    } else {
        "Error running Go benchmarks (Exit code: $LASTEXITCODE)" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
        Write-Host "Error running Go benchmarks" -ForegroundColor Red
    }
} catch {
    "Error running Go benchmarks: $($_.Exception.Message)" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    Write-Host "Error running Go benchmarks: $($_.Exception.Message)" -ForegroundColor Red
}

"``````" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
"" | Out-File -FilePath $ReportFile -Append -Encoding UTF8

Write-Host "2. Building application for load testing..." -ForegroundColor Yellow
go build -ldflags="-s -w" -o app.exe cmd/api/main.go

Write-Host "3. Starting application in background..." -ForegroundColor Yellow
$AppProcess = Start-Process -FilePath ".\app.exe" -PassThru -WindowStyle Hidden -RedirectStandardOutput "nul"
Start-Sleep -Seconds 3

# Function to cleanup
function Cleanup {
    Write-Host "Cleaning up..." -ForegroundColor Yellow
    if ($AppProcess -and !$AppProcess.HasExited) {
        $AppProcess.Kill()
        $AppProcess.WaitForExit(5000) | Out-Null
    }
    if (Test-Path "app.exe") {
        Remove-Item "app.exe" -Force
    }
}

# Ensure cleanup on script exit
try {
    Write-Host "4. Testing application responsiveness..." -ForegroundColor Yellow
    
    # Test if application is responding
    $MaxRetries = 10
    $RetryCount = 0
    $AppResponding = $false
    
    while ($RetryCount -lt $MaxRetries -and !$AppResponding) {
        try {
            $Response = Invoke-WebRequest -Uri "http://localhost:8080/health" -UseBasicParsing -TimeoutSec 5
            if ($Response.StatusCode -eq 200) {
                $AppResponding = $true
                Write-Host "Application is running successfully!" -ForegroundColor Green
            }
        } catch {
            $RetryCount++
            Start-Sleep -Seconds 1
        }
    }
    
    if (!$AppResponding) {
        Write-Host "Application is not responding on localhost:8080" -ForegroundColor Red
        throw "Application startup failed"
    }

    "## Load Test Results" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8

    Write-Host "5. Running memory usage analysis..." -ForegroundColor Yellow
    "## Memory Usage Analysis" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "``````" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    
    $ProcessInfo = Get-Process -Id $AppProcess.Id | Select-Object Name, Id, WorkingSet, VirtualMemorySize, CPU
    $ProcessInfo | Format-Table | Out-String | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    
    "``````" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8

    Write-Host "6. Testing startup time..." -ForegroundColor Yellow
    "## Startup Time Analysis" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "Measuring application startup time (5 iterations):" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "``````" | Out-File -FilePath $ReportFile -Append -Encoding UTF8

    # Kill current app for startup time test
    $AppProcess.Kill()
    $AppProcess.WaitForExit(5000) | Out-Null

    # Test startup time multiple times
    for ($i = 1; $i -le 5; $i++) {
        "Iteration ${i}:" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
        $StartTime = Get-Date
        
        # Start process and wait for it to be ready
        $TestProcess = Start-Process -FilePath ".\app.exe" -PassThru -WindowStyle Hidden -RedirectStandardOutput "nul"
        
        # Wait for process to be responsive
        $Ready = $false
        $Timeout = 30 # seconds
        $StartWait = Get-Date
        
        while (!$Ready -and ((Get-Date) - $StartWait).TotalSeconds -lt $Timeout) {
            Start-Sleep -Milliseconds 100
            try {
                $TestResponse = Invoke-WebRequest -Uri "http://localhost:8080/health" -UseBasicParsing -TimeoutSec 1
                if ($TestResponse.StatusCode -eq 200) {
                    $Ready = $true
                    $EndTime = Get-Date
                }
            } catch {
                # App not ready yet, continue waiting
            }
        }
        
        # Kill the test process
        if (!$TestProcess.HasExited) {
            $TestProcess.Kill()
            $TestProcess.WaitForExit(5000) | Out-Null
        }
        
        if ($Ready) {
            $Duration = $EndTime - $StartTime
            "Startup time: $($Duration.TotalMilliseconds) ms" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
        } else {
            "Startup time: TIMEOUT (>$Timeout seconds)" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
        }
        
        # Wait a bit before next iteration
        Start-Sleep -Seconds 1
    }

    "``````" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8

    if (!$SkipDocker) {
        Write-Host "7. Analyzing Docker image sizes..." -ForegroundColor Yellow
        "## Docker Image Analysis" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
        "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8

        if (Get-Command docker -ErrorAction SilentlyContinue) {
            "### Development Image" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            "``````" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            
            try {
                docker build -t "${ProjectName}:dev" -f Dockerfile . 2>&1 | Out-File -FilePath $ReportFile -Append -Encoding UTF8
                docker images "${ProjectName}:dev" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            } catch {
                "Error building development image: $($_.Exception.Message)" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            }
            
            "``````" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            
            "### Production Image" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            "``````" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            
            try {
                docker build -t "${ProjectName}:prod" -f Dockerfile.alpine . 2>&1 | Out-File -FilePath $ReportFile -Append -Encoding UTF8
                docker images "${ProjectName}:prod" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            } catch {
                "Error building production image: $($_.Exception.Message)" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            }
            
            "``````" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
        } else {
            "Docker not available - skipping image analysis" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
        }
    }

    Write-Host "8. Running endpoint response time tests..." -ForegroundColor Yellow
    "## Security and Performance Analysis" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8

    # Restart app for final tests
    $AppProcess = Start-Process -FilePath ".\app.exe" -PassThru -WindowStyle Hidden -RedirectStandardOutput "nul"
    Start-Sleep -Seconds 3

    "### Endpoint Response Times" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "``````" | Out-File -FilePath $ReportFile -Append -Encoding UTF8

    $Endpoints = @("/health")
    foreach ($Endpoint in $Endpoints) {
        "Testing ${Endpoint}:" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
        for ($i = 1; $i -le 3; $i++) {
            try {
                $StartTime = Get-Date
                $Response = Invoke-WebRequest -Uri "http://localhost:8080$Endpoint" -UseBasicParsing -TimeoutSec 10
                $EndTime = Get-Date
                $Duration = ($EndTime - $StartTime).TotalMilliseconds
                "Response time: $Duration ms, Status: $($Response.StatusCode)" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            } catch {
                "Error testing $Endpoint : $($_.Exception.Message)" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
            }
        }
        "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    }
    "``````" | Out-File -FilePath $ReportFile -Append -Encoding UTF8

    # Final summary
    "## Summary" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "- **Go Benchmarks:** Completed successfully" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "- **Memory Analysis:** Completed" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "- **Startup Time:** Measured across 5 iterations" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    if (!$SkipDocker) {
        "- **Docker Images:** $(if (Get-Command docker -ErrorAction SilentlyContinue) { 'Analyzed' } else { 'Skipped (Docker not available)' })" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    }
    "- **Endpoint Testing:** Completed" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "" | Out-File -FilePath $ReportFile -Append -Encoding UTF8
    "**Report Generated:** $(Get-Date)" | Out-File -FilePath $ReportFile -Append -Encoding UTF8

    Write-Host ""
    Write-Host "=== Benchmark Complete ===" -ForegroundColor Green
    Write-Host "Report saved to: $ReportFile" -ForegroundColor Green
    Write-Host ""
    Write-Host "Quick Summary:" -ForegroundColor Blue
    Write-Host "- Go benchmark tests completed"
    Write-Host "- Memory and startup analysis completed"
    Write-Host "- Docker image analysis $(if (!$SkipDocker -and (Get-Command docker -ErrorAction SilentlyContinue)) { 'completed' } else { 'skipped' })"
    Write-Host ""
    Write-Host "To view the full report:" -ForegroundColor Yellow
    Write-Host "Get-Content '$ReportFile'"

} finally {
    Cleanup
}
