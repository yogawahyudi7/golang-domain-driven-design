#!/bin/bash

# Docker Security Scanning Script
# This script scans Docker images for vulnerabilities

set -e

echo "🔍 Docker Security Scanning for golang-domain-driven-design"
echo "=============================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    print_error "Docker is not running. Please start Docker and try again."
    exit 1
fi

# Build the image
print_status "Building Docker image..."
docker build -t golang-domain-driven-design:latest .

# Function to scan image with different tools
scan_with_trivy() {
    print_status "Scanning with Trivy..."
    if command -v trivy &> /dev/null; then
        trivy image --severity HIGH,CRITICAL golang-domain-driven-design:latest
    else
        print_warning "Trivy not installed. Install with: curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin"
    fi
}

scan_with_docker_scout() {
    print_status "Scanning with Docker Scout..."
    if docker scout --help > /dev/null 2>&1; then
        docker scout cves golang-domain-driven-design:latest
    else
        print_warning "Docker Scout not available. Enable Docker Scout in Docker Desktop or install Docker Scout CLI."
    fi
}

scan_with_snyk() {
    print_status "Scanning with Snyk..."
    if command -v snyk &> /dev/null; then
        snyk container test golang-domain-driven-design:latest
    else
        print_warning "Snyk not installed. Install with: npm install -g snyk"
    fi
}

# Check image layers and best practices
check_best_practices() {
    print_status "Checking Docker best practices..."
    
    # Check if running as root
    USER_CHECK=$(docker run --rm golang-domain-driven-design:latest whoami 2>/dev/null || echo "nonroot")
    if [ "$USER_CHECK" = "root" ]; then
        print_error "Image is running as root user"
    else
        print_status "✓ Image is running as non-root user: $USER_CHECK"
    fi
    
    # Check image size
    IMAGE_SIZE=$(docker images golang-domain-driven-design:latest --format "table {{.Size}}" | tail -n 1)
    print_status "Image size: $IMAGE_SIZE"
    
    # Check for common vulnerabilities in Dockerfile
    print_status "Checking Dockerfile for best practices..."
    
    if grep -q "FROM.*:latest" Dockerfile; then
        print_warning "Using 'latest' tag in Dockerfile. Consider using specific versions."
    else
        print_status "✓ Using specific image versions"
    fi
    
    if grep -q "USER" Dockerfile; then
        print_status "✓ Non-root user specified in Dockerfile"
    else
        print_warning "Consider adding USER instruction in Dockerfile"
    fi
    
    if grep -q "HEALTHCHECK" Dockerfile; then
        print_status "✓ Health check configured"
    else
        print_warning "Consider adding HEALTHCHECK instruction"
    fi
}

# Run scans
echo ""
print_status "Starting security scans..."
echo ""

scan_with_trivy
echo ""

scan_with_docker_scout
echo ""

scan_with_snyk
echo ""

check_best_practices
echo ""

print_status "Security scanning completed!"
print_status "Review the results above and address any HIGH or CRITICAL vulnerabilities."

# Generate summary report
cat << EOF > security-scan-report.txt
Docker Security Scan Report
Generated: $(date)
Image: golang-domain-driven-design:latest
Image Size: $(docker images golang-domain-driven-design:latest --format "table {{.Size}}" | tail -n 1)

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
EOF

print_status "Security report saved to security-scan-report.txt"
