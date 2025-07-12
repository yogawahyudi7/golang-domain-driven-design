# Quick Test Guide

## 🚀 Cara Cepat Testing Aplikasi

### 1. Testing Tanpa Database

Untuk testing cepat tanpa setup database PostgreSQL, ikuti langkah berikut:

#### Setup Environment untuk Testing
```bash
# Copy dan edit .env untuk testing
cp .env.example .env
```

Edit `.env` file dengan konfigurasi sederhana:
```env
# Database Configuration (akan diabaikan jika tidak ada PostgreSQL)
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=golang_domain_driven_design
DB_SSLMODE=disable

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=localhost

# JWT Configuration
JWT_SECRET=test-secret-key-for-development
JWT_EXPIRES_IN=24h

# Application Configuration
APP_ENV=development
APP_NAME=golang-domain-driven-design
```

### 2. Menjalankan dengan Make Alternative

#### Windows PowerShell:
```powershell
# Install dependencies
.\make.ps1 deps

# Build aplikasi
.\make.ps1 build

# Run tests (tidak memerlukan database)
.\make.ps1 test

# Format code
.\make.ps1 fmt
```

#### Windows Command Prompt:
```cmd
# Install dependencies
make.bat deps

# Build aplikasi
make.bat build

# Run tests
make.bat test
```

#### Linux/Mac/WSL dengan Make:
```bash
# Install dependencies
make deps

# Build aplikasi
make build

# Run tests
make test

# Run aplikasi
make run
```

### 3. Testing dengan Docker (Recommended)

Cara termudah untuk testing dengan database:

```bash
# Jalankan aplikasi dengan database PostgreSQL
docker-compose up -d

# Check logs
docker-compose logs -f app

# Test endpoints
curl http://localhost:8080/health

# Stop services
docker-compose down
```

### 4. Testing Manual Commands

```bash
# 1. Install dependencies
go mod download
go mod tidy

# 2. Run tests (tidak perlu database)
go test -v ./internal/domain/...

# 3. Build aplikasi
go build -o build/golang-domain-driven-design cmd/api/main.go

# 4. Format code
go fmt ./...

# 5. Check code
go vet ./...
```

### 5. API Testing

Setelah aplikasi berjalan (port 8080):

```bash
# Health check
curl http://localhost:8080/health

# Atau menggunakan PowerShell
Invoke-RestMethod -Uri "http://localhost:8080/health" -Method GET
```

### 6. Troubleshooting

#### Database Connection Error
Jika muncul error database connection:
1. Pastikan PostgreSQL running
2. Atau gunakan Docker Compose
3. Atau skip database dengan run tests saja

#### Port Already in Use
Jika port 8080 sudah digunakan:
1. Ubah `SERVER_PORT` di `.env`
2. Atau stop aplikasi yang menggunakan port 8080

#### Permission Denied (PowerShell)
Jika error execution policy:
```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

### 7. Development Workflow

```bash
# 1. Install dependencies
make deps  # atau .\make.ps1 deps

# 2. Run tests
make test  # atau .\make.ps1 test

# 3. Format code
make fmt   # atau .\make.ps1 fmt

# 4. Build aplikasi
make build # atau .\make.ps1 build

# 5. Run dengan Docker (full environment)
docker-compose up -d
```
