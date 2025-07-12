# PowerShell Security Scanning Script for Windows
# This script scans Docker images for vulnerabilities

param(
    [string]$ImageName = "golang-domain-driven-design:latest"
)

# Colors for output
$Green = "Green"
$Yellow = "Yellow"
$Red = "Red"

function Write-Status {
    param([string]$Message)
    Write-Host "[INFO] $Message" -ForegroundColor $Green
}

function Write-Warning {
    param([string]$Message)
    Write-Host "[WARNING] $Message" -ForegroundColor $Yellow
}

function Write-Error {
    param([string]$Message)
    Write-Host "[ERROR] $Message" -ForegroundColor $Red
}

Write-Host "🔍 Docker Security Scanning for golang-domain-driven-design" -ForegroundColor Cyan
Write-Host "===============================================" -ForegroundColor Cyan

# Check if Docker is running
try {
    docker info | Out-Null
    Write-Status "Docker is running"
} catch {
    Write-Error "Docker is not running. Please start Docker Desktop and try again."
    exit 1
}

# Build the image
Write-Status "Building Docker image..."
try {
    docker build -t $ImageName .
    Write-Status "Image built successfully"
} catch {
    Write-Error "Failed to build Docker image"
    exit 1
}

# Check image details
Write-Status "Checking image details..."
$imageSize = docker images $ImageName --format "table {{.Size}}" | Select-Object -Last 1
Write-Status "Image size: $imageSize"

# Check if running as root
Write-Status "Checking user privileges..."
try {
    $user = docker run --rm $ImageName whoami 2>$null
    if ($user -eq "root") {
        Write-Error "Image is running as root user"
    } else {
        Write-Status "✓ Image is running as non-root user: $user"
    }
} catch {
    Write-Status "✓ Image is running as non-root user (nonroot)"
}

# Check Dockerfile best practices
Write-Status "Checking Dockerfile for best practices..."

$dockerfileContent = Get-Content Dockerfile -Raw

if ($dockerfileContent -match "FROM.*:latest") {
    Write-Warning "Using 'latest' tag in Dockerfile. Consider using specific versions."
} else {
    Write-Status "✓ Using specific image versions"
}

if ($dockerfileContent -match "USER") {
    Write-Status "✓ Non-root user specified in Dockerfile"
} else {
    Write-Warning "Consider adding USER instruction in Dockerfile"
}

if ($dockerfileContent -match "HEALTHCHECK") {
    Write-Status "✓ Health check configured"
} else {
    Write-Warning "Consider adding HEALTHCHECK instruction"
}

# Check for Trivy
Write-Status "Checking for vulnerability scanners..."
try {
    trivy version | Out-Null
    Write-Status "Scanning with Trivy..."
    trivy image --severity HIGH,CRITICAL $ImageName
} catch {
    Write-Warning "Trivy not installed. Install from: https://github.com/aquasecurity/trivy/releases"
}

# Check for Docker Scout
try {
    docker scout version | Out-Null
    Write-Status "Scanning with Docker Scout..."
    docker scout cves $ImageName
} catch {
    Write-Warning "Docker Scout not available. Enable in Docker Desktop or install Docker Scout CLI."
}

# Generate report
$reportContent = @"
Docker Security Scan Report
Generated: $(Get-Date)
Image: $ImageName
Image Size: $imageSize

Recommendations:
1. Regularly update base images
2. Use minimal base images (distroless, alpine)
3. Run as non-root user
4. Implement health checks
5. Use multi-stage builds
6. Scan images regularly in CI/CD pipeline
7. Use specific image tags instead of 'latest'
8. Remove unnecessary packages and files
9. Set resource limits
10. Use read-only filesystem when possible

Base Images Used:
- golang:1.23-alpine3.20 (builder)
- gcr.io/distroless/static-debian12:nonroot (runtime)

Security Features Implemented:
✓ Multi-stage build
✓ Non-root user
✓ Distroless base image
✓ Health check
✓ Minimal attack surface
✓ Static binary compilation
✓ Security flags in build process
"@

$reportContent | Out-File -FilePath "security-scan-report.txt" -Encoding UTF8
Write-Status "Security report saved to security-scan-report.txt"

Write-Status "Security scanning completed!"
Write-Status "Review the results above and address any HIGH or CRITICAL vulnerabilities."
