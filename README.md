# Ohara API

A modern REST API built with Go and Fiber framework for managing books and users. This project features a clean architecture with PostgreSQL integration, comprehensive middleware, and Docker support.

## Features

- 🚀 **High Performance**: Built with Go Fiber framework
- 🗄️ **Database Integration**: PostgreSQL with GORM
- 🔒 **Security**: Helmet, CORS, and rate limiting
- 📚 **API Documentation**: Swagger/OpenAPI integration
- 🐳 **Containerized**: Docker and Docker Compose support
- 🔄 **Hot Reload**: Air for development
- 🏗️ **Clean Architecture**: Hexagonal architecture pattern
- 📝 **Request Logging**: Custom middleware with request IDs

## Tech Stack

- **Language**: Go 1.24
- **Framework**: Fiber v2
- **Database**: PostgreSQL
- **ORM**: GORM
- **Documentation**: Swagger
- **Containerization**: Docker
- **Development**: Air (hot reload)

## Project Structure

```
ohara-api/
├── cmd/
│   └── server.go           # Application entry point
├── internal/
│   ├── adapters/
│   │   ├── primary/http/   # HTTP handlers and middleware
│   │   └── secondary/      # Database and external services
│   ├── config/
│   │   └── env/           # Environment configuration
│   └── core/
│       └── domain/        # Domain models (Books, Users)
├── docs/                  # Swagger documentation
├── pkg/                   # Shared utilities
├── docker-compose.yml     # Docker services
├── Dockerfile            # Production Docker image
├── dev.Dockerfile        # Development Docker image
└── Makefile              # Build and development commands
```

## Getting Started

### Prerequisites

- Go 1.24 or higher
- PostgreSQL
- Docker (optional)
- Make (optional, for convenience commands)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/nNottp33/ohara-api.git
   cd ohara-api
   ```

2. **Copy environment file**
   ```bash
   cp env.example .env
   ```

3. **Configure environment variables**
   Edit `.env` file with your settings:
   ```env
   # Application
   APP_ENV=development
   APP_PORT=3033
   PREFIX_REQUEST_ID=prefix-request
   
   # Database
   DB_HOST=127.0.0.1
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=postgres
   DB_MAX_CONNECTION=10
   DB_MAX_IDLE_CONNECTION=10
   ```

4. **Install dependencies**
   ```bash
   go mod download
   ```

### Running the Application

#### Option 1: Local Development (with hot reload)
```bash
# Install Air if not already installed
go install github.com/air-verse/air@latest

# Run with hot reload
make run
```

#### Option 2: Docker Compose
```bash
# Start all services
make up

# View logs
make logs

# Stop services
make down
```

#### Option 3: Manual Build and Run
```bash
# Build the application
make build

# Run the binary
./tmp/main
```

## API Documentation

When running in development mode, Swagger documentation is available at:
```
http://localhost:3033/swagger
```

To regenerate Swagger docs:
```bash
make swag
```

## Available Make Commands

- `make run` - Run with hot reload (development)
- `make build` - Build the application
- `make up` - Start Docker services
- `make down` - Stop Docker services
- `make restart SERVICE=app` - Restart specific service
- `make logs` - View application logs
- `make swag` - Generate Swagger documentation
- `make clean` - Clean build artifacts
- `make prune` - Clean Docker resources

## Database Models

### Books
- ID, Name, Description
- Image URL, Author, Publisher
- ISBN (unique), Price
- Timestamps (created_at, updated_at)

### Users
- ID, Name
- Timestamps (created_at, updated_at)

## Middleware

The API includes several middleware layers:

- **Request ID**: Custom request ID generation
- **Logger**: Request/response logging
- **CORS**: Cross-origin resource sharing
- **Helmet**: Security headers
- **Compress**: Response compression
- **Rate Limiter**: API rate limiting (20 requests per 30 seconds)

## Environment Configuration

The application supports different environments:
- `development` - Enables Swagger docs, verbose logging
- `production` - Optimized for production use

## Docker Support

### Development
```bash
docker-compose up --build
```

### Production
```bash
docker-compose -f docker-compose.deploy.yml up --build
```

## API Endpoints

- `GET /` - Health check endpoint
- API documentation available at `/swagger` (development only)

*Note: Additional endpoints will be documented as the API grows*

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

Project Link: [https://github.com/nNottp33/ohara-api](https://github.com/nNottp33/ohara-api)