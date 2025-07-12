package errors

import "fmt"

// AppError represents the base application error
type AppError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	return e.Message
}

// ValidationError represents a validation error
type ValidationError struct {
	*AppError
	Field string `json:"field"`
}

// NewValidationError creates a new validation error
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		AppError: &AppError{
			Type:    "validation_error",
			Message: message,
		},
		Field: field,
	}
}

// NotFoundError represents a not found error
type NotFoundError struct {
	*AppError
	Resource string `json:"resource"`
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(resource, message string) *NotFoundError {
	return &NotFoundError{
		AppError: &AppError{
			Type:    "not_found_error",
			Message: message,
		},
		Resource: resource,
	}
}

// UnauthorizedError represents an unauthorized error
type UnauthorizedError struct {
	*AppError
}

// NewUnauthorizedError creates a new unauthorized error
func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		AppError: &AppError{
			Type:    "unauthorized_error",
			Message: message,
		},
	}
}

// ConflictError represents a conflict error
type ConflictError struct {
	*AppError
	Resource string `json:"resource"`
}

// NewConflictError creates a new conflict error
func NewConflictError(resource, message string) *ConflictError {
	return &ConflictError{
		AppError: &AppError{
			Type:    "conflict_error",
			Message: message,
		},
		Resource: resource,
	}
}

// InternalError represents an internal server error
type InternalError struct {
	*AppError
	OriginalError error `json:"-"`
}

// NewInternalError creates a new internal error
func NewInternalError(message string, originalError error) *InternalError {
	return &InternalError{
		AppError: &AppError{
			Type:    "internal_error",
			Message: message,
		},
		OriginalError: originalError,
	}
}

// Unwrap returns the original error
func (e *InternalError) Unwrap() error {
	return e.OriginalError
}

// String returns a string representation of the error
func (e *InternalError) String() string {
	if e.OriginalError != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.OriginalError)
	}
	return e.Message
}
