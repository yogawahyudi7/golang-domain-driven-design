package valueobjects

import (
	"errors"
	"regexp"
	"strings"
)

// Email represents an email value object
type Email struct {
	value string
}

// NewEmail creates a new Email value object
func NewEmail(email string) (*Email, error) {
	email = strings.TrimSpace(strings.ToLower(email))

	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}

	return &Email{value: email}, nil
}

// Value returns the email value
func (e *Email) Value() string {
	return e.value
}

// String returns the string representation of email
func (e *Email) String() string {
	return e.value
}

// isValidEmail validates email format
func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
