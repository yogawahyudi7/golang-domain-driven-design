package repositories

import (
	"context"

	"golang-domain-driven-design/internal/domain/entities"

	"github.com/google/uuid"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, user *entities.User) error

	// GetByID retrieves a user by ID
	GetByID(ctx context.Context, id uuid.UUID) (*entities.User, error)

	// GetByEmail retrieves a user by email
	GetByEmail(ctx context.Context, email string) (*entities.User, error)

	// GetByUsername retrieves a user by username
	GetByUsername(ctx context.Context, username string) (*entities.User, error)

	// Update updates an existing user
	Update(ctx context.Context, user *entities.User) error

	// Delete deletes a user by ID
	Delete(ctx context.Context, id uuid.UUID) error

	// List retrieves a list of users with pagination
	List(ctx context.Context, offset, limit int) ([]*entities.User, error)

	// Count returns the total number of users
	Count(ctx context.Context) (int64, error)

	// ExistsByEmail checks if a user exists by email
	ExistsByEmail(ctx context.Context, email string) (bool, error)

	// ExistsByUsername checks if a user exists by username
	ExistsByUsername(ctx context.Context, username string) (bool, error)
}
