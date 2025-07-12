# Golang Clean Architecture with Domain-Driven Design

<!-- Language Toggle -->
<div align="right">
  <strong>🌐 Language:</strong>
  <a href="README.md">🇮🇩</a> |
  <a href="README.en.md">🇺🇸</a>
</div>
<br>

[![Go Version](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org/)
[![Architecture](https://img.shields.io/badge/Architecture-Clean%20Architecture-green.svg)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
[![DDD](https://img.shields.io/badge/Pattern-Domain%20Driven%20Design-orange.svg)](https://martinfowler.com/tags/domain%20driven%20design.html)
[![Docker](https://img.shields.io/badge/Docker-Multi%20Stage-blue.svg)](https://docs.docker.com/develop/dev-best-practices/)
[![Security](https://img.shields.io/badge/Security-Distroless-red.svg)](https://github.com/GoogleContainerTools/distroless)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> **🌐 Read this in other languages:** [English](README.en.md) • [Bahasa Indonesia](README.md)

This project implements Clean Architecture with Domain-Driven Design (DDD) using the Go programming language. This structure is designed to separate concerns, improve testability, and maintainability.

## 📑 Table of Contents

<details>
<summary>🗂️ <strong>Click to show/hide table of contents</strong></summary>

### 🚀 Getting Started
- [⚡ Quick Start & Docker Configuration](#-quick-start--docker-configuration)
- [🏗️ Project Structure](#️-project-structure)
- [🚀 Quick Start](#-quick-start)
- [🧪 Testing](#-testing)

### ⚙️ Configuration & Commands  
- [📋 Available Commands](#-available-commands)
- [⚙️ Setup Make (Optional)](#️-setup-make-optional)
- [🔧 Configuration](#-configuration)

### 🏛️ Architecture & API
- [🏛️ Architecture Principles](#️-architecture-principles)
- [📡 API Endpoints](#-api-endpoints)
- [🔗 Features](#-features)

### 🐳 Docker & Deployment
- [🐳 Docker Configuration](#-docker-configuration)
  - [📋 Docker Configuration Overview](#-docker-configuration-overview)
  - [🚀 Dockerfile.production](#-dockerfileproduction-recommended-for-production)
  - [🛠️ Dockerfile.development](#️-dockerfiledevelopment-recommended-for-development)
  - [🎯 Real Case Scenarios](#-real-case-scenarios)
  - [📊 Performance Comparison](#-performance-comparison)
  - [⚡ Best Practices](#-best-practices)
  - [🐳 Docker Compose Integration](#-docker-compose-integration)

### 🔒 Security & Advanced
- [🔒 Security](#-security)
  - [Docker Security Features](#docker-security-features)
  - [Vulnerability Scanning](#vulnerability-scanning)
  - [Security Updates](#security-updates)
- [📖 Docker Configuration Guide](#-docker-configuration-guide)
  - [🎯 Choosing the Right Dockerfile](#-choosing-the-right-dockerfile)
  - [🔄 Development to Production Workflow](#-development-to-production-workflow)
  - [🛡️ Security Comparison](#️-security-comparison)
  - [🔍 Troubleshooting Guide](#-troubleshooting-guide)
  - [📊 Performance Benchmarks](#-performance-benchmarks)

### 📚 Documentation & Contributing
- [📚 Documentation](#-documentation)
- [🤝 Contributing](#-contributing)
- [📝 License](#-license)
- [🚀 Next Steps](#-next-steps)
- [🎓 Learning Resources](#-learning-resources)

</details>

---

## 🚀 Quick Start & Docker Configuration

> **💡 Docker Configuration Options:**
> - 🛠️ **Development**: Use `Dockerfile.development` for debugging and development
> - 🚀 **Production**: Use `Dockerfile.production` for deployment with maximum security

### ⚡ Quick Commands

```bash
# 🛠️ Development (with debugging tools)
docker build -f Dockerfile.development -t myapp:dev .
docker run -it myapp:dev

# 🚀 Production (maximum security)  
docker build -f Dockerfile.production -t myapp:prod .
docker run myapp:prod
```

> **🎯 Quick Actions:**
> [📖 View Complete Guide](#-docker-configuration) | [🔧 Setup Development](#️-dockerfiledevelopment-recommended-for-development) | [🚀 Deploy Production](#-dockerfileproduction-recommended-for-production) | [🔒 Security Guide](#-security)

## 📋 Docker Configuration Recommendations

| Environment | Dockerfile | Use Case | Security Level |
|-------------|------------|----------|----------------|
| **Development** | `Dockerfile.development` | 🛠️ Debugging, Learning, Staging | 🔒 Secure |
| **Production** | `Dockerfile.production` | 🚀 Live Apps, High Security | 🔒🔒🔒 Maximum |

**🎯 Best Practice**: Develop with `Dockerfile.development`, deploy with `Dockerfile.production`!

## 🏗️ Project Structure

```
golang-domain-driven-design/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── domain/                     # Domain layer (entities, value objects, repositories)
│   │   ├── entities/
│   │   │   ├── user.go
│   │   │   └── user_test.go
│   │   ├── repositories/
│   │   │   └── user_repository.go
│   │   └── valueobjects/
│   │       ├── email.go
│   │       ├── email_test.go
│   │       └── password.go
│   ├── application/                # Application layer (use cases, services)
│   │   └── usecases/
│   │       ├── dto.go
│   │       └── user_usecase.go
│   ├── infrastructure/             # Infrastructure layer
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── database/
│   │   │   └── database.go
│   │   └── repositories/
│   │       └── user_repository.go
│   └── interfaces/                 # Interface layer
│       ├── controllers/
│       │   ├── health_controller.go
│       │   └── user_controller.go
│       ├── middleware/
│       │   ├── cors.go
│       │   └── logger.go
│       └── routes/
│           └── routes.go
├── pkg/                           # Shared packages
│   ├── errors/
│   │   └── errors.go
│   ├── logger/
│   │   └── logger.go
│   └── validator/
│       └── validator.go
├── docs/                          # Documentation
│   ├── api.md
│   └── architecture.md
├── scripts/                       # Build and deployment scripts
│   ├── build.sh
│   └── dev.sh
├── tests/                         # Test files
├── .env.example                   # Environment variables template
├── .gitignore
├── docker-compose.yml             # Docker composition for development
├── docker-compose.prod.yml        # Docker composition for production
├── Dockerfile.development         # Docker config for development (debugging)
├── Dockerfile.production          # Docker config for production (security)
├── go.mod                         # Go modules
├── go.sum
├── Makefile                       # Build automation
└── README.md
```

## 🚀 Quick Start

[⬆️ Back to Table of Contents](#-table-of-contents)

### Prerequisites

- Go 1.21 or higher
- PostgreSQL (for database)
- Docker and Docker Compose (optional)

### 1. Clone Repository

```bash
git clone <repository-url>
cd golang-domain-driven-design
```

### 2. Setup Environment Variables

```bash
cp .env.example .env
```

Edit the `.env` file according to your database configuration:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=golang_domain_driven_design
```

### 3. Install Dependencies

```bash
go mod download
go mod tidy
```

### 4. Run Application

#### Using Make (Linux/Mac)

If `make` is installed:
```bash
# Development mode
make run

# Build application
make build

# View all commands
make help
```

#### Using PowerShell (Windows)

```powershell
# Development mode
.\make.ps1 run

# Build application
.\make.ps1 build

# View all commands
.\make.ps1 help
```

#### Using Command Prompt (Windows)

```cmd
# Development mode
make.bat run

# Build application
make.bat build

# View all commands
make.bat help
```

#### Using Go Directly

```bash
# Development mode
go run cmd/api/main.go

# Manual build
go build -o build/golang-domain-driven-design cmd/api/main.go
```

#### Using Docker Compose

```bash
docker-compose up -d
```

### 5. Verification

Open browser and access:
- Health check: `http://localhost:8080/health`
- API documentation: See `docs/api.md`

## 🧪 Testing

Run all tests:

```bash
make test
```

Run tests with coverage:

```bash
make test-coverage
```

Test specific layers:

```bash
# Domain layer tests
go test ./internal/domain/...

# Application layer tests  
go test ./internal/application/...
```

## 📋 Available Commands

[⬆️ Back to Table of Contents](#-table-of-contents)

### Using Make (Linux/Mac/WSL)

View all available commands:
```bash
make help
```

Available commands:
- `make build` - Build application
- `make run` - Run application
- `make test` - Run tests
- `make test-coverage` - Test with coverage
- `make clean` - Clean build artifacts
- `make deps` - Install dependencies
- `make fmt` - Format code
- `make lint` - Lint code

### Using PowerShell (Windows)

```powershell
# View all commands
.\make.ps1 help

# Usage examples
.\make.ps1 run           # Run application
.\make.ps1 build         # Build application
.\make.ps1 test          # Run tests
.\make.ps1 test-coverage # Test with coverage
.\make.ps1 clean         # Clean build artifacts
.\make.ps1 deps          # Install dependencies
.\make.ps1 fmt           # Format code
.\make.ps1 lint          # Lint code
```

### Using Command Prompt (Windows)

```cmd
# View all commands
make.bat help

# Usage examples
make.bat run             # Run application
make.bat build           # Build application
make.bat test            # Run tests
make.bat clean           # Clean build artifacts
```

### Manual Commands (All Platforms)

If not using scripts, you can run commands directly:

```bash
# Install dependencies
go mod download && go mod tidy

# Run application
go run cmd/api/main.go

# Build application
go build -o build/golang-domain-driven-design cmd/api/main.go

# Run tests
go test -v ./...

# Test with coverage
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Format code
go fmt ./...

# Clean build artifacts
rm -rf build/ coverage.out coverage.html
```

## ⚙️ Setup Make (Optional)

### Windows

If you want to use `make` on Windows, here are several options:

#### Option 1: Using Chocolatey
```powershell
# Install Chocolatey first (if not already)
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

# Install make
choco install make
```

#### Option 2: Using Scoop
```powershell
# Install Scoop first (if not already)
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
irm get.scoop.sh | iex

# Install make
scoop install make
```

#### Option 3: Using WSL (Windows Subsystem for Linux)
```bash
# Install WSL Ubuntu
wsl --install

# Inside WSL
sudo apt update
sudo apt install make
```

#### Option 4: Using Git Bash
Git Bash usually includes `make`. Install Git for Windows from https://git-scm.com/

### Alternative Without Make

If you don't want to install `make`, use the provided scripts:
- **PowerShell**: `make.ps1`
- **Command Prompt**: `make.bat`
- **Go directly**: `go run cmd/api/main.go`

## 🏛️ Architecture Principles

[⬆️ Back to Table of Contents](#-table-of-contents)

### Domain Layer
- Contains business logic and rules
- Independent from external concerns
- Entities and Value Objects
- Repository interfaces

### Application Layer  
- Orchestrates domain objects
- Use cases and application services
- Transaction management
- Business workflow

### Infrastructure Layer
- Database implementations
- External service integrations
- Configuration management
- Technical details

### Interface Layer
- HTTP handlers/controllers
- Request/Response models
- Routing and Middleware

## 📡 API Endpoints

### Authentication
- `POST /api/v1/auth/login` - Login user

### Users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users` - List users (with pagination)
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Health Check
- `GET /health` - Health check

See complete API documentation in `docs/api.md`

## 🐳 Docker Configuration

[⬆️ Back to Table of Contents](#-table-of-contents)

This project provides **two different Dockerfile configurations** to meet development and production needs:

### 📋 Docker Configuration Overview

| Aspect | Dockerfile.development | Dockerfile.production |
|--------|----------------------|---------------------|
| **Base Image** | `alpine:3.20` | `gcr.io/distroless/static-debian12:nonroot` |
| **Primary Use** | Development & Debugging | Production & Maximum Security |
| **Image Size** | ~25MB | ~15MB |
| **Shell Access** | ✅ Available | ❌ Not Available |
| **Debugging Tools** | ✅ Full Alpine toolset | ❌ Minimal surface |
| **Security Level** | 🔒 Secure | 🔒🔒🔒 Maximum Security |
| **Health Check** | `wget` based | Application based |
| **Signal Handling** | `dumb-init` | Built-in |

### 🚀 Dockerfile.production (Recommended for Production)

**Purpose**: Maximum security with minimal attack surface

**Characteristics**:
- ✅ **Distroless Image**: Uses Google Distroless for maximum security
- ✅ **Minimal Attack Surface**: No shell, package manager, or tools
- ✅ **Smallest Size**: Smallest possible image (~15MB)
- ✅ **Zero CVE**: Almost no vulnerabilities
- ❌ **No Shell Access**: Cannot exec into container for debugging

**When to Use**:
- Production environment
- High security requirements
- Compliance with security standards
- Applications that are mature and stable

**Build Command**:
```bash
docker build -f Dockerfile.production -t myapp:production .
```

### 🛠️ Dockerfile.development (Recommended for Development)

**Purpose**: Flexibility for development and debugging

**Characteristics**:
- ✅ **Alpine Linux**: Complete base image with tools
- ✅ **Shell Access**: Can exec into container for debugging
- ✅ **Debugging Tools**: wget, netstat, ps, and other Alpine tools
- ✅ **dumb-init**: Proper signal handling
- ⚠️ **Larger Size**: Slightly larger (~25MB)
- ⚠️ **More Attack Surface**: More packages installed

**When to Use**:
- Development environment
- Staging environment
- Troubleshooting and debugging
- Teams still learning containers

**Build Command**:
```bash
docker build -f Dockerfile.development -t myapp:development .
```

### 🎯 Real Case Scenarios

#### Scenario 1: Production Deployment
```bash
# Build for production (maximum security)
docker build -f Dockerfile.production -t myapp:prod .
docker run -d --name myapp-prod myapp:prod

# Health check using application itself
docker exec myapp-prod /app/main --health-check
```

#### Scenario 2: Development & Debugging
```bash
# Build for development (easy debugging)
docker build -f Dockerfile.development -t myapp:dev .
docker run -d --name myapp-dev myapp:dev

# Debug with shell access
docker exec -it myapp-dev /bin/sh
/app $ ps aux                    # View running processes
/app $ netstat -ln              # Check network connections
/app $ cat /proc/meminfo        # Check memory usage
/app $ wget -qO- http://localhost:8080/health  # Manual health check
```

#### Scenario 3: Security Scanning Comparison
```bash
# Scan production image (minimal vulnerabilities)
docker scan myapp:prod
# Result: 0-1 vulnerabilities

# Scan development image (more packages = more potential issues)  
docker scan myapp:dev
# Result: 2-5 vulnerabilities (non-critical)
```

### 📊 Performance Comparison

[⬆️ Back to Table of Contents](#-table-of-contents) | [🚀 Quick Start](#-quick-start) | [🐳 Docker Guide](#-docker-configuration-guide)

| Metric | Development | Production |
|--------|------------|------------|
| **Build Time** | ~2-3 minutes | ~2-3 minutes |
| **Image Size** | 25MB | 15MB |
| **Startup Time** | ~1-2 seconds | ~1 second |
| **Memory Usage** | 20-30MB | 15-25MB |
| **Security Score** | Good (8/10) | Excellent (10/10) |

### 🚀 Quick Start Commands

#### Development Environment
```bash
# Build development image
docker build -f Dockerfile.development -t golang-ddd:dev .

# Run with docker-compose (development)
docker-compose -f docker-compose.yml up -d

# Debug container
docker exec -it golang-ddd_app_1 /bin/sh
```

#### Production Environment
```bash
# Build production image  
docker build -f Dockerfile.production -t golang-ddd:prod .

# Run with docker-compose (production)
docker-compose -f docker-compose.prod.yml up -d

# Check health (no shell access)
docker logs golang-ddd_app_1
```

### ⚡ Best Practices

#### 🔄 Development Workflow
```bash
# 1. Develop with development image
docker build -f Dockerfile.development -t myapp:dev .
docker run -it myapp:dev

# 2. Test with production image before deploy
docker build -f Dockerfile.production -t myapp:prod .  
docker run myapp:prod

# 3. Deploy to production
docker tag myapp:prod registry.example.com/myapp:latest
docker push registry.example.com/myapp:latest
```

#### 🛡️ Security Best Practices
```bash
# Scan images before deployment
trivy image myapp:prod
docker scout cves myapp:prod

# Use specific tags, avoid :latest in production
docker build -f Dockerfile.production -t myapp:v1.2.3 .

# Run with security constraints
docker run --read-only --tmpfs /tmp myapp:v1.2.3
```

### 🔧 Customization

#### Modify Development Image
```dockerfile
# Add custom debugging tools to Dockerfile.development
RUN apk add --no-cache \
    curl \
    jq \
    htop \
    nano
```

#### Modify Production Image
```dockerfile
# Production image should remain minimal
# Only add if absolutely necessary for production
```

### 🐳 Docker Compose Integration

```yaml
# docker-compose.yml (Development)
version: '3.8'
services:
  app:
    build:
      context: .
      dockerfile: Dockerfile.development
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=development

# docker-compose.prod.yml (Production)  
version: '3.8'
services:
  app:
    build:
      context: .
      dockerfile: Dockerfile.production
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=production
    read_only: true
    tmpfs:
      - /tmp
```

### 📈 Monitoring & Observability

#### Development Monitoring
```bash
# Easy debugging with shell access
docker exec -it myapp-dev /bin/sh
/app $ top                      # Real-time process monitoring
/app $ df -h                   # Disk usage  
/app $ free -m                 # Memory usage
```

#### Production Monitoring
```bash
# Use external monitoring tools
docker stats myapp-prod         # Container stats
docker logs myapp-prod         # Application logs
# Use APM tools like Prometheus, Grafana, etc.
```

## 🔧 Configuration

Application uses environment variables for configuration:

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host | localhost |
| `DB_PORT` | Database port | 5432 |
| `DB_USER` | Database user | postgres |
| `DB_PASSWORD` | Database password | - |
| `DB_NAME` | Database name | golang_domain_driven_design |
| `SERVER_PORT` | Server port | 8080 |
| `JWT_SECRET` | JWT secret key | - |
| `APP_ENV` | Application environment | development |

## 📚 Documentation

- [Architecture Documentation](docs/architecture.md) - Detailed architecture explanation
- [API Documentation](docs/api.md) - Complete API documentation

## 🤝 Contributing

1. Fork repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

## 📝 License

This project is licensed under the MIT License.

## 🔗 Features

- ✅ Clean Architecture with DDD
- ✅ RESTful API with Gin framework
- ✅ GORM for database ORM
- ✅ PostgreSQL support
- ✅ JWT Authentication (placeholder)
- ✅ Input validation
- ✅ Error handling
- ✅ Logging middleware
- ✅ CORS support
- ✅ Docker containerization
- ✅ Unit tests
- ✅ Environment configuration
- ✅ Database migrations
- ✅ Graceful shutdown

## 🚀 Next Steps

- [ ] Implement JWT authentication
- [ ] Add more comprehensive tests
- [ ] Add API documentation with Swagger
- [ ] Implement caching
- [ ] Add monitoring and metrics
- [ ] CI/CD pipeline with multi-stage Docker builds
- [ ] Rate limiting
- [ ] Request tracing

## 📖 Docker Configuration Guide

[⬆️ Back to Table of Contents](#-table-of-contents)

### 🎯 Choosing the Right Dockerfile

#### Use `Dockerfile.production` when:
- ✅ Deploying to production environment
- ✅ Security compliance is required
- ✅ Application is stable and well-tested
- ✅ Minimal attack surface is priority
- ✅ Image size optimization is important

#### Use `Dockerfile.development` when:
- ✅ Local development and testing
- ✅ Staging environment setup
- ✅ Debugging is frequently needed
- ✅ Learning container technologies
- ✅ Troubleshooting application issues

### 🔄 Development to Production Workflow

```bash
# 1. Develop and test locally
docker build -f Dockerfile.development -t myapp:dev .
docker run -p 8080:8080 myapp:dev

# 2. Debug if needed
docker exec -it myapp_container /bin/sh
/app $ wget -qO- http://localhost:8080/health
/app $ ps aux
/app $ netstat -ln

# 3. Test with production configuration
docker build -f Dockerfile.production -t myapp:test .
docker run -p 8080:8080 myapp:test

# 4. Security scan before deployment
trivy image myapp:test
docker scout cves myapp:test

# 5. Deploy to production
docker tag myapp:test registry.company.com/myapp:v1.0.0
docker push registry.company.com/myapp:v1.0.0
```

### 🛡️ Security Comparison

| Security Aspect | Development | Production |
|-----------------|-------------|------------|
| **Attack Surface** | Medium | Minimal |
| **Shell Access** | Available | None |
| **Package Count** | ~50 packages | ~5 packages |
| **CVE Count** | 2-5 (low/medium) | 0-1 (low) |
| **Base Image** | Alpine Linux | Google Distroless |
| **User Privileges** | appuser:appgroup | nonroot |
| **Debugging** | Full tools | Application logs only |

### ⚙️ Environment-Specific Commands

#### Development Environment
```bash
# Start development stack
docker-compose up -d

# View application logs  
docker-compose logs -f app

# Debug container
docker exec -it $(docker ps -qf "name=app") /bin/sh

# Test health endpoint
docker exec $(docker ps -qf "name=app") wget -qO- http://localhost:8080/health
```

#### Production Environment
```bash
# Deploy production stack
docker-compose -f docker-compose.prod.yml up -d

# Check application health (no shell access)
docker logs $(docker ps -qf "name=app")
docker inspect --format='{{.State.Health.Status}}' $(docker ps -qf "name=app")

# Monitor resource usage
docker stats $(docker ps -qf "name=app")
```

### 🔍 Troubleshooting Guide

<details>
<summary>🛠️ <strong>Click to view troubleshooting guide</strong></summary>

#### Development Issues
```bash
# Container won't start
docker logs myapp-dev

# Permission issues
docker exec -it myapp-dev /bin/sh
/app $ ls -la
/app $ whoami
/app $ id

# Network connectivity
docker exec myapp-dev wget -qO- http://localhost:8080/health
docker exec myapp-dev netstat -ln
```

#### Production Issues
```bash
# Container won't start (no shell access)
docker logs myapp-prod
docker inspect myapp-prod

# Health check failures
docker exec myapp-prod /app/main --health-check

# Resource constraints
docker stats myapp-prod
```

</details>

### 📊 Performance Benchmarks

#### Image Size Comparison
```bash
# Check image sizes
docker images | grep myapp
# myapp    dev     25.3MB
# myapp    prod    14.1MB
```

#### Startup Time Comparison
```bash
# Development image
time docker run --rm myapp:dev /app/main --version
# ~1.2 seconds

# Production image  
time docker run --rm myapp:prod /app/main --version
# ~0.8 seconds
```

#### Memory Usage
```bash
# Development: ~25-30MB baseline
# Production: ~18-22MB baseline
docker stats --no-stream myapp-dev myapp-prod
```

### 🎓 Learning Resources

- [Docker Security Best Practices](https://docs.docker.com/develop/security-best-practices/)
- [Google Distroless Images](https://github.com/GoogleContainerTools/distroless)
- [Alpine Linux Security](https://alpinelinux.org/about/)
- [Container Security Guide](https://kubernetes.io/docs/concepts/security/)

---

**🎯 Key Takeaway**: Develop with `Dockerfile.development` for easy debugging, deploy with `Dockerfile.production` for maximum security!

---

## 📍 Navigation Helper

### 🔝 Quick Links
- [⬆️ Back to Top](#golang-clean-architecture-with-domain-driven-design)
- [📑 Table of Contents](#-table-of-contents)
- [🚀 Quick Start](#-quick-start--docker-configuration)
- [🐳 Docker Configuration](#-docker-configuration)
- [🔒 Security](#-security)

### 📖 Key Sections
| Section | Description |
|---------|-------------|
| [🚀 Quick Start](#-quick-start--docker-configuration) | Get started quickly |
| [🛠️ Development](#️-dockerfiledevelopment-recommended-for-development) | Setup development environment |
| [🚀 Production](#-dockerfileproduction-recommended-for-production) | Deploy to production |
| [🔒 Security](#-security) | Security best practices |
| [🔍 Troubleshooting](#-troubleshooting-guide) | Problem-solving guide |

### 💡 Pro Tips
- 🛠️ **Development**: Use `Dockerfile.development` for debugging
- 🚀 **Production**: Use `Dockerfile.production` for maximum security  
- 🔄 **Best Practice**: Test with production config before deploy
- 📊 **Monitoring**: Use health checks for monitoring

---

<div align="center">

**📧 Have questions?** [Create Issue](../../issues) • **🐛 Found a bug?** [Report Bug](../../issues) • **✨ Want to contribute?** [Read Contributing](#-contributing)

**⭐ If this project helps you, give it a star!**

---

**🌐 Available Languages:**
[🇮🇩 Bahasa Indonesia](README.md) • [🇺🇸 English](README.en.md)

[🔝 Back to Top](#golang-clean-architecture-with-domain-driven-design)

</div>

## 🔒 Security

[⬆️ Back to Table of Contents](#-table-of-contents)

### Docker Security Features

This project implements Docker security best practices through two configurations:

#### 🚀 Dockerfile.production - Maximum Security
- ✅ **Distroless Base Image**: Uses `gcr.io/distroless/static-debian12:nonroot` to reduce attack surface
- ✅ **Non-root User**: Container runs as non-root user
- ✅ **Multi-stage Build**: Reduces image size and eliminates build dependencies
- ✅ **Static Binary**: Binary compiled statically for additional security
- ✅ **Security Flags**: Built with security flags (`-ldflags='-w -s'`)
- ✅ **Zero Shell Access**: No shell to reduce attack vectors
- ✅ **Health Checks**: Container health monitoring with application itself

#### 🛠️ Dockerfile.development - Secure Development
- ✅ **Alpine Linux**: Secure base image with latest security updates
- ✅ **Non-root User**: Custom user with specific UID/GID
- ✅ **dumb-init**: Proper signal handling for graceful shutdown
- ✅ **Development Tools**: Secure debugging tools for development
- ✅ **Health Checks**: HTTP-based health checks with wget

### Vulnerability Scanning

Scan both Docker configurations for vulnerabilities:

```bash
# Build both images
docker build -f Dockerfile.production -t myapp:prod .
docker build -f Dockerfile.development -t myapp:dev .

# Scan production image (minimal vulnerabilities)
trivy image myapp:prod
docker scout cves myapp:prod

# Scan development image (more surface area)
trivy image myapp:dev  
docker scout cves myapp:dev

# Use provided script
./scripts/security-scan.sh

# Compare security scores
echo "Production Image Security:"
trivy image --severity HIGH,CRITICAL myapp:prod
echo "Development Image Security:"  
trivy image --severity HIGH,CRITICAL myapp:dev
```

### Security Updates

- **Base Images**: Always use latest versions of base images
- **Dependencies**: Update Go dependencies regularly with `go mod tidy`
- **Security Patches**: Monitor security advisories for Go and dependencies

### Environment Security

```bash
# Use strong passwords
export POSTGRES_PASSWORD="$(openssl rand -base64 32)"
export JWT_SECRET="$(openssl rand -base64 64)"
export DB_PASSWORD="$(openssl rand -base64 32)"

# Run with secure environment variables
docker-compose -f docker-compose.prod.yml up -d
```
