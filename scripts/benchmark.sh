#!/bin/bash

# Performance Benchmark Script for Golang DDD Project
# This script runs comprehensive performance tests and generates reports

set -e

PROJECT_NAME="golang-domain-driven-design"
REPORT_DIR="benchmarks/reports"
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
REPORT_FILE="$REPORT_DIR/benchmark_report_$TIMESTAMP.md"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Performance Benchmark Suite ===${NC}"
echo "Starting comprehensive performance testing..."
echo "Report will be saved to: $REPORT_FILE"
echo ""

# Create reports directory if it doesn't exist
mkdir -p "$REPORT_DIR"

# Initialize report file
cat > "$REPORT_FILE" << EOF
# Performance Benchmark Report

**Project:** $PROJECT_NAME  
**Date:** $(date)  
**Go Version:** $(go version)  
**Git Commit:** $(git rev-parse --short HEAD)  
**Branch:** $(git branch --show-current)

## System Information
- **OS:** $(uname -s)
- **Architecture:** $(uname -m)
- **CPU Cores:** $(nproc)
- **Memory:** $(free -h | awk '/^Mem:/ {print $2}')

---

EOF

echo -e "${YELLOW}1. Running Go Benchmark Tests...${NC}"
echo "## Go Benchmark Results" >> "$REPORT_FILE"
echo "" >> "$REPORT_FILE"
echo "\`\`\`" >> "$REPORT_FILE"
go test -bench=. -benchmem ./benchmarks >> "$REPORT_FILE" 2>&1
echo "\`\`\`" >> "$REPORT_FILE"
echo "" >> "$REPORT_FILE"

echo -e "${YELLOW}2. Building application for load testing...${NC}"
go build -ldflags="-s -w" -o app cmd/api/main.go

echo -e "${YELLOW}3. Starting application in background...${NC}"
./app > /dev/null 2>&1 &
APP_PID=$!
sleep 3

# Function to cleanup
cleanup() {
    echo -e "${YELLOW}Cleaning up...${NC}"
    if kill -0 $APP_PID 2>/dev/null; then
        kill $APP_PID
        wait $APP_PID 2>/dev/null || true
    fi
    rm -f app
}

# Trap to ensure cleanup on script exit
trap cleanup EXIT

echo -e "${YELLOW}4. Testing application responsiveness...${NC}"
if ! curl -s http://localhost:8080/health > /dev/null; then
    echo -e "${RED}Application is not responding on localhost:8080${NC}"
    exit 1
fi

echo -e "${GREEN}Application is running successfully!${NC}"

echo -e "${YELLOW}5. Running load tests with wrk (if available)...${NC}"
echo "## Load Test Results" >> "$REPORT_FILE"
echo "" >> "$REPORT_FILE"

if command -v wrk &> /dev/null; then
    echo -e "${GREEN}Running wrk load test...${NC}"
    echo "### wrk Load Test (Health Endpoint)" >> "$REPORT_FILE"
    echo "\`\`\`" >> "$REPORT_FILE"
    wrk -t4 -c100 -d10s http://localhost:8080/health >> "$REPORT_FILE" 2>&1
    echo "\`\`\`" >> "$REPORT_FILE"
    echo "" >> "$REPORT_FILE"
else
    echo -e "${YELLOW}wrk not found, skipping load test${NC}"
    echo "wrk not available - skipping load test" >> "$REPORT_FILE"
    echo "" >> "$REPORT_FILE"
fi

echo -e "${YELLOW}6. Running load tests with hey (if available)...${NC}"
if command -v hey &> /dev/null; then
    echo -e "${GREEN}Running hey load test...${NC}"
    echo "### hey Load Test (Health Endpoint)" >> "$REPORT_FILE"
    echo "\`\`\`" >> "$REPORT_FILE"
    hey -n 10000 -c 100 http://localhost:8080/health >> "$REPORT_FILE" 2>&1
    echo "\`\`\`" >> "$REPORT_FILE"
    echo "" >> "$REPORT_FILE"
else
    echo -e "${YELLOW}hey not found, installing...${NC}"
    if command -v go &> /dev/null; then
        go install github.com/rakyll/hey@latest
        if command -v hey &> /dev/null; then
            echo -e "${GREEN}Running hey load test...${NC}"
            echo "### hey Load Test (Health Endpoint)" >> "$REPORT_FILE"
            echo "\`\`\`" >> "$REPORT_FILE"
            hey -n 10000 -c 100 http://localhost:8080/health >> "$REPORT_FILE" 2>&1
            echo "\`\`\`" >> "$REPORT_FILE"
            echo "" >> "$REPORT_FILE"
        fi
    else
        echo "hey not available - skipping load test" >> "$REPORT_FILE"
        echo "" >> "$REPORT_FILE"
    fi
fi

echo -e "${YELLOW}7. Measuring memory usage...${NC}"
echo "## Memory Usage Analysis" >> "$REPORT_FILE"
echo "" >> "$REPORT_FILE"
echo "\`\`\`" >> "$REPORT_FILE"
ps -o pid,rss,vsz,pcpu,pmem,comm -p $APP_PID >> "$REPORT_FILE"
echo "\`\`\`" >> "$REPORT_FILE"
echo "" >> "$REPORT_FILE"

