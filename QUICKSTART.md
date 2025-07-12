# Quick Start Guide

## 🚀 Panduan Cepat Menjalankan Aplikasi

### 1. Persiapan Database

Pastikan PostgreSQL sudah terinstall dan berjalan, kemudian buat database:

```sql
CREATE DATABASE golang_domain_driven_design;
```

### 2. Setup Environment

```bash
# Copy environment template
cp .env.example .env

# Edit .env file dengan konfigurasi database Anda
# Minimal yang perlu diubah:
DB_PASSWORD=your_database_password
```

### 3. Install Dependencies

```bash
go mod download
go mod tidy
```

### 4. Jalankan Aplikasi

```bash
# Menggunakan Makefile
make run

# Atau langsung
go run cmd/api/main.go
```

### 5. Test API

Health Check:
```bash
curl http://localhost:8080/health
```

Create User:
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "username": "johndoe",
    "first_name": "John",
    "last_name": "Doe",
    "password": "SecurePass123!"
  }'
```

Login:
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "password": "SecurePass123!"
  }'
```

### 6. Dengan Docker Compose (Alternatif)

Jika ingin menggunakan Docker:

```bash
# Start semua services (app + database)
docker-compose up -d

# Lihat logs
docker-compose logs -f app

# Stop services
docker-compose down
```

## 🔧 Commands Berguna

```bash
# Build aplikasi
make build

# Jalankan tests
make test

# Format code
make fmt

# Lihat semua commands
make help
```

## 📁 Struktur yang Sudah Dibuat

✅ Domain Layer - Entities, Value Objects, Repository Interfaces
✅ Application Layer - Use Cases, DTOs
✅ Infrastructure Layer - Database, Config, Repository Implementations  
✅ Interface Layer - Controllers, Middleware, Routes
✅ Shared Packages - Errors, Logger, Validator
✅ Tests - Unit tests untuk domain layer
✅ Documentation - API docs & Architecture docs
✅ Docker Support - Dockerfile & docker-compose
✅ Build Scripts - Makefile & shell scripts

## 🎯 Fitur yang Sudah Implementasi

- ✅ User CRUD operations
- ✅ Email validation dengan Value Object
- ✅ Password validation dengan Value Object  
- ✅ Error handling dengan custom error types
- ✅ Database migrations dengan GORM
- ✅ Environment configuration
- ✅ HTTP middleware (CORS, Logging)
- ✅ Graceful shutdown
- ✅ Health check endpoint

## 🔜 Next Steps

1. Implement JWT authentication yang real
2. Tambah validasi input dengan go-playground/validator
3. Tambah logging yang lebih comprehensive
4. Tambah rate limiting
5. Tambah API documentation dengan Swagger
6. Implement caching
7. Tambah monitoring & metrics
