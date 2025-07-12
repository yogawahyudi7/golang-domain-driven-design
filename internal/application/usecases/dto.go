package usecases

import (
	"time"

	"github.com/google/uuid"
)

// CreateUserRequest represents the request to create a user
type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required,min=3,max=50"`
	FirstName string `json:"first_name" validate:"required,min=1,max=50"`
	LastName  string `json:"last_name" validate:"required,min=1,max=50"`
	Password  string `json:"password" validate:"required,min=8"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	FirstName string `json:"first_name" validate:"required,min=1,max=50"`
	LastName  string `json:"last_name" validate:"required,min=1,max=50"`
}

// UserResponse represents the user response
type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	FullName  string    `json:"full_name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ListUsersResponse represents the response for listing users
type ListUsersResponse struct {
	Users  []*UserResponse `json:"users"`
	Total  int64           `json:"total"`
	Offset int             `json:"offset"`
	Limit  int             `json:"limit"`
}

// AuthRequest represents the authentication request
type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	User        *UserResponse `json:"user"`
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
	ExpiresIn   int64         `json:"expires_in"`
	ExpiresAt   int64         `json:"expires_at"`
}
