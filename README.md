# Mini App - Task Management REST API

[![Go Version](https://img.shields.io/badge/Go-1.25.1-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.10.1-00ADD8?style=flat)](https://github.com/gin-gonic/gin)
[![MySQL](https://img.shields.io/badge/MySQL-8.0-4479A1?style=flat&logo=mysql&logoColor=white)](https://www.mysql.com/)
[![GORM](https://img.shields.io/badge/GORM-v1.31.1-00ADD8?style=flat)](https://gorm.io/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

A lightweight RESTful API built with Go (Golang) for managing users and tasks. This project demonstrates clean architecture principles, following best practices for building scalable backend applications.

## 📋 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Prerequisites](#-prerequisites)
- [Installation](#-installation)
- [Configuration](#-configuration)
- [Running the Application](#-running-the-application)
- [Database Migrations](#-database-migrations)
- [API Documentation](#-api-documentation)
- [API Endpoints](#-api-endpoints)
- [Usage Examples](#-usage-examples)
- [Development](#-development)
- [Testing](#-testing)
- [Docker Support](#-docker-support)
- [Contributing](#-contributing)
- [License](#-license)

## ✨ Features

- **RESTful API** - Clean and intuitive REST endpoints
- **User Management** - Create, read, update, and delete users
- **Task Management** - Manage tasks with status tracking (pending/done)
- **Database Migrations** - Version-controlled database schema changes
- **Swagger Documentation** - Interactive API documentation
- **Docker Support** - Easy deployment with Docker Compose
- **Clean Architecture** - Separation of concerns with layered architecture
- **Input Validation** - Request validation using go-playground/validator
- **GORM ORM** - Efficient database operations with GORM
- **Error Handling** - Consistent error responses

## 🛠 Tech Stack

- **Language**: Go 1.25.1
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin) - High-performance HTTP web framework
- **Database**: MySQL 8.0
- **ORM**: [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- **Validation**: [go-playground/validator](https://github.com/go-playground/validator)
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Documentation**: [Swagger/OpenAPI](https://swagger.io/) with [swaggo](https://github.com/swaggo/swag)
- **Environment**: [godotenv](https://github.com/joho/godotenv)
- **Containerization**: Docker & Docker Compose

## 📁 Project Structure

```
miniapp/
├── cmd/
│   ├── main.go                 # Application entry point
│   └── migrate/
│       └── main.go             # Database migration runner
├── internal/
│   ├── config/
│   │   └── database.go         # Database configuration
│   ├── data/
│   │   └── models/             # Data models
│   │       ├── user.go
│   │       └── task.go
│   ├── dto/                    # Data Transfer Objects
│   │   ├── user_dto.go
│   │   └── task_dto.go
│   ├── handlers/               # HTTP request handlers
│   │   ├── user_handler.go
│   │   └── task_handler.go
│   ├── repositories/           # Data access layer
│   │   ├── user_repository.go
│   │   └── task_repository.go
│   ├── services/               # Business logic layer
│   │   ├── user_service.go
│   │   └── task_service.go
│   └── routes/
│       └── routes.go           # API route definitions
├── migrations/                 # SQL migration files
│   ├── 000001_create_users_table.up.sql
│   ├── 000001_create_users_table.down.sql
│   ├── 000002_create_tasks_table.up.sql
│   └── 000002_create_tasks_table.down.sql
├── pkg/
│   └── utils/                  # Utility functions
│       ├── response.go
│       └── validator.go
├── docs/                       # Swagger documentation
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── deployments/
│   └── docker-compose.yml      # Docker Compose configuration
├── go.mod                      # Go module dependencies
├── go.sum                      # Dependency checksums
└── README.md
```

### Architecture Layers

- **Handlers**: Handle HTTP requests and responses
- **Services**: Contain business logic
- **Repositories**: Handle database operations
- **Models**: Define data structures
- **DTOs**: Define request/response data structures

## 📦 Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.25.1 or higher)
- [MySQL](https://dev.mysql.com/downloads/mysql/) (version 8.0 or higher)
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerized deployment)
- [Docker Compose](https://docs.docker.com/compose/install/) (optional)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) (for running migrations)

## 🚀 Installation

### 1. Clone the Repository

```bash
git clone https://github.com/hvmidrezv/miniapp.git
cd miniapp
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Install Swagger CLI (for generating docs)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## ⚙️ Configuration

### 1. Create Environment File

Create a `.env` file in the root directory:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=miniapp
DB_PASSWORD=miniappt0lk2o20
DB_NAME=miniapp_db

# Server Configuration
SERVER_PORT=8080

# Application
APP_ENV=development
```

### 2. Database Setup

#### Option A: Using Docker (Recommended)

```bash
# Start MySQL container
docker-compose -f deployments/docker-compose.yml up -d
```

The Docker Compose file will:
- Create a MySQL 8.0 instance
- Expose it on port `3308` (mapped from container port 3306)
- Create the `miniapp_db` database
- Set up user credentials

#### Option B: Manual MySQL Setup

Create the database manually:

```sql
CREATE DATABASE miniapp_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'miniapp'@'localhost' IDENTIFIED BY 'miniappt0lk2o20';
GRANT ALL PRIVILEGES ON miniapp_db.* TO 'miniapp'@'localhost';
FLUSH PRIVILEGES;
```

## 🏃 Running the Application

### 1. Run Database Migrations

```bash
# Set the database connection string
migrate -path migrations -database "mysql://miniapp:miniappt0lk2o20@tcp(localhost:3308)/miniapp_db" up
```

Or use the migration tool:

```bash
go run cmd/migrate/main.go
```

### 2. Generate Swagger Documentation

```bash
swag init -g cmd/main.go -o docs
```

### 3. Start the Server

```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080`

### 4. Verify Installation

Open your browser and visit:
- API Health: `http://localhost:8080/api/users`
- Swagger UI: `http://localhost:8080/swagger/index.html`

## 🗄️ Database Migrations

### Running Migrations

```bash
# Apply all up migrations
migrate -path migrations -database "mysql://miniapp:miniappt0lk2o20@tcp(localhost:3308)/miniapp_db" up

# Rollback last migration
migrate -path migrations -database "mysql://miniapp:miniappt0lk2o20@tcp(localhost:3308)/miniapp_db" down 1

# Check migration version
migrate -path migrations -database "mysql://miniapp:miniappt0lk2o20@tcp(localhost:3308)/miniapp_db" version
```

### Creating New Migrations

```bash
migrate create -ext sql -dir migrations -seq create_new_table
```

## 📚 API Documentation

Interactive API documentation is available via Swagger UI:

**URL**: `http://localhost:8080/swagger/index.html`

The Swagger documentation provides:
- Complete API endpoint reference
- Request/response schemas
- Interactive API testing
- Example requests and responses

## 🔌 API Endpoints

### User Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/users` | Get all users |
| GET | `/api/users/:id` | Get user by ID |
| POST | `/api/users` | Create new user |
| PUT | `/api/users/:id` | Update user |
| DELETE | `/api/users/:id` | Delete user |

### Task Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/tasks` | Get all tasks |
| GET | `/api/tasks/:id` | Get task by ID |
| POST | `/api/tasks` | Create new task |
| PUT | `/api/tasks/:id` | Update task |
| DELETE | `/api/tasks/:id` | Delete task |

## 💡 Usage Examples

### Create a User

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john.doe@example.com"
  }'
```

**Response:**
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "created_at": "2025-11-20T10:00:00Z",
    "updated_at": "2025-11-20T10:00:00Z"
  }
}
```

### Get All Users

```bash
curl -X GET http://localhost:8080/api/users
```

**Response:**
```json
{
  "success": true,
  "message": "Users retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "created_at": "2025-11-20T10:00:00Z",
      "updated_at": "2025-11-20T10:00:00Z"
    }
  ]
}
```

### Create a Task

```bash
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "title": "Complete project documentation",
    "status": "pending"
  }'
```

**Response:**
```json
{
  "success": true,
  "message": "Task created successfully",
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Complete project documentation",
    "status": "pending",
    "created_at": "2025-11-20T10:05:00Z",
    "updated_at": "2025-11-20T10:05:00Z"
  }
}
```

### Update Task Status

```bash
curl -X PUT http://localhost:8080/api/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Complete project documentation",
    "status": "done",
    "user_id": 1
  }'
```

### Get User's Tasks

```bash
curl -X GET http://localhost:8080/api/users/1
```

### Delete a Task

```bash
curl -X DELETE http://localhost:8080/api/tasks/1
```

## 🔧 Development

### Code Structure Best Practices

This project follows clean architecture principles:

1. **Handler Layer**: Receives HTTP requests, validates input, calls service layer
2. **Service Layer**: Contains business logic, orchestrates data operations
3. **Repository Layer**: Handles database operations, abstracts data access
4. **Model Layer**: Defines data structures and relationships

### Adding New Features

1. **Define the model** in `internal/data/models/`
2. **Create migration files** in `migrations/`
3. **Create DTO** in `internal/dto/`
4. **Implement repository** in `internal/repositories/`
5. **Implement service** in `internal/services/`
6. **Create handler** in `internal/handlers/`
7. **Register routes** in `internal/routes/routes.go`
8. **Update Swagger docs** with appropriate annotations

### Regenerate Swagger Docs

After modifying API endpoints or adding Swagger annotations:

```bash
swag init -g cmd/main.go -o docs
```

### Code Formatting

```bash
# Format code
go fmt ./...
```

## 🐳 Docker Support

### Using Docker Compose

The project includes a `docker-compose.yml` for the MySQL database.

#### Start Services

```bash
docker-compose -f deployments/docker-compose.yml up -d
```

#### Stop Services

```bash
docker-compose -f deployments/docker-compose.yml down
```

#### View Logs

```bash
docker-compose -f deployments/docker-compose.yml logs -f
```

### Database Connection

When using Docker Compose:
- **Host**: `localhost`
- **Port**: `3308` (external), `3306` (internal)
- **Database**: `miniapp_db`
- **Username**: `miniapp`
- **Password**: `miniappt0lk2o20`
- **Root Password**: `miniappRoo7t0lk2o20`

### Connecting to MySQL Container

```bash
docker exec -it miniapp-database mysql -u miniapp -p
# Enter password: miniappt0lk2o20
```

## 🏗️ Building for Production

### Build Binary

```bash
# Build for current platform
go build -o miniapp cmd/main.go
```

### Run Production Build

```bash
./miniapp
```


