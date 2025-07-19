# Performance Benchmarks

This directory contains comprehensive performance benchmarking tools and results for the Golang Clean Architecture with Domain-Driven Design project.

## 📊 Benchmark Types

### 1. Go Benchmark Tests
Located in `/benchmarks/*.go` files, these provide micro-benchmarks for:
- HTTP endpoint performance
- Middleware overhead
- Memory allocation patterns
- Concurrent request handling
- Panic recovery performance

### 2. Load Testing
External tools integration for realistic performance testing:
- **wrk**: HTTP benchmarking tool
- **hey**: Load testing tool  
- **Apache Benchmark (ab)**: Classic benchmarking tool

### 3. Memory Analysis
- Runtime memory usage tracking
- Garbage collection analysis
- Memory leak detection
- Resource consumption monitoring

### 4. Startup Time Measurement
- Application initialization time
- Cold start performance
- Warm-up behavior analysis

### 5. Container Analysis
- Docker image size comparison
- Container startup performance
- Resource usage in containerized environment

## 🚀 Quick Start

### Run All Benchmarks (Linux/Mac)
```bash
# Make script executable
chmod +x scripts/benchmark.sh

# Run comprehensive benchmark suite
./scripts/benchmark.sh
```

### Run All Benchmarks (Windows)
```powershell
# Run comprehensive benchmark suite
.\scripts\benchmark.ps1

# Quick benchmark (skip Docker analysis)
.\scripts\benchmark.ps1 -Quick -SkipDocker
```

### Run Individual Go Benchmarks
```bash
# Run all benchmark tests
go test -bench=. -benchmem ./benchmarks

# Run specific benchmark
go test -bench=BenchmarkHealthEndpoint -benchmem ./benchmarks

# Run with CPU profiling
go test -bench=. -cpuprofile=cpu.prof ./benchmarks

# Run with memory profiling  
go test -bench=. -memprofile=mem.prof ./benchmarks
```

### Manual Load Testing
```bash
# Start the application
go run cmd/api/main.go

# In another terminal, run load tests
# With wrk
wrk -t4 -c100 -d30s http://localhost:8080/health

# With hey
hey -n 100000 -c 200 http://localhost:8080/health

# With Apache Benchmark
ab -n 100000 -c 200 http://localhost:8080/health
```

## 📈 Benchmark Results

Results are automatically saved to `benchmarks/reports/` directory with timestamps.

### Expected Performance Metrics

Based on our testing, here are the expected performance characteristics:

#### Framework Performance
- **Gin Framework**: ~40,000+ requests/second (varies by hardware)
- **Memory Usage**: ~10-20MB baseline (without business logic)
- **Cold Start**: <100ms startup time
- **Container Size**: 15-25MB (depending on Dockerfile choice)

#### Endpoint Performance
- **Health Check**: ~50,000+ req/s
- **JSON Response**: ~45,000+ req/s  
- **With Full Middleware**: ~35,000+ req/s
- **POST with JSON**: ~30,000+ req/s

#### Memory Characteristics
- **Per Request Allocation**: <1KB average
- **GC Frequency**: Low pressure under normal load
- **Memory Growth**: Linear with concurrent connections

## 🛠️ Tools Installation

### Install Load Testing Tools

#### Linux (Ubuntu/Debian)
```bash
# Install wrk
sudo apt-get update
sudo apt-get install wrk

# Install hey
go install github.com/rakyll/hey@latest

# Install Apache Benchmark
sudo apt-get install apache2-utils
```

#### macOS
```bash
# Install wrk
brew install wrk

# Install hey
go install github.com/rakyll/hey@latest

# Install Apache Benchmark (included in httpd)
brew install httpd
```

#### Windows
```powershell
# Install hey (requires Go)
go install github.com/rakyll/hey@latest

# For wrk on Windows, use WSL or Docker
wsl --install
# Then follow Linux instructions in WSL
```

### Install Analysis Tools
```bash
# Install Go profiling tools (included with Go)
go install golang.org/x/tools/cmd/pprof@latest

# Install graphviz for pprof visualization
# Linux: sudo apt-get install graphviz
# macOS: brew install graphviz
# Windows: Download from https://graphviz.org/download/
```

## 📊 Analyzing Results

### View Go Benchmark Results
```bash
# View latest benchmark report
cat benchmarks/reports/benchmark_report_*.md | tail -1

# Compare benchmarks over time
go test -bench=. ./benchmarks > current.txt
benchcmp baseline.txt current.txt
```

### Profile Analysis
```bash
# Analyze CPU profile
go tool pprof cpu.prof
# In pprof: web, top, list

# Analyze memory profile  
go tool pprof mem.prof
# In pprof: web, top, list

# Generate web visualization
go tool pprof -http=:8080 cpu.prof
```

### Load Test Analysis
```bash
# wrk provides detailed latency distribution
wrk -t4 -c100 -d30s --latency http://localhost:8080/health

# hey provides percentile analysis
hey -n 10000 -c 100 -m GET http://localhost:8080/health
```

## 🎯 Performance Targets

### Minimum Acceptable Performance
- **Throughput**: >10,000 req/s (health endpoint)
- **Latency**: <10ms p99 (under normal load)
- **Memory**: <50MB resident (steady state)
- **Startup**: <500ms (cold start)

### Target Performance Goals  
- **Throughput**: >30,000 req/s (health endpoint)
- **Latency**: <5ms p99 (under normal load)
- **Memory**: <30MB resident (steady state)
- **Startup**: <100ms (cold start)

### Exceptional Performance
- **Throughput**: >50,000 req/s (health endpoint)
- **Latency**: <2ms p99 (under normal load)  
- **Memory**: <20MB resident (steady state)
- **Startup**: <50ms (cold start)

## 🔧 Optimization Tips

### Code Optimizations
1. **Use release mode**: `gin.SetMode(gin.ReleaseMode)`
2. **Minimize allocations**: Reuse objects, use object pools
3. **Optimize JSON handling**: Use faster JSON libraries if needed
4. **Database connection pooling**: Tune pool size and timeouts
5. **Reduce middleware overhead**: Only use necessary middleware

### Deployment Optimizations
1. **Build with optimizations**: `-ldflags="-s -w"`
2. **Use production Docker images**: Distroless or Alpine
3. **Configure GC**: Tune `GOGC` environment variable
4. **Resource limits**: Set appropriate CPU and memory limits
5. **Load balancing**: Distribute traffic across instances

### Monitoring in Production
1. **Continuous benchmarking**: Run benchmarks in CI/CD
2. **APM tools**: Use Prometheus, Grafana, or similar
3. **Profile in production**: Regular profiling of live systems
4. **Alert on regressions**: Set up performance regression alerts

## 📝 Contributing

When making changes that might affect performance:

1. **Run benchmarks before and after** your changes
2. **Document performance impact** in pull requests
3. **Add specific benchmarks** for new features
4. **Update expected metrics** if targets change

### Adding New Benchmarks
```go
func BenchmarkNewFeature(b *testing.B) {
    // Setup
    router := setupRouter()
    req := httptest.NewRequest("GET", "/new-endpoint", nil)
    
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            w := httptest.NewRecorder()
            router.ServeHTTP(w, req)
        }
    })
}
```

## 🔗 Related Documentation

- [Architecture Documentation](../docs/architecture.md)
- [API Documentation](../docs/api.md)
- [Docker Configuration](../README.md#docker-configuration)
- [Security Guidelines](../docs/security.md)

## 📞 Support

For performance-related questions or optimization help:
1. Check existing benchmark reports
2. Review profiling data
3. Open an issue with benchmark results
4. Consider discussing in project discussions
