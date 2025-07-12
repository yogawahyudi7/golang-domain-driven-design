# Clean Architecture with Domain-Driven Design

## Overview

This project implements Clean Architecture principles with Domain-Driven Design (DDD) patterns in Go. The architecture promotes separation of concerns, testability, and maintainability.

## Layers

### 1. Domain Layer (`internal/domain/`)
The domain layer contains the business logic and rules. It's the heart of the application and should be independent of external concerns.

#### Entities (`internal/domain/entities/`)
- Represent business objects with identity
- Contain business rules and invariants
- Independent of data persistence concerns

#### Value Objects (`internal/domain/valueobjects/`)
- Immutable objects that describe aspects of the domain
- No identity, compared by value
- Encapsulate validation rules

#### Repository Interfaces (`internal/domain/repositories/`)
- Define contracts for data access
- Part of the domain but implementation is in infrastructure
- Allow dependency inversion

### 2. Application Layer (`internal/application/`)
The application layer orchestrates domain objects to fulfill use cases. It contains application-specific business logic.

#### Use Cases (`internal/application/usecases/`)
- Implement application business rules
- Orchestrate domain objects
- Handle application workflow
- Manage transactions

#### DTOs (`internal/application/usecases/dto.go`)
- Data Transfer Objects for use case inputs/outputs
- Validation rules for input data
- Response formatting

### 3. Infrastructure Layer (`internal/infrastructure/`)
The infrastructure layer provides implementations for external concerns like databases, external services, and configuration.

#### Database (`internal/infrastructure/database/`)
- Database connection and setup
- Migration management
- Health checks

#### Repositories (`internal/infrastructure/repositories/`)
- Concrete implementations of repository interfaces
- Database-specific logic
- Data mapping between domain entities and database models

#### Configuration (`internal/infrastructure/config/`)
- Application configuration management
- Environment variable handling
- Configuration validation

### 4. Interface Layer (`internal/interfaces/`)
The interface layer handles external communication, such as HTTP requests, CLI commands, or message queues.

#### Controllers (`internal/interfaces/controllers/`)
- HTTP request handling
- Input validation and transformation
- Response formatting
- Error handling

#### Middleware (`internal/interfaces/middleware/`)
- Cross-cutting concerns
- Logging, CORS, authentication
- Request/response processing

#### Routes (`internal/interfaces/routes/`)
- URL routing configuration
- Route grouping
- Middleware application

## Shared Packages (`pkg/`)

### Errors (`pkg/errors/`)
- Custom error types
- Error handling utilities
- Consistent error responses

### Logger (`pkg/logger/`)
- Structured logging
- Log level management
- Context-aware logging

### Validator (`pkg/validator/`)
- Input validation utilities
- Custom validation rules
- Error message formatting

## Dependency Flow

```
Interfaces -> Application -> Domain
     ↓
Infrastructure
```

- **Interfaces** depend on **Application**
- **Application** depends on **Domain**
- **Infrastructure** depends on **Domain** (implements interfaces)
- **Domain** has no dependencies (pure business logic)

## Benefits

1. **Testability**: Each layer can be tested independently
2. **Maintainability**: Clear separation of concerns
3. **Flexibility**: Easy to change external dependencies
4. **Scalability**: Clean structure supports team collaboration
5. **Domain Focus**: Business logic is protected and highlighted

## Testing Strategy

1. **Unit Tests**: Test domain entities and value objects
2. **Integration Tests**: Test use cases with mock repositories
3. **End-to-End Tests**: Test complete user journeys
4. **Contract Tests**: Test repository implementations

## Best Practices

1. Keep domain layer pure (no external dependencies)
2. Use dependency injection for loose coupling
3. Implement interfaces in the layer that needs them
4. Keep use cases focused on single responsibilities
5. Use value objects for validation and data integrity
6. Handle errors at appropriate layers
7. Maintain clear boundaries between layers
