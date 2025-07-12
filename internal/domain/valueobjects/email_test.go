package valueobjects

import (
	"testing"
)

func TestNewEmail_Valid(t *testing.T) {
	validEmails := []string{
		"test@example.com",
		"user.name@domain.co.uk",
		"user+tag@example.org",
		"123@domain.com",
	}

	for _, email := range validEmails {
		emailVO, err := NewEmail(email)
		if err != nil {
			t.Errorf("Expected valid email %s, got error: %v", email, err)
		}

		if emailVO.Value() != email {
			t.Errorf("Expected email value %s, got %s", email, emailVO.Value())
		}
	}
}

func TestNewEmail_Invalid(t *testing.T) {
	invalidEmails := []string{
		"",
		"invalid-email",
		"@domain.com",
		"user@",
		"user@domain",
		"user space@domain.com",
	}

	for _, email := range invalidEmails {
		_, err := NewEmail(email)
		if err == nil {
			t.Errorf("Expected invalid email %s to return error", email)
		}
	}
}

func TestNewEmail_Normalization(t *testing.T) {
	email := "  Test@EXAMPLE.COM  "
	expectedEmail := "test@example.com"

	emailVO, err := NewEmail(email)
	if err != nil {
		t.Errorf("Expected valid email, got error: %v", err)
	}

	if emailVO.Value() != expectedEmail {
		t.Errorf("Expected normalized email %s, got %s", expectedEmail, emailVO.Value())
	}
}
