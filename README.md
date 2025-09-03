# Go E-Commerce Backend

A robust, scalable e-commerce backend built with Go (Golang). This project demonstrates a clean architecture with separation of concerns, following industry best practices for building maintainable web services.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Clone the Repository](#clone-the-repository)
  - [Environment Setup](#environment-setup)
  - [Database Setup](#database-setup)
  - [Run the Application](#run-the-application)
- [API Endpoints](#api-endpoints)
- [Project Architecture](#project-architecture)
- [Makefile Commands](#makefile-commands)
- [Contributing](#contributing)
- [License](#license)

## Features

- User registration and authentication
- Password hashing and security
- Database migration support
- RESTful API design
- Clean architecture with separation of concerns
- Environment-based configuration
- Input validation
- Structured logging

## Tech Stack

- **Language**: Go 1.25.0
- **Web Framework**: Gorilla Mux
- **Database**: PostgreSQL (via pgx)
- **Migration Tool**: golang-migrate
- **Authentication**: JWT
- **Validation**: go-playground/validator
- **Environment Management**: godotenv

## Project Structure

```
src/
├── api/            # HTTP routing and handler registration
├── config/         # Configuration loading (from .env)
├── db/             # Database connection and migration setup
│   └── migrate/    # Migration files and runner
├── feature/        # Feature modules
│   └── user/       # User-related functionality
├── service/        # Business services
│   └── auth/       # Authentication utilities
└── utils/          # General utility functions
```

## Prerequisites

- Go 1.25.0 or higher
- PostgreSQL 12+ (or compatible database)
- golang-migrate CLI tool

### Install golang-migrate

On macOS:
```bash
brew install golang-migrate
```

On Linux:
```bash
curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
apt-get update
apt-get install -y migrate
```

For other platforms, refer to the [official installation guide](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md).

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/MKSinghDev/go-ecom.git
cd go-ecom
```

### Environment Setup

Create a `.env` file in the project root with the following variables (optional - fallback values will be used if not provided):

```env
# Server Configuration
PORT=8000
PUBLIC_HOST=http://localhost

# Database Configuration
DB_USERNAME=postgres
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=go-ecom
DB_SSL_MODE=disable

# JWT Configuration
JWT_SECRET=your_jwt_secret_key
JWT_EXP=604800
```

### Database Setup

1. Create the database in PostgreSQL:
   ```sql
   CREATE DATABASE go-ecom;
   ```

2. Run database migrations:
   ```bash
   make migrate-up
   ```

### Run the Application

```bash
# Build and run the application
make run

# Or just build
make build

# Or run directly
make run
```

The server will start on `http://localhost:8000` (or your configured PORT).

## API Endpoints

### User Management

| Method | Endpoint               | Description           |
|--------|------------------------|-----------------------|
| POST   | `/api/v1/register`     | Register a new user   |
| POST   | `/api/v1/login`        | User login            |

### Example Requests

**Register a new user:**
```bash
curl -X POST http://localhost:8000/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "password": "securepassword"
  }'
```

**User login:**
```bash
curl -X POST http://localhost:8000/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "password": "securepassword"
  }'
```

## Project Architecture

This project follows a layered architecture with clear separation of concerns:

1. **API Layer**: Handles HTTP routing and request/response handling (`api/`)
2. **Service Layer**: Business logic implementation (`service/`)
3. **Data Access Layer**: Database operations and repository patterns (`feature/*/repo.go`)
4. **Configuration Layer**: Environment-based configuration loading (`config/`)
5. **Migration Layer**: Database schema management (`db/migrate/`)

### Key Design Patterns

- **Dependency Injection**: Services and handlers are loosely coupled
- **Repository Pattern**: Data access is abstracted through interfaces
- **Singleton Pattern**: For configuration and database connection
- **Middleware Pattern**: For request processing pipeline

## Makefile Commands

| Command         | Description                          |
|----------------|--------------------------------------|
| `make build`    | Builds the application binary        |
| `make run`      | Builds and runs the application      |
| `make test`     | Runs all tests                       |
| `make migration name=migration_name` | Creates new migration files |
| `make migrate-up` | Applies all pending migrations     |
| `make migrate-down` | Rolls back the last migration    |

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.