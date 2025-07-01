# Sinbitus

A modern habit tracking application built with Go, designed to help users build and maintain positive habits through a clean REST API.

## 🚀 Features

- **User Authentication**: Secure JWT-based authentication system
- **Habit Management**: Create, track, and manage personal habits
- **Event Tracking**: Record habit completion events with timestamps
- **RESTful API**: Clean HTTP API with comprehensive endpoints
- **Database Persistence**: PostgreSQL backend with proper indexing
- **Monitoring**: Built-in health checks and Prometheus metrics
- **Logging**: Structured logging with configurable levels
- **Dependency Injection**: Clean architecture using Uber FX

## 🏗️ Architecture

The project follows clean architecture principles with clear separation of concerns:

```
├── cmd/sinbitus/          # Application entry point
├── configs/               # Configuration files
├── internal/              # Private application code
│   ├── appfx/            # Application framework setup
│   ├── config/           # Configuration management
│   ├── constants/        # Application constants
│   ├── controller/       # HTTP controllers and middleware
│   ├── dto/              # Data Transfer Objects
│   ├── entity/           # Domain entities
│   ├── infra/            # Infrastructure layer (database)
│   ├── pkg/              # Internal packages (logger, time)
│   └── usecase/          # Business logic layer
├── migrations/           # Database migrations
└── deployment/           # Docker and deployment configs
```

## 🛠️ Tech Stack

- **Language**: Go 1.24.4
- **Framework**: Fiber v3 (HTTP server)
- **Database**: PostgreSQL 15
- **ORM**: pgx with Squirrel query builder
- **Authentication**: JWT tokens
- **Dependency Injection**: Uber FX
- **Logging**: Structured logging with slog
- **Configuration**: Viper with YAML support
- **Monitoring**: Prometheus metrics
- **Containerization**: Docker & Docker Compose

## ⚙️ Configuration

The application configuration is located in `configs/sinbitus/config.yml`:

```yaml
log:
  level: debug
  timestamp: yes
  formatter: text
  color: yes

http:
  address: "0.0.0.0"
  port: 9060

postgres:
  url: postgres://admin:CHANGE_ME@localhost:5432/sinbitus?sslmode=disable
  max_pool_size: 1
  conn_attempts: 10
  conn_timeout: 1s

jwt:
  secret_key: CHANGE_ME
  ttl: 24h
```

**Important**: Update the PostgreSQL URL and JWT secret key before running in production.

## 📡 API Endpoints

### Authentication
- `POST /api/v1/signup` - User registration
- `POST /api/v1/login` - User authentication

### System
- `GET /health` - Health check
- `GET /metrics` - Prometheus metrics
- `GET /api/v1/profile` - Get user profile
- `POST /api/v1/habits` - Create habit
- `GET /api/v1/habits` - List user habits
- `GET /api/v1/habits/:id` - Get specific habit
- `PUT /api/v1/habits/:id` - Update habit
- `DELETE /api/v1/habits/:id` - Delete habit
- `POST /api/v1/habits/:id/track` - Track habit event
- `GET /api/v1/habits/:id/track` - Get habit tracking history

## 🗄️ Database Schema

The application uses three main tables:

- **users**: User accounts with authentication
- **habits**: User-defined habits with descriptions
- **habit_events**: Tracking events for habit completion

## 🛠️ Development

### Available Make Commands

```bash
make migrate-create <migration_name>  # Create new migration
make migrate-up                       # Run database migrations
```

## 🔧 Environment Variables

The application supports environment variables for configuration:

- `POSTGRES_URL`: PostgreSQL connection string
- `JWT_SECRET_KEY`: Secret key for JWT token generation
- `HTTP_PORT`: HTTP server port
- `LOG_LEVEL`: Logging level
