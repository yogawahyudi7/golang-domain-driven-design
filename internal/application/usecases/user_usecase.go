package usecases

import (
	"context"
	"fmt"

	"golang-domain-driven-design/internal/domain/entities"
	"golang-domain-driven-design/internal/domain/repositories"
	"golang-domain-driven-design/internal/domain/valueobjects"
	apperrors "golang-domain-driven-design/pkg/errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserUseCase defines the interface for user use cases
type UserUseCase interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*UserResponse, error)
	GetUser(ctx context.Context, id uuid.UUID) (*UserResponse, error)
	UpdateUser(ctx context.Context, id uuid.UUID, req UpdateUserRequest) (*UserResponse, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	ListUsers(ctx context.Context, offset, limit int) (*ListUsersResponse, error)
	AuthenticateUser(ctx context.Context, email, password string) (*UserResponse, error)
}

// userUseCase implements UserUseCase interface
type userUseCase struct {
	userRepo repositories.UserRepository
}

// NewUserUseCase creates a new user use case
func NewUserUseCase(userRepo repositories.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user
func (uc *userUseCase) CreateUser(ctx context.Context, req CreateUserRequest) (*UserResponse, error) {
	// Validate email
	email, err := valueobjects.NewEmail(req.Email)
	if err != nil {
		return nil, apperrors.NewValidationError("email", err.Error())
	}

	// Validate password
	password, err := valueobjects.NewPassword(req.Password)
	if err != nil {
		return nil, apperrors.NewValidationError("password", err.Error())
	}

	// Check if email already exists
	exists, err := uc.userRepo.ExistsByEmail(ctx, email.Value())
	if err != nil {
		return nil, fmt.Errorf("failed to check email existence: %w", err)
	}
	if exists {
		return nil, apperrors.NewValidationError("email", "email already exists")
	}

	// Check if username already exists
	exists, err = uc.userRepo.ExistsByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check username existence: %w", err)
	}
	if exists {
		return nil, apperrors.NewValidationError("username", "username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.Value()), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user entity
	user := entities.NewUser(
		email.Value(),
		req.Username,
		req.FirstName,
		req.LastName,
		string(hashedPassword),
	)

	// Save user
	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return toUserResponse(user), nil
}

// GetUser retrieves a user by ID
func (uc *userUseCase) GetUser(ctx context.Context, id uuid.UUID) (*UserResponse, error) {
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return nil, apperrors.NewNotFoundError("user", "user not found")
	}

	return toUserResponse(user), nil
}

// UpdateUser updates an existing user
func (uc *userUseCase) UpdateUser(ctx context.Context, id uuid.UUID, req UpdateUserRequest) (*UserResponse, error) {
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return nil, apperrors.NewNotFoundError("user", "user not found")
	}

	// Update user profile
	user.UpdateProfile(req.FirstName, req.LastName)

	// Save updated user
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return toUserResponse(user), nil
}

// DeleteUser deletes a user by ID
func (uc *userUseCase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return apperrors.NewNotFoundError("user", "user not found")
	}

	if err := uc.userRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// ListUsers retrieves a list of users with pagination
func (uc *userUseCase) ListUsers(ctx context.Context, offset, limit int) (*ListUsersResponse, error) {
	users, err := uc.userRepo.List(ctx, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	total, err := uc.userRepo.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count users: %w", err)
	}

	userResponses := make([]*UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = toUserResponse(user)
	}

	return &ListUsersResponse{
		Users:  userResponses,
		Total:  total,
		Offset: offset,
		Limit:  limit,
	}, nil
}

// AuthenticateUser authenticates a user with email and password
func (uc *userUseCase) AuthenticateUser(ctx context.Context, email, password string) (*UserResponse, error) {
	// Validate email format
	emailVO, err := valueobjects.NewEmail(email)
	if err != nil {
		return nil, apperrors.NewValidationError("email", err.Error())
	}

	// Get user by email
	user, err := uc.userRepo.GetByEmail(ctx, emailVO.Value())
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	if user == nil {
		return nil, apperrors.NewUnauthorizedError("invalid email or password")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, apperrors.NewUnauthorizedError("account is deactivated")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, apperrors.NewUnauthorizedError("invalid email or password")
	}

	return toUserResponse(user), nil
}

// toUserResponse converts user entity to user response
func toUserResponse(user *entities.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		FullName:  user.GetFullName(),
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
