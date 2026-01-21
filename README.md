# CashInvoice  Todo Assignment

A simple yet powerful **Todo Task Management API** built with Go, featuring user authentication, task tracking with status management, and automatic task completion features.

---

## 📋 Table of Contents

- [Features](#features)
- [Technology Stack](#technology-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [API Endpoints](#api-endpoints)
- [Usage Examples](#usage-examples)
- [Database Schema](#database-schema)
- [Error Handling](#error-handling)

---

## ✨ Features

- **User Authentication**: Register and login with secure password hashing using bcrypt
- **JWT Token-based Access**: Secure API endpoints with JWT authentication
- **Todo Management**: Create, read, update, and delete todos
- **Task Status Tracking**: Track todos as pending, in-progress, or completed
- **Pagination**: Get todos with customizable page size
- **Filtering**: Filter todos by status
- **Role-Based Access**: Support for different user roles (user/admin)
- **Admin Access**: Admins can view all todos across users
- **Auto-Complete Feature**: Automatically mark old pending todos as completed
- **User Profile**: Access your user information

---

## 🛠 Technology Stack

- **Language**: Go 1.25.1
- **Web Framework**: Fiber v2 (Fast HTTP framework)
- **Database**: PostgreSQL
- **ORM**: GORM (Go Object-Relational Mapping)
- **Authentication**: JWT (JSON Web Tokens)
- **Password Security**: bcrypt
- **Database Migrations**: GORM AutoMigrate

---

## 📁 Project Structure

```
cashinvoice-assignment/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── database/
│   │   └── postgresql.go         # Database connection
│   ├── handler/
│   │   ├── auth_handler.go       # Authentication endpoints
│   │   └── todo_handler.go       # Todo endpoints
│   ├── middleware/
│   │   └── auth_middleware.go    # JWT authentication middleware
│   ├── model/
│   │   ├── base.go              # Base model with common fields
│   │   ├── user.go              # User model
│   │   └── todo.go              # Todo model
│   ├── repository/
│   │   ├── user_repository.go    # User data access
│   │   └── todo_repository.go    # Todo data access
│   ├── service/
│   │   ├── auth_service.go       # Authentication logic
│   │   └── todo_service.go       # Todo business logic
│   ├── router/
│   │   └── router.go             # API route definitions
│   ├── utils/
│   │   ├── password.go           # Password hashing utilities
│   │   └── jwt.go                # JWT token utilities
│   └── errors/
│       └── errors.go             # Custom error definitions
├── go.mod                        # Go module definition
└── README.md                     # This file
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.25.1 or higher
- PostgreSQL 12 or higher
- Git

### Installation Steps

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd cashinvoice-assignment
   ```

2. **Install dependencies**
   ```bash
   go get ./...
   ```

3. **Set up PostgreSQL database**
   ```bash
   createdb cashinvoice_db
   ```

4. **Configure environment variables** (see Configuration section)

5. **Run the application**
   ```bash
   go run cmd/server/main.go
   ```

The server will start on `http://localhost:8080`

### Verify Installation

Test if the server is running:
```bash
curl http://localhost:8080/ping
```

Expected response:
```json
{"status":"ok"}
```

---

## ⚙️ Configuration

Create a `.env` file in the project root with the following variables:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres123
DB_NAME=cashinvoice_db
DB_SSLMODE=disable

# JWT Secret (keep this secret!)
JWT_SECRET=your-secret-key-here

# Auto-Complete Feature
AUTO_COMPLETE_DELAY=10
```

### Configuration Details(Dummy Values)

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | localhost | PostgreSQL host address |
| `DB_PORT` | 5432 | PostgreSQL port |
| `DB_USER` | postgres | Database username |
| `DB_PASSWORD` | postgres123 | Database password |
| `DB_NAME` | postgres | Database name |
| `DB_SSLMODE` | disable | SSL mode for database connection |
| `JWT_SECRET` | supersecretkey | Secret key for JWT token signing |
| `AUTO_COMPLETE_DELAY` | 10 | Minutes to wait before auto-completing pending todos |

---

## 📡 API Endpoints

### Base URL
```
http://localhost:8080
```

### Authentication Endpoints

#### 1. Register a New User
Create a new user account.

**Endpoint**: `POST /auth/register`

**Request Body**:
```json
{
  "name": "Kartik",
  "email": "kartik@example.com",
  "password": "password123",
  "role": "user"
}
```

**Response** (201 Created):
```json
{
  "message": "user registered successfully"
}
```

**Error Response** (400):
```json
{
  "error": "user already exists"
}
```

---

#### 2. Login
Get an authentication token to access protected endpoints.

**Endpoint**: `POST /auth/login`

**Request Body**:
```json
{
  "email": "kartik@example.com",
  "password": "password123"
}
```

**Response** (200 OK):
```json
{
  "message": "login successful",
  "user_id": 1,
  "access_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Error Response** (401):
```json
{
  "error": "invalid credentials"
}
```

**How to use the token**: Include it in the `Authorization` header for protected endpoints:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
```

---

### Protected Endpoints (Require Authentication)

All endpoints below require the `Authorization: Bearer <token>` header.

#### 3. Get User Profile
View your own user information.

**Endpoint**: `GET /api/profile`

**Headers**:
```
Authorization: Bearer <your-token>
```

**Response** (200 OK):
```json
{
  "user_id": 1,
  "email": "kartik@example.com"
}
```

---

### Todo Endpoints

#### 4. Create a New Todo
Add a new task to your todo list.

**Endpoint**: `POST /api/todos`

**Headers**:
```
Authorization: Bearer <your-token>
Content-Type: application/json
```

**Request Body**:
```json
{
  "title": "Buy groceries",
  "description": "Milk, eggs, bread, and vegetables",
  "status": "pending"
}
```

**Status Values**: `pending`, `in_progress`, `completed`

**Response** (201 Created):
```json
{
  "ID": "550e8400-e29b-41d4-a716-446655440000",
  "Title": "Buy groceries",
  "Description": "Milk, eggs, bread, and vegetables",
  "Status": "pending",
  "CreatedAt": "2026-01-22T10:30:00Z",
  "UpdatedAt": "2026-01-22T10:30:00Z",
  "UserID": 1
}
```

---

#### 5. Get All Todos (with Pagination & Filtering)
Retrieve your todos with optional pagination and status filtering.

**Endpoint**: `GET /api/todos`

**Headers**:
```
Authorization: Bearer <your-token>
```

**Query Parameters**:
| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `page` | integer | 1 | Page number for pagination |
| `limit` | integer | 10 | Number of todos per page (max: 100) |
| `status` | string | (none) | Filter by status: `pending`, `in_progress`, or `completed` |

**Example Requests**:
```bash
# Get first 10 todos
GET /api/todos

# Get page 2 with 20 items per page
GET /api/todos?page=2&limit=20

# Get only pending todos
GET /api/todos?status=pending

# Get completed todos on page 3
GET /api/todos?page=3&status=completed
```

**Response** (200 OK):
```json
{
  "data": [
    {
      "ID": "550e8400-e29b-41d4-a716-446655440000",
      "Title": "Buy groceries",
      "Description": "Milk, eggs, bread, and vegetables",
      "Status": "pending",
      "CreatedAt": "2026-01-22T10:30:00Z",
      "UpdatedAt": "2026-01-22T10:30:00Z",
      "UserID": 1
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 45,
    "total_pages": 5
  }
}
```

**Admin Access**: Admins see all todos from all users (when `role: "admin"`).

---

#### 6. Update a Todo
Modify an existing todo's title, description, or status.

**Endpoint**: `PUT /api/todos/:id`

**Headers**:
```
Authorization: Bearer <your-token>
Content-Type: application/json
```

**URL Parameters**:
- `id` (required): The UUID of the todo to update

**Request Body**:
```json
{
  "title": "Buy groceries and cook dinner",
  "description": "Updated description",
  "status": "in_progress"
}
```

**Response** (200 OK):
```json
{
  "ID": "550e8400-e29b-41d4-a716-446655440000",
  "Title": "Buy groceries and cook dinner",
  "Description": "Updated description",
  "Status": "in_progress",
  "CreatedAt": "2026-01-22T10:30:00Z",
  "UpdatedAt": "2026-01-22T11:45:00Z",
  "UserID": 1
}
```

**Error Response** (403):
```json
{
  "error": "forbidden"
}
```
(Only the todo owner can update their todos, except admins)

---

#### 7. Delete a Todo
Remove a todo from your list.

**Endpoint**: `DELETE /api/todos/:id`

**Headers**:
```
Authorization: Bearer <your-token>
```

**URL Parameters**:
- `id` (required): The UUID of the todo to delete

**Response** (204 No Content)
Empty response body indicates successful deletion.

**Error Response** (403):
```json
{
  "error": "forbidden"
}
```
(Only the todo owner can delete their todos, except admins)

---

## 💡 Usage Examples

### Complete Workflow

**1. Register a new user**
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice",
    "email": "alice@example.com",
    "password": "securepass123",
    "role": "user"
  }'
```

**2. Login to get token**
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alice@example.com",
    "password": "securepass123"
  }'
```

**3. Create a todo**
```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Finish project",
    "description": "Complete the assignment",
    "status": "pending"
  }'
```

**4. Get todos with filters**
```bash
# Get pending todos
curl -X GET "http://localhost:8080/api/todos?status=pending" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"

# Get completed todos with pagination
curl -X GET "http://localhost:8080/api/todos?status=completed&page=1&limit=5" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**5. Update a todo**
```bash
curl -X PUT http://localhost:8080/api/todos/550e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Finish project",
    "description": "Completed the assignment",
    "status": "completed"
  }'
```

**6. Delete a todo**
```bash
curl -X DELETE http://localhost:8080/api/todos/550e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

---

## 💾 Database Schema

### Users Table
Stores user account information.

| Column | Type | Constraints |
|--------|------|-------------|
| id | BIGINT | PRIMARY KEY, Auto-increment |
| name | VARCHAR(100) | NOT NULL |
| email | VARCHAR(255) | UNIQUE, NOT NULL |
| password | VARCHAR(255) | NOT NULL (bcrypt hashed) |
| role | VARCHAR(20) | DEFAULT 'user' |
| created_at | TIMESTAMP | Auto-generated |
| updated_at | TIMESTAMP | Auto-generated |
| deleted_at | TIMESTAMP | For soft delete |

### Todos Table
Stores user task information.

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY |
| title | VARCHAR(255) | NOT NULL |
| description | TEXT | (optional) |
| status | VARCHAR(20) | DEFAULT 'pending' |
| user_id | BIGINT | FOREIGN KEY (users.id) |
| created_at | TIMESTAMP | Auto-generated |
| updated_at | TIMESTAMP | Auto-generated |
| deleted_at | TIMESTAMP | For soft delete |

---

## ⚠️ Error Handling

The API returns standard HTTP status codes and error messages in JSON format.

### Common HTTP Status Codes

| Status Code | Meaning |
|------------|---------|
| 200 | OK - Request successful |
| 201 | Created - Resource successfully created |
| 204 | No Content - Request successful but no content to return |
| 400 | Bad Request - Invalid input or parameters |
| 401 | Unauthorized - Missing or invalid authentication token |
| 403 | Forbidden - You don't have permission to access this resource |
| 404 | Not Found - Resource not found |
| 500 | Internal Server Error - Server error |

### Error Response Format

```json
{
  "error": "error message describing what went wrong"
}
```

### Common Errors

| Error | Cause | Solution |
|-------|-------|----------|
| `missing token` | No authorization header provided | Add `Authorization: Bearer <token>` header |
| `invalid token` | Token is expired or malformed | Login again to get a fresh token |
| `invalid credentials` | Wrong email or password | Check your email and password |
| `user already exists` | Email is already registered | Use a different email or login |
| `forbidden` | Trying to modify another user's todo | Only owners or admins can modify todos |
| `invalid status` | Invalid status value provided | Use: pending, in_progress, or completed |

---

## 🔒 Security Notes

- **Passwords**: All passwords are hashed using bcrypt before storage
- **JWT Tokens**: Tokens expire after 24 hours
- **API Authentication**: All `/api/*` endpoints require valid JWT tokens
- **Role-Based Access**: Only admins can see all users' todos
- **Ownership**: Users can only modify/delete their own todos
- **Environment Variables**: Keep `JWT_SECRET` and `DB_PASSWORD` secure

