package test

import (
	"testing"

	"golang-domain-driven-design/internal/domain/entities"

	"github.com/google/uuid"
)

func TestNewUser(t *testing.T) {
	// Test data
	email := "test@example.com"
	username := "testuser"
	firstName := "John"
	lastName := "Doe"
	password := "hashedpassword"

	// Create user
	user := entities.NewUser(email, username, firstName, lastName, password)

	// Assertions
	if user.Email != email {
		t.Errorf("Expected email %s, got %s", email, user.Email)
	}

	if user.Username != username {
		t.Errorf("Expected username %s, got %s", username, user.Username)
	}

	if user.FirstName != firstName {
		t.Errorf("Expected first name %s, got %s", firstName, user.FirstName)
	}

	if user.LastName != lastName {
		t.Errorf("Expected last name %s, got %s", lastName, user.LastName)
	}

	if user.Password != password {
		t.Errorf("Expected password %s, got %s", password, user.Password)
	}

	if !user.IsActive {
		t.Error("Expected user to be active")
	}

	if user.ID == (uuid.UUID{}) {
		t.Error("Expected ID to be generated")
	}
}

func TestUser_UpdateProfile(t *testing.T) {
	user := entities.NewUser("test@example.com", "testuser", "John", "Doe", "password")
	originalUpdatedAt := user.UpdatedAt

	// Update profile
	newFirstName := "Jane"
	newLastName := "Smith"
	user.UpdateProfile(newFirstName, newLastName)

	// Assertions
	if user.FirstName != newFirstName {
		t.Errorf("Expected first name %s, got %s", newFirstName, user.FirstName)
	}

	if user.LastName != newLastName {
		t.Errorf("Expected last name %s, got %s", newLastName, user.LastName)
	}

	if !user.UpdatedAt.After(originalUpdatedAt) {
		t.Error("Expected UpdatedAt to be updated")
	}
}

func TestUser_GetFullName(t *testing.T) {
	user := entities.NewUser("test@example.com", "testuser", "John", "Doe", "password")
	expectedFullName := "John Doe"

	if user.GetFullName() != expectedFullName {
		t.Errorf("Expected full name %s, got %s", expectedFullName, user.GetFullName())
	}
}

func TestUser_Deactivate(t *testing.T) {
	user := entities.NewUser("test@example.com", "testuser", "John", "Doe", "password")
	originalUpdatedAt := user.UpdatedAt

	// Deactivate user
	user.Deactivate()

	// Assertions
	if user.IsActive {
		t.Error("Expected user to be inactive")
	}

	if !user.UpdatedAt.After(originalUpdatedAt) {
		t.Error("Expected UpdatedAt to be updated")
	}
}

func TestUser_Activate(t *testing.T) {
	user := entities.NewUser("test@example.com", "testuser", "John", "Doe", "password")
	user.Deactivate() // First deactivate
	originalUpdatedAt := user.UpdatedAt

	// Activate user
	user.Activate()

	// Assertions
	if !user.IsActive {
		t.Error("Expected user to be active")
	}

	if !user.UpdatedAt.After(originalUpdatedAt) {
		t.Error("Expected UpdatedAt to be updated")
	}
}
