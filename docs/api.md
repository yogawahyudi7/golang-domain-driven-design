# API Documentation

## Authentication

### Login
- **URL**: `POST /api/v1/auth/login`
- **Description**: Authenticate user and get access token
- **Request Body**:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```
- **Response**:
```json
{
  "message": "Authentication successful",
  "data": {
    "user": {
      "id": "uuid",
      "email": "user@example.com",
      "username": "username",
      "first_name": "John",
      "last_name": "Doe",
      "full_name": "John Doe",
      "is_active": true,
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    },
    "access_token": "jwt-token",
    "token_type": "Bearer",
    "expires_in": 86400
  }
}
```

## Users

### Create User
- **URL**: `POST /api/v1/users`
- **Description**: Create a new user
- **Request Body**:
```json
{
  "email": "user@example.com",
  "username": "username",
  "first_name": "John",
  "last_name": "Doe",
  "password": "SecurePassword123!"
}
```

### Get User
- **URL**: `GET /api/v1/users/:id`
- **Description**: Get user by ID
- **Parameters**: 
  - `id` (UUID): User ID

### Update User
- **URL**: `PUT /api/v1/users/:id`
- **Description**: Update user information
- **Request Body**:
```json
{
  "first_name": "John",
  "last_name": "Doe"
}
```

### Delete User
- **URL**: `DELETE /api/v1/users/:id`
- **Description**: Delete user by ID
- **Parameters**: 
  - `id` (UUID): User ID

### List Users
- **URL**: `GET /api/v1/users`
- **Description**: Get list of users with pagination
- **Query Parameters**:
  - `offset` (int): Offset for pagination (default: 0)
  - `limit` (int): Limit for pagination (default: 10, max: 100)

## Health Check

### Health Check
- **URL**: `GET /health`
- **Description**: Check service health
- **Response**:
```json
{
  "status": "healthy",
  "message": "Service is running",
  "service": "golang-domain-driven-design"
}
```

## Error Responses

### Validation Error (400)
```json
{
  "error": "Validation Error",
  "message": "Invalid email format",
  "field": "email"
}
```

### Not Found Error (404)
```json
{
  "error": "Not Found",
  "message": "User not found"
}
```

### Unauthorized Error (401)
```json
{
  "error": "Unauthorized",
  "message": "Invalid email or password"
}
```

### Internal Server Error (500)
```json
{
  "error": "Internal Server Error",
  "message": "An unexpected error occurred"
}
```
