# HRM Backend API

Human Resource Management Backend API พร้อม Swagger Documentation

## Features

- ✅ User Registration
- ✅ Leave Request Management
- ✅ RESTful API
- ✅ Swagger Documentation
- ✅ CORS Support
- ✅ MySQL Database Integration

## Getting Started

### Prerequisites

- Go 1.24+
- MySQL Database
- Swag CLI tool

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```

3. Install Swag CLI (if not already installed):
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

### Configuration

Create your database configuration in the config files.

### Running the Application

1. Generate Swagger docs:
   ```bash
   swag init -g cmd/main.go -o docs/
   ```

2. Run the application:
   ```bash
   go run cmd/main.go
   ```

3. The server will start on `http://localhost:3000`

### API Documentation

Once the server is running, you can access the Swagger documentation at:

**📚 Swagger UI: [http://localhost:3000/swagger/](http://localhost:3000/swagger/)**

### API Endpoints

#### User Management

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/users/register` | Register a new user |

#### Leave Management

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/leaves/request` | Submit a new leave request |
| GET | `/api/v1/leaves/all` | Get all leave requests |
| PUT | `/api/v1/leaves/{id}/approve` | Approve a leave request |
| PUT | `/api/v1/leaves/{id}/reject` | Reject a leave request |

### Example API Requests

#### Register User
```bash
curl -X POST http://localhost:3000/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "password123",
    "password_confirmation": "password123"
  }'
```

#### Request Leave
```bash
curl -X POST http://localhost:3000/api/v1/leaves/request \
  -H "Content-Type: application/json" \
  -d '{
    "employee_id": 123,
    "reason": "Family vacation",
    "from_date": "2024-01-15T00:00:00Z",
    "to_date": "2024-01-20T00:00:00Z"
  }'
```

#### Get All Leaves
```bash
curl -X GET http://localhost:3000/api/v1/leaves/all
```

#### Approve Leave
```bash
curl -X PUT http://localhost:3000/api/v1/leaves/1/approve
```

### Development

#### Make Commands

```bash
# Generate Swagger docs
make swag-init

# Format Swagger comments
make swag-fmt

# Run the application
make run

# Build the application
make build

# Generate docs and run
make dev
```

### Project Structure

```
backend/
├── cmd/
│   └── main.go              # Application entry point
├── docs/                    # Swagger documentation
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── adapter/
│   │   ├── config/          # Configuration
│   │   ├── handler/         # HTTP handlers
│   │   └── storage/         # Database layer
│   └── core/
│       ├── domain/          # Domain models
│       ├── port/            # Interfaces
│       └── service/         # Business logic
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

### Contributing

1. Fork the repository
2. Create your feature branch
3. Add Swagger annotations for new endpoints
4. Regenerate Swagger docs: `make swag-init`
5. Commit your changes
6. Push to the branch
7. Create a Pull Request

## License

This project is licensed under the MIT License.