echo -e "${YELLOW}8. Testing startup time...${NC}"
echo "## Startup Time Analysis" >> "$REPORT_FILE"
echo "" >> "$REPORT_FILE"
echo "Measuring application startup time (5 iterations):" >> "$REPORT_FILE"
echo "\`\`\`" >> "$REPORT_FILE"

# Kill current app
kill $APP_PID
wait $APP_PID 2>/dev/null || true

# Test startup time multiple times
for i in {1..5}; do
    echo "Iteration $i:" >> "$REPORT_FILE"
    { time ./app --quick-exit 2>/dev/null; } 2>&1 | grep real >> "$REPORT_FILE"
done
echo "\`\`\`" >> "$REPORT_FILE"
echo "" >> "$REPORT_FILE"

echo -e "${YELLOW}9. Analyzing Docker image sizes...${NC}"
echo "## Docker Image Analysis" >> "$REPORT_FILE"
echo "" >> "$REPORT_FILE"

if command -v docker &> /dev/null; then
    echo "### Development Image" >> "$REPORT_FILE"
    echo "\`\`\`" >> "$REPORT_FILE"
    docker build -t ${PROJECT_NAME}:dev -f Dockerfile.development . >> "$REPORT_FILE" 2>&1
    docker images ${PROJECT_NAME}:dev >> "$REPORT_FILE"
    echo "\`\`\`" >> "$REPORT_FILE"
    echo "" >> "$REPORT_FILE"
    
    echo "### Production Image" >> "$REPORT_FILE"
    echo "\`\`\`" >> "$REPORT_FILE"
    docker build -t ${PROJECT_NAME}:prod -f Dockerfile.production . >> "$REPORT_FILE" 2>&1
    docker images ${PROJECT_NAME}:prod >> "$REPORT_FILE"
    echo "\`\`\`" >> "$REPORT_FILE"
    echo "" >> "$REPORT_FILE"
else
    echo "Docker not available - skipping image analysis" >> "$REPORT_FILE"
    echo "" >> "$REPORT_FILE"
fi

echo -e "${YELLOW}10. Running security and performance analysis...${NC}"
echo "## Security and Performance Analysis" >> "$REPORT_FILE"
echo "" >> "$REPORT_FILE"

# Test different endpoints
echo "### Endpoint Response Times" >> "$REPORT_FILE"
echo "\`\`\`" >> "$REPORT_FILE"

# Restart app for final tests
./app > /dev/null 2>&1 &
APP_PID=$!
sleep 2

endpoints=("/health" "/test/success" "/test/error")
for endpoint in "${endpoints[@]}"; do
    echo "Testing $endpoint:" >> "$REPORT_FILE"
    for i in {1..3}; do
        curl -s -w "Response time: %{time_total}s, Status: %{http_code}\n" \
             -o /dev/null http://localhost:8080$endpoint >> "$REPORT_FILE" 2>&1
    done
    echo "" >> "$REPORT_FILE"
done
echo "\`\`\`" >> "$REPORT_FILE"

# Final summary
echo "## Summary" >> "$REPORT_FILE"
echo "" >> "$REPORT_FILE"
echo "- **Go Benchmarks:** Completed successfully" >> "$REPORT_FILE"
echo "- **Load Tests:** $(if command -v wrk &> /dev/null || command -v hey &> /dev/null; then echo "Completed"; else echo "Skipped (tools not available)"; fi)" >> "$REPORT_FILE"
echo "- **Memory Analysis:** Completed" >> "$REPORT_FILE"
echo "- **Startup Time:** Measured across 5 iterations" >> "$REPORT_FILE"
echo "- **Docker Images:** $(if command -v docker &> /dev/null; then echo "Analyzed"; else echo "Skipped (Docker not available)"; fi)" >> "$REPORT_FILE"
echo "- **Endpoint Testing:** Completed" >> "$REPORT_FILE"
echo "" >> "$REPORT_FILE"
echo "**Report Generated:** $(date)" >> "$REPORT_FILE"

echo ""
echo -e "${GREEN}=== Benchmark Complete ===${NC}"
echo -e "${GREEN}Report saved to: $REPORT_FILE${NC}"
echo ""
echo -e "${BLUE}Quick Summary:${NC}"
echo "- Go benchmark tests completed"
echo "- Load testing $(if command -v wrk &> /dev/null || command -v hey &> /dev/null; then echo "completed"; else echo "skipped (install wrk or hey for load testing)"; fi)"
echo "- Memory and startup analysis completed"
echo "- Docker image analysis $(if command -v docker &> /dev/null; then echo "completed"; else echo "skipped (Docker not available)"; fi)"
echo ""
echo -e "${YELLOW}To view the full report:${NC}"
echo "cat $REPORT_FILE"
echo ""
echo -e "${YELLOW}To install load testing tools:${NC}"
echo "# Install wrk (Ubuntu/Debian): sudo apt-get install wrk"
echo "# Install hey: go install github.com/rakyll/hey@latest"
