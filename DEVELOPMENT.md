# Development Documentation

## Table of Contents

1. [Folder Structure](#folder-structure)
2. [Setup Instructions](#setup-instructions)
3. [Coding Standards](#coding-standards)
4. [Database Schema](#database-schema)
5. [Deployment Guide](#deployment-guide)
6. [Changelog](#changelog)

## Folder Structure

```
go-hexagonal-demo/
├── _cmd/                  # Command-line entry points
├── internal/              # Internal application code
│   ├── core/             # Core business logic
│   │   ├── domains/      # Business entities and rules
│   │   ├── ports/        # Interface definitions
│   │   └── services/     # Business logic implementation
│   └── adapters/         # Interface adapters
│       ├── inbound/      # Input adapters (HTTP, gRPC, etc.)
│       └── outbound/     # Output adapters (Database, External Services)
├── server.go             # Main server file
├── go.mod               # Go module definition
└── go.sum               # Go dependencies checksums
```

## Setup Instructions

### Prerequisites

- Go 1.23.6 or later
- MongoDB
- Elasticsearch 8.x

### Local Development Setup

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd go-hexagonal-demo
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Set up environment variables:

   ```bash
   # Create a .env file with the following variables
   MONGODB_URI=mongodb://localhost:27017
   ELASTICSEARCH_URL=http://localhost:9200
   ```

4. Run the application:
   ```bash
   go run server.go
   ```

### Running Tests

```bash
go test ./...
```

## Coding Standards

### General Guidelines

1. Follow Go's official [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
2. Use meaningful variable and function names
3. Write comments for exported functions and types
4. Keep functions small and focused

### Project-Specific Standards

#### Hexagonal Architecture Principles

1. Core business logic should be independent of external concerns
2. Use interfaces (ports) to define boundaries
3. Implement adapters for external systems
4. Dependency injection for flexible component binding

#### Code Organization

1. Keep domain logic in the `core` package
2. Place interfaces in the `ports` package
3. Implement adapters in the `adapters` package
4. Use dependency injection for service initialization

#### Error Handling

1. Use custom error types for domain-specific errors
2. Wrap errors with context using `fmt.Errorf`
3. Handle errors at the appropriate level

#### Testing

1. Write unit tests for core business logic
2. Use mocks for external dependencies
3. Include integration tests for adapters
4. Maintain high test coverage

## Database Schema

### MongoDB Collections

#### Users Collection

```json
{
  "_id": "ObjectId",
  "email": "string",
  "name": "string",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

### Elasticsearch Indices

#### Search Index

```json
{
  "mappings": {
    "properties": {
      "id": {"type": "keyword"},
      "title": {"type": "text"},
      "content": {"type": "text"},
      "created_at": {"type": "date"}
    }
  }
}
```

## Deployment Guide

### Production Deployment

1. Build the application:

   ```bash
   go build -o app server.go
   ```

2. Set up environment variables:

   ```bash
   export MONGODB_URI=<production-mongodb-uri>
   export ELASTICSEARCH_URL=<production-elasticsearch-url>
   ```

3. Run the application:
   ```bash
   ./app
   ```

### Docker Deployment

1. Build the Docker image:

   ```bash
   docker build -t go-hexagonal-demo .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 \
     -e MONGODB_URI=<mongodb-uri> \
     -e ELASTICSEARCH_URL=<elasticsearch-url> \
     go-hexagonal-demo
   ```

### Kubernetes Deployment

1. Apply the Kubernetes manifests:

   ```bash
   kubectl apply -f k8s/
   ```

2. Verify the deployment:
   ```bash
   kubectl get pods
   kubectl get services
   ```

## Changelog

### [Unreleased]

- Initial project setup
- Implemented hexagonal architecture
- Added MongoDB integration
- Added Elasticsearch integration
- Set up basic HTTP server with Gin

### [0.1.0] - YYYY-MM-DD

- Initial release
- Basic CRUD operations
- Search functionality
- API documentation
