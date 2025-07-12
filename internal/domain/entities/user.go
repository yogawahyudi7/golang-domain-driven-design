package entities

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user entity in the domain
type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Username  string    `json:"username" gorm:"unique;not null"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	Password  string    `json:"-" gorm:"not null"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser creates a new user entity
func NewUser(email, username, firstName, lastName, hashedPassword string) *User {
	return &User{
		ID:        uuid.New(),
		Email:     email,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Password:  hashedPassword,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// UpdateProfile updates user profile information
func (u *User) UpdateProfile(firstName, lastName string) {
	u.FirstName = firstName
	u.LastName = lastName
	u.UpdatedAt = time.Now()
}

// ChangePassword updates user password
func (u *User) ChangePassword(hashedPassword string) {
	u.Password = hashedPassword
	u.UpdatedAt = time.Now()
}

// Deactivate deactivates the user
func (u *User) Deactivate() {
	u.IsActive = false
	u.UpdatedAt = time.Now()
}

// Activate activates the user
func (u *User) Activate() {
	u.IsActive = true
	u.UpdatedAt = time.Now()
}

// GetFullName returns the full name of the user
func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}
