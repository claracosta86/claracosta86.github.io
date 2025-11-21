# Domain-Driven Design (DDD) Architecture

This document describes the new DDD architecture implemented in the `back` folder.

## Architecture Overview

The application follows a layered DDD architecture with clear separation of concerns:

```
back/
├── application/      # Application Layer - Use cases
│   ├── comment/  # Card Comment Section  
│   │   └── use_case.go      
│   ├── cultural/        # Events/Tourist Attractions 
│   │   └── use_case.go
│   ├── notification/        # Notification
│   │   └── use_case.go         
│   └── user/        # User use cases
│       └── use_case.go      
├── cmd/     
│   └── api/        
│       └── main.go 
├── domain/          # Domain Layer - Core business logic
│   ├── comment/        # Card Comment Section domain
│   │   ├── entity.go        # Entity with business rules
│   │   ├── value_objects.go # Value objects
│   │   ├── repository.go    # Repository interface
│   │   └── service.go       # Domain service
│   ├── cultural/        # Events/Tourist Attractions domain
│   │   ├── entity.go         
│   │   ├── value_objects.go 
│   │   ├── repository.go   
│   │   └── service.go      
│   ├── notification/        # Notification domain
│   │   ├── entity.go         
│   │   ├── value_objects.go 
│   │   ├── repository.go    
│   │   └── service.go       
│   └── user/        # User domain
│       ├── entity.go         
│       ├── value_objects.go 
│       ├── repository.go    
│       └── service.go 
├── infrastructure/   # Infrastructure Layer - External concerns
│   ├── persistence/ # Data persistence
│   │   └── mysql/   # MySQL implementation
│   │       ├── queries 
│   │       │   ├── comment.sql
│   │       │   ├── cultural.sql
│   │       │   ├── notification.sql
│   │       │   └── user.sql
│   │       ├── comment_repository.go
│   │       ├── cultural_repository.go
│   │       ├── notification_repository.go
│   │       └── user_repository.go
│   └── container/   # Dependency injection
│       └── container.go
├── interface/       # Interface Layer - HTTP handlers
│   ├── http/        # HTTP interface
│   │   ├── handlers
│   │   │   ├── comment_handler.go
│   │   │   ├── cultural_handler.go 
│   │   │   ├── notification_handler.go 
│   │   │   └── user_handler.go
│   │   ├── router
│   │   │    └── router.go
│   └── model       # HTTP response/request models
│       ├── comment.go
│       ├── cultural.go
│       ├── notifications.go
│       └── user.go
├── mocks/ # Mocks de funções para os testes
│   ├
│   ├
│   ... 
├── go.mod
└── go.sum
```

## Layer Responsibilities

### 1. Domain Layer (`domain/`)
- **Entities**: Core business objects with identity and lifecycle
- **Value Objects**: Immutable objects without identity
- **Repository Interfaces**: Contracts for data access
- **Domain Services**: Business logic that doesn't belong to entities

### 2. Application Layer (`application/`)
- **Use Cases**: Application-specific business rules
- **DTOs**: Data transfer objects for input/output
- **Orchestration**: Coordinates domain objects and services

### 3. Infrastructure Layer (`infrastructure/`)
- **Persistence**: Database implementations
- **External Services**: Third-party integrations
- **Configuration**: Environment-specific settings
- **Dependency Injection**: Wiring of components

### 4. Interface Layer (`interface/`)
- **HTTP Handlers**: REST API endpoints
- **Request/Response**: HTTP-specific data handling
- **Validation**: Input validation and sanitization

## Key Benefits

1. **Separation of Concerns**: Each layer has a single responsibility
2. **Testability**: Easy to mock dependencies and test in isolation
3. **Maintainability**: Clear structure makes code easier to understand
4. **Scalability**: Easy to add new features without affecting existing code
5. **Domain Focus**: Business logic is centralized and protected

## Implementation Details

### Comment Domain
- **Entity**: `Comment` with business validation rules
- **Value Objects**: 
- **Repository**: Interface for data persistence
- **Service**: Business logic for user operations

### Cultural Domain
- **Entity**: `Cultural` with business validation rules
- **Value Objects**: 
- **Repository**: Interface for data persistence
- **Service**: Business logic for user operations

### Notification Domain
- **Entity**: `Notification` with business validation rules
- **Value Objects**: 
- **Repository**: Interface for data persistence
- **Service**: Business logic for user operations


### User Domain
- **Entity**: `User` with business validation rules
- **Value Objects**: `Email`, `Password`, `Name`, `Document`
- **Repository**: Interface for data persistence
- **Service**: Business logic for user operations

### Dependency Injection
- **Container**: Centralized dependency management
- **Wiring**: Automatic composition of components
- **Lifecycle**: Proper initialization and cleanup

## Usage Example

```go
// Create container with dependencies
container := container.NewContainer(db)

// Use handlers from container
userHandler := container.UserHandler

// Handler automatically uses the correct use case and domain service
```

## Next Steps

1. Add more value objects and business rules
2. Add comprehensive error handling
3. Implement logging and monitoring
4. Add unit tests for each layer

## Migration Notes

The old structure has been preserved in separate files for reference. The new DDD structure provides:
- Better separation of concerns
- More maintainable code
- Easier testing
- Clearer business logic organization
