package valueobjects

import (
	"errors"
	"strings"
	"unicode"
)

// Password represents a password value object
type Password struct {
	value string
}

// NewPassword creates a new Password value object
func NewPassword(password string) (*Password, error) {
	password = strings.TrimSpace(password)

	if err := validatePassword(password); err != nil {
		return nil, err
	}

	return &Password{value: password}, nil
}

// Value returns the password value
func (p *Password) Value() string {
	return p.value
}

// validatePassword validates password strength
func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	if len(password) > 128 {
		return errors.New("password must be less than 128 characters")
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}

	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}

	if !hasNumber {
		return errors.New("password must contain at least one number")
	}

	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
